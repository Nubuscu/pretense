"""A script to pull data from spotify and put it in the db.

Gets all liked albums and inserts them, along with their artists

Intended as a one-off pull but should be safely re-runnable.
"""
import requests

from db.db import get_cursor

# a token for the spotify api. e.g.:
# https://developer.spotify.com/console/get-current-user-saved-albums/?limit=&offset=&market=
# may expire relatively quickly
TOKEN = ""

# note: hardcoded market
URL_FMT = "https://api.spotify.com/v1/me/albums?limit={limit}&offset={offset}&market=NZ"

ALBUM_INSERT_SQL = (
    "INSERT INTO album (title) VALUES (%(title)s) ON CONFLICT DO NOTHING;"
)
ALBUM_SELECT_SQL = "SELECT id FROM album WHERE title = %(title)s;"
ARTIST_INSERT_SQL = (
    "INSERT INTO artist (name) VALUES (%(name)s) ON CONFLICT DO NOTHING;"
)
ARTIST_SELECT_SQL = "SELECT id FROM artist WHERE name = %(name)s;"
REL_INSERT_SQL = "INSERT INTO rel_album_artist (album_id, artist_id) VALUES (%(album_id)s, %(artist_id)s) ON CONFLICT DO NOTHING;"


def parse_and_insert(items):
    for item in items:
        album_title = item.get("album", {}).get("name")
        raw_artists = item.get("album", {}).get("artists", [])
        artist_names = [a.get("name") for a in raw_artists]

        with get_cursor() as cursor:
            print(f"Inserting {album_title} by {artist_names}")
            cursor.execute(ALBUM_INSERT_SQL, {"title": album_title})
            cursor.execute(ALBUM_SELECT_SQL, {"title": album_title})
            album_id = cursor.fetchone()

            for name in artist_names:
                cursor.execute(ARTIST_INSERT_SQL, {"name": name})
                cursor.execute(ARTIST_SELECT_SQL, {"name": name})
                artist_id = cursor.fetchone()
                cursor.execute(
                    REL_INSERT_SQL, {"album_id": album_id, "artist_id": artist_id}
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
