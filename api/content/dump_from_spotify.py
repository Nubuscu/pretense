"""A script to pull data from spotify and put it in the db.

Gets all liked albums and inserts them, along with their artists

Intended as a one-off pull but should be safely re-runnable.
"""
import os
import logging
import requests
from pymongo.collection import Collection

from src.graph import GraphRepository

logging.basicConfig(level=logging.INFO)
LOG = logging.getLogger("spotify")
# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly
TOKEN = "BQBRNnSF-1X6CCxuewH2U1O_i5o519ACrdH3o3u4sHMArUjtD4jHgQ3uw06aIuheacZBgaryTn5-gl3V1ZXS3SHD7_AvQlbq6qPdAHCrwMlkq1a3e_xsQTv0YI_NWLSGlqeR6W9kT6lsMt6F3Ijf97aX2EQWyGvr1jCyzzxyJ4dt_AiwD5FnoIbR"

# note: hardcoded market
URL_FMT = "https://api.spotify.com/v1/me/albums?limit={limit}&offset={offset}&market=NZ"

repo = GraphRepository(
    host=os.environ["DB_HOST"], port=os.environ["DB_PORT"], username=None, password=None
)


def parse_and_insert(items):
    with repo as ctx_repo:
        for item in items:
            album_title = item.get("album", {}).get("name")
            raw_artists = item.get("album", {}).get("artists", [])
            artist_names = [a.get("name") for a in raw_artists]
            ctx_repo.upsert_album(filter_album_name(album_title), artist_names)


def filter_album_name(name):
    if any(substr in name.lower() for substr in ["(deluxe", "(special", "(remixed"]):
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
    main()