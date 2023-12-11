.PHONY: start-psql
start-psql:
	docker-compose \
		-f deployment/docker/postgres.docker-compose.yaml \
		--env-file=.env \
		up -d
args=
.PHONY: stop-psql
stop-psql:
	docker-compose \
		-f deployment/docker/postgres.docker-compose.yaml \
		--env-file=.env \
		down ${args}
