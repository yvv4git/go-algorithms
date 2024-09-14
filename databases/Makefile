# Environment file
ENV_FILE := .env
ifneq (,$(wildcard ${ENV_FILE}))
    include ${ENV_FILE}
endif

# Test environment
DOCKER_COMPOSE_DEV_ENV ?= docker-compose.yaml
DOCKER_COMPOSE_DEV_ENV_SERVICES ?= postgresql mysql
DOCKER_COMPOSE_RUN?=docker compose -f ${DOCKER_COMPOSE_DEV_ENV} --project-name=db-algorithms --env-file=${ENV_FILE}

# Dev environment
.PHONY: dev-env-up
dev-env-up:
	${DOCKER_COMPOSE_RUN} up -d ${DOCKER_COMPOSE_DEV_ENV_SERVICES}

.PHONY: dev-env-down
dev-env-down:
	${DOCKER_COMPOSE_RUN} down

.PHONY: dev-env-clear
dev-env-clear: dev-env-down
	# Remove containers and volumes
	${DOCKER_COMPOSE_RUN} down -v