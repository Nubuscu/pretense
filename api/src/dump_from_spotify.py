"""A script to pull data from spotify and put it in the db.

Gets all liked albums and inserts them, along with their artists

Intended as a one-off pull but should be safely re-runnable.
"""
import requests
from pymongo.collection import Collection

from src.graph import GraphRepository

# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly
TOKEN = ""

# note: hardcoded market
URL_FMT = "https://api.spotify.com/v1/me/albums?limit={limit}&offset={offset}&market=NZ"

repo = GraphRepository()


def parse_and_insert(items):
    for item in items:
        album_title = item.get("album", {}).get("name")
        raw_artists = item.get("album", {}).get("artists", [])
        artist_names = [a.get("name") for a in raw_artists]
        repo.upsert_album(album_title, artist_names)


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
