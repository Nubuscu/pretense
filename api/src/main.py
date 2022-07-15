import logging
import os

from fastapi import APIRouter, FastAPI
from fastapi.middleware.cors import CORSMiddleware
from src.models import Album, Artist, Topic
from src.db.db import get_topics, get_albums, get_artists
from src.routes import router_for_collection

from src.routes import health

from fastapi.logger import logger as fastapi_logger

gunicorn_error_logger = logging.getLogger("gunicorn.error")
gunicorn_logger = logging.getLogger("gunicorn")
uvicorn_access_logger = logging.getLogger("uvicorn.access")
uvicorn_access_logger.handlers = gunicorn_error_logger.handlers

fastapi_logger.handlers = gunicorn_error_logger.handlers

if __name__ != "__main__":
    fastapi_logger.setLevel(gunicorn_logger.level)
else:
    fastapi_logger.setLevel(logging.DEBUG)
LOG = logging.getLogger(__name__)


def list_from_env_var(raw):
    try:
        return [clean for item in raw.split(",") if (clean := item.strip())]
    except AttributeError:
        return []


def create_app():
    app = FastAPI()
    origins = ["*"]
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
    topics_router = router_for_collection(get_topics, Topic)
    albums_router = router_for_collection(get_albums, Album)
    artists_router = router_for_collection(get_artists, Artist)
    v1_router.include_router(topics_router, prefix="/topics")
    v1_router.include_router(albums_router, prefix="/albums")
    v1_router.include_router(artists_router, prefix="/artists")
    v1_router.include_router(health, prefix="/health")

    app.include_router(v1_router)
    return app


app = create_app()
