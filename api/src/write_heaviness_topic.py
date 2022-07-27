from graph import get_cursor

INSERT_TOPIC_SQL = "INSERT INTO topic (name) VALUES (%(name)s)"
SELECT_TOPIC_SQL = "SELECT id FROM topic WHERE name = %(name)s"
INSERT_REVIEW_SQL = "INSERT INTO review (title, body) VALUES (%(title)s, %(body)s)"
SELECT_REVIEW_SQL = "SELECT id FROM review WHERE title = %(title)s"
INSERT_REVIEW_TOPIC_REL_SQL = "INSERT INTO rel_review_topic (review_id, topic_id) VALUES (%(review_id)s, %(topic_id)s)"

INSERT_ALBUM_RELATIONS_SQL = """
WITH cte_albums AS (
    SELECT id, %(topic_id)s FROM album
    WHERE title IN (
        'Koloss',
        'I Let It in and It Took Everything',
        'Thegodmachine',
        'Awakened',
        'måsstaden under vatten',
        'Oh What The Future Holds',
        'May Our Chambers Be Full',
        'Oxidized',
        'Absolute',
        'Menschenmühle',
        'Lifeblood',
        'Melancholy',
        'A Tear in the Fabric of Life',
        'Laugh Tracks'
        )
)
INSERT INTO rel_album_topic (album_id, topic_id) SELECT * FROM cte_albums
"""
title_id = "heaviness"
title = "An Investigation Into Perceived Heaviness"

body = """
One of the qualitative metrics that some metalheads like to use to describe their
favourite bands is how "heavy" they are.
This is usually a good thing, i.e. `heavier ~= more enjoyable`.
Of course heavier isn't always better, and taking it to the extreme tends to result in more niche bands

Considering the variety of subgenres, there appears to be a range of definitions of heavy.
Some of the contributing factors I've picked up by observation:

1. Atmosphere
1. Breakdowns
1. Lyrical content
1. Big riffs
1. Context and build-up


### Atmosphere
This is typically the focus of doom metal (and probably black metal but that's not my cup of tea),
creating a dark ambiance to either scare or depress the listener. Or both.
That looks worse written down but some people watch horror movies for fun ¯\_(ツ)_/¯.
Examples might include, in their own ways:
1. May Our Chambers Be Full by Emma Ruth Rundle and Thou
2. Loathe's I let it in...,
3. that one song on every Meshuggah album that's not 100% djenty math noises.
4. Menschenmühle by Kanonenfieber
I'd argue something like Frontierer's Oxidized fits this category too, but in the opposite way.
It's loud and oppressively noisy.

### Breakdowns
The obvious answer for most metalcore listeners, and in live shows, usually associated with circle pits,
walls of death, disrespecting one's surroundings, or hardcore dancing. This was my introduction to heavier sounding
things, especially in (hard) rock bands where they could scream during the bridge and still be radio friendly.

some more words go here, especially about metalcore, deathcore, etc.

### Lyrical content
Although metal bands don't always write the happiest of songs, A factor of heaviness is often the lyrics.
More of a factor to some, I'd like to think a little less so for me personally. 
Sometimes it's the clear, vocalized emotion attached with lyrical topics (King 810's song The Trauma Model),
the realization of what's being talked about, or just a breakdown line that makes you go "oof".

See also:
- The entire A Tear in the Fabric of Life EP (and the video for it too)
- Deadringer from Knocked Loose's Laugh Tracks album

### Big riffs
A catch-all for fun musicality that isn't strictly in the breakdown. Also the hardest to narrow down

more words go here too

### Context and build-up
Sometimes newer bands, especially metalcore, seem to write a long series of breakdowns rather than albums of songs.
Clearly if people are buying it, it can't be all bad, but I find it monotonous after a while.
By context, I'm meaning the rest of the song that makes it feel heavy by comparison, or contains callbacks to use later.
An example:

Halo by Machine Head. There are probably others but this one came to mind.
The outro riff is practically identical to the intro, but hits notably harder as (well probably some mixing shenanigans and half-time drums)
you've heard that musical idea before, spent the last while building up out of a mellower section.
And monkey brain likes repetition.


(title stolen from the 2021 Lionfight album of the same name)
"""

with get_cursor() as cursor:
    cursor.execute(INSERT_TOPIC_SQL, {"name": title_id})
    cursor.execute(SELECT_TOPIC_SQL, {"name": title_id})
    topic_id = cursor.fetchone()

    cursor.execute(INSERT_REVIEW_SQL, {"title": title, "body": body})
    cursor.execute(SELECT_REVIEW_SQL, {"title": title})
    review_id = cursor.fetchone()

    cursor.execute(
        INSERT_REVIEW_TOPIC_REL_SQL, {"review_id": review_id, "topic_id": topic_id}
    )

    cursor.execute(INSERT_ALBUM_RELATIONS_SQL, {"topic_id": topic_id})
