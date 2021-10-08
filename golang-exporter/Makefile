PROJECT_NAME = prom-go

.PHONY: build tag push docker_all
build:
	docker build . -t $(PROJECT_NAME)

tag:
	docker tag $(PROJECT_NAME) localhost:5000/$(PROJECT_NAME)

push:
	docker push localhost:5000/$(PROJECT_NAME)

docker_all: build tag push
