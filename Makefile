run:
	docker compose up --build -d

build:
	docker compose build

down:
	docker compose down

restart:
	docker compose down
	docker compose up -d

ps:
	docker compose ps

logs:
	docker compose logs -f