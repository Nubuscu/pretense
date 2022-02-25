from typing import Optional

from fastapi import APIRouter, FastAPI

from routes import (
    albums,
    artists,
    reviews,
    topics,
)


def create_app():
    app = FastAPI()
    v1_router = APIRouter(prefix="/v1")
    v1_router.include_router(albums, prefix="/albums")
    v1_router.include_router(artists, prefix="/artists")
    v1_router.include_router(reviews, prefix="/reviews")
    v1_router.include_router(topics, prefix="/topics")

    app.include_router(v1_router)
    return app


app = create_app()
