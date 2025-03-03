include .env.development
CURRENT_DIR=$(shell pwd)

APP=$(shell basename $(shell pwd))
APP_CMD_DIR=$(shell pwd)/cmd

createdb:
	docker exec -it ${POSTGRES_DATABASE} createdb --username=${PostgresDockerContainerUser} --owner=${PostgresDockerContainerUser} ${PostgresDatabase}

dropdb:
	docker exec -it ${POSTGRES_DATABASE} dropdb ${PostgresDatabase} --username=${PostgresDockerContainerUser}

psqlcontainer:
	docker run --name ${POSTGRES_DATABASE} -d -p ${PostgresPort}:5432 -e POSTGRES_PASSWORD=${PostgresDockerContainerPassword} postgres:15-alpine3.16

rediscontainer:
	docker run --name ${RedisDockerContainerName} -p ${RedisPort}:6379 -d --env-file .env.development redis:7.0.5-alpine3.16

# redisdocker:
# 	docker run -d --cap-add sys_resource --name ${RedisDockerContainerName} -p ${REDIS_HTTPS_PORT}:8443 -p ${REDIS_REST_API_PORT}:9443 -p ${RedisPort}:12000 redislabs/redis

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rsync -rv --exclude={'/.git','LICENSE','README.md'} ${CURRENT_DIR}/protos/* ${CURRENT_DIR}/okslibprotos

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://${PostgresDockerContainerUser}:${PostgresDockerContainerPassword}@${PostgresHost}:${PostgresPort}/${PostgresDatabase}?sslmode=${PostgresSSLMode}' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://${PostgresDockerContainerUser}:${PostgresDockerContainerPassword}@${PostgresHost}:${PostgresPort}/${PostgresDatabase}?sslmode=${PostgresSSLMode}' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

swag-init:
	swag init -g api/api.go -o api/docs

.PHONE: pull-proto-module, update-proto-module, copy-proto-module, migration-up, migration-down, build, build-image, push-image, swag-init, psqlcontainer, rediscontainer
