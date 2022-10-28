FROM python:3.10-slim-buster as base

COPY requirements.in /requirements.in
RUN pip install --no-cache-dir --upgrade -r requirements.in
COPY src /app/src

WORKDIR /app

EXPOSE 5000

CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "5000"]