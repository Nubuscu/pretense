"""A script to pull data from spotify and put it in the db.

Gets all liked albums and inserts them, along with their artists

Intended as a one-off pull but should be safely re-runnable.
"""
import requests
from pymongo.collection import Collection

from src.db.db import get_db, safe_insert

# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly
TOKEN = "BQDGI4Y7ADQsBLKNIEogowgMw56yqjF1mXLkpiSUE_B6AkQPqUji6ZkHKCEJjl_k3_BB41OQdD4A76Bo6_hZDQ-rbvJUUXZjjFBB58V_vLy6gJWh9-gWDcpruQJGY7xZeM-KlofbzS5vw9xiO76wYXke6CsPPs-1FJi7GfFK-kZDKQqfn_ZbkzeZ"

# note: hardcoded market
URL_FMT = "https://api.spotify.com/v1/me/albums?limit={limit}&offset={offset}&market=NZ"


def parse_and_insert(items):
    for item in items:
        album_title = item.get("album", {}).get("name")
        raw_artists = item.get("album", {}).get("artists", [])
        artist_names = [a.get("name") for a in raw_artists]

        db = get_db()
        artists_table: Collection = db["artists"]
        albums_table: Collection = db["albums"]
        print(f"Processing {album_title} by {artist_names}")
        # search artists for name
        # search albums for name
        # otherwise insert
        db_artists = []
        for artist_name in artist_names:
            db_artists.append(
                safe_insert(
                    artists_table,
                    {"name": artist_name},
                    {"name": artist_name},
                )
            )
        safe_insert(
            albums_table,
            {"name": album_title},
            {"name": album_title, "artists": [a["id"] for a in db_artists]},
        )


def main():
    session = requests.Session()
    session.headers.update({"Authorization": f"Bearer {TOKEN}"})

    limit = 25
    offset = 0
    keep_going = True
    while keep_going:
        resp = session.get(URL_FMT.format(limit=limit, offset=offset))
        print(resp, offset)
        assert resp.ok, resp.text
        items = resp.json().get("items")
        offset += limit
        keep_going = len(items) >= limit
        parse_and_insert(items)


if __name__ == "__main__":
    main()
