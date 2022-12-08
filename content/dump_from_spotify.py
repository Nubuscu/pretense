"""A script to pull data from spotify and put it in the db.
Gets all liked albums and inserts them, along with their artists
Intended as a one-off pull but should be safely re-runnable.
"""
import os
import logging
import requests

from python_graphql_client import GraphqlClient


logging.basicConfig(level=logging.INFO)
LOG = logging.getLogger("spotify")
# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly
TOKEN = "BQD-SXJOByDOM2wVxG_uCI4Dp7mTyf-U7C0NeA6dAoHtH3irtAerdQUIWgrJY-patSaIHLFt2OcBzbKtB1cT2W6eLrJ1rZHmp4EuN9ihe0p6Y9I2QqC9F6jMKXLRN30FB0FaqmGzc2Kg3a5wH7-cpX8J9_sm9Pwexv76aD6lRplR96X3CR1CgCJZ"

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
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {TOKEN}"})

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
