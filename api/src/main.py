from typing import Optional

from fastapi import APIRouter, FastAPI

from routes.album import albums


def create_app():
    app = FastAPI()
    v1_router = APIRouter(prefix="/v1")
    v1_router.include_router(albums, prefix="/albums")

    app.include_router(v1_router)
    return app


app = create_app()
