run:
	docker compose up --build -d

run-nocache:
	docker compose build --no-cache
	docker compose up -d

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


# remove unused <none> images 
remove-none-images:
	docker images -f "dangling=true"
	docker images -f "dangling=true" -q | xargs -r docker rmi
