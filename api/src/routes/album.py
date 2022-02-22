from fastapi import APIRouter

albums = APIRouter()


@albums.get("/")
def get():
    return {"foo": "bar"}
