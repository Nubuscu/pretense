"""A script to pull data from spotify and put it in the db.
Gets all liked albums and inserts them, along with their artists
Intended as a one-off pull but should be safely re-runnable.
"""
import multiprocessing
import base64
import hashlib
import logging
import random
from fastapi import FastAPI
import fastapi
import requests
import string

from python_graphql_client import GraphqlClient
import uvicorn


logging.basicConfig(level=logging.INFO)
LOG = logging.getLogger("spotify")
# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly

CLIENT_ID = "4f5bb5185cd044c5a5e58e22d224587c"
REDIRECT_URL = "http://localhost:12345/callback"

# note: hardcoded market
URL_FMT = "https://api.spotify.com/v1/me/albums?limit={limit}&offset={offset}&market=NZ"


CLIENT = GraphqlClient(endpoint="http://localhost:8081/query", verify=False)

MUTATION = """
mutation createAlbumAndArtists($name: String!, $artists: [CreateArtistInput!]!) {
  createAlbumAndArtists(album: {name: $name}, artists: $artists) {
    id
  }
}
"""


def generate_code() -> tuple[str, str]:
    rand = random.SystemRandom()
    code_verifier = "".join(rand.choices(string.ascii_letters + string.digits, k=128))

    code_sha_256 = hashlib.sha256(code_verifier.encode("utf-8")).digest()
    b64 = base64.urlsafe_b64encode(code_sha_256)
    code_challenge = b64.decode("utf-8").replace("=", "")

    return code_verifier, code_challenge


def fastapi_localhost(queue: multiprocessing.Queue):
    app = FastAPI()

    @app.get("/callback")
    def callback(request: fastapi.Request):
        queue.put(request.url)
        return fastapi.responses.HTMLResponse(
            """
    <html>
      <body>
        <h1>All done.</h1>
        <h3>You may close this tab now.</h3>
        <script>
          window.setTimeout(() => {
            window.close()
          }, 1000)
        </script>
      </body>
    </html>
    """
        )

    return uvicorn.run(app, host="127.0.0.1", port=12345, reload=False, workers=1)


def login() -> str:
    verifier, challenge = generate_code()
    state = "abcde"
    url = "https://accounts.spotify.com/authorize?" + "&".join(
        [
            f"{k}={v}"
            for k, v in {
                "response_type": "code",
                "client_id": CLIENT_ID,
                "scope": "user-library-read",
                "redirect_uri": REDIRECT_URL,
                "state": state,
                "code_challenge_method": "S256",
                "code_challenge": challenge,
            }.items()
        ]
    )
    queue = multiprocessing.Queue()
    fastapi_proc = multiprocessing.Process(target=fastapi_localhost, args=(queue,))
    fastapi_proc.start()
    print(f"go to {url}")

    callback_url = queue.get(block=True)
    fastapi_proc.terminate()

    raw_query = callback_url.query
    query_params = {k: v for k, v in [pair.split("=") for pair in raw_query.split("&")]}
    code = query_params["code"]

    # now we have the code for pkce, get an access token
    response = requests.post(
        url="https://accounts.spotify.com/api/token",
        headers={"Content-Type": "application/x-www-form-urlencoded"},
        data={
            "grant_type": "authorization_code",
            "code": code,
            "redirect_uri": REDIRECT_URL,  # not actually used, just for verification
            "client_id": CLIENT_ID,
            "code_verifier": verifier,
        },
    )
    if not response.ok:
        print(response.text)
        breakpoint()

    # format like:
    # {
    #     "access_token": "",
    #     "token_type": "Bearer",
    #     "expires_in": 3600,
    #     "refresh_token": "",
    # }
    return response.json()


def parse_and_insert(items):
    for item in items:
        album_name = item.get("album", {}).get("name")
        raw_artists = item.get("album", {}).get("artists", [])
        _vars = {
            "name": filter_album_name(album_name),
            "artists": [{"name": a.get("name")} for a in raw_artists],
        }
        LOG.info("inserting %s", _vars)
        CLIENT.execute(query=MUTATION, variables=_vars)


def filter_album_name(name):
    if any(
        substr in name.lower()
        for substr in ["(deluxe", "(special", "(remixed", "(with bonus"]
    ):
        clean = name.split("(")[0].strip()
        LOG.warning("modified album name %s, original was %s", clean, name)
        return clean
    return name


def main():
    token_resp = login()
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {token_resp['access_token']}"})

    limit = 50
    offset = 0
    keep_going = True
    while keep_going:
        resp = session.get(URL_FMT.format(limit=limit, offset=offset))
        LOG.info("resp: %s, offset: %s", resp, offset)
        assert resp.ok, resp.text
        items = resp.json().get("items")
        offset += limit
        keep_going = len(items) >= limit
        parse_and_insert(items)


if __name__ == "__main__":
    try:
        main()
    except Exception as err:
        print(
            "Try a new token: "
            "https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market="
        )
        raise
