from fastapi import APIRouter

health = APIRouter()


@health.get("/")
def say_hi():
    return "hello there!"
