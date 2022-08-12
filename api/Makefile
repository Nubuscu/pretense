
run:
	@uvicorn src.main:app --reload --port=5000

format:
	isort src/*
	black src/


re-up:
	docker-compose down
	docker-compose up --build --detach
	# PYTHONPATH=. python content/dump_from_spotify.py
	# PYTHONPATH=. python content/write_topic.py all