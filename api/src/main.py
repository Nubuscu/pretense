import logging
import os

from fastapi import APIRouter, FastAPI
from fastapi.middleware.cors import CORSMiddleware

from routes import albums, artists, reviews, topics

LOG = logging.getLogger(__name__)


def list_from_env_var(raw):
    try:
        return [clean for item in raw.split(",") if (clean := item.strip())]
    except AttributeError:
        return []


def create_app():
    app = FastAPI()
    origins = [
        "http://localhost",
        "http://localhost:3000",
    ]
    origins += list_from_env_var(os.environ.get("CORS_ALLOW_LIST"))
    LOG.info("Allowed origins: %s", origins)
    app.add_middleware(
        CORSMiddleware,
        allow_origins=origins,
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    v1_router = APIRouter(prefix="/v1")
    v1_router.include_router(albums, prefix="/albums")
    v1_router.include_router(artists, prefix="/artists")
    v1_router.include_router(reviews, prefix="/reviews")
    v1_router.include_router(topics, prefix="/topics")

    app.include_router(v1_router)
    return app


app = create_app()
