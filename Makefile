test:
	go test ./... -v -bench . -failfast -cover -count=1

docker_up:
	docker-compose up -d

docker_stop:
	docker-compose stop

docker_down:
	docker-compose down

docker_status:
	docker-compose ps

