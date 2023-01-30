API_IMAGE_NAME = goal_account_api
API_CONTAINER_NAME = goal_account_api_service
API_PORT = 8100

ADMIN_IMAGE_NAME = goal_account_admin
ADMIN_CONTAINER_NAME = goal_account_admin_service
ADMIN_PORT = 8200

build-api:
	docker build --build-arg GOAL_APP=account_api -t $(API_IMAGE_NAME) .

build-admin:
	docker build --build-arg GOAL_APP=account_admin -t $(ADMIN_IMAGE_NAME) .

run-api:
	docker run -d --name $(API_CONTAINER_NAME) -p $(API_PORT):$(API_PORT) $(API_IMAGE_NAME)

run-admin:
	docker run -d --name $(ADMIN_CONTAINER_NAME) -p $(ADMIN_PORT):$(ADMIN_PORT) $(ADMIN_IMAGE_NAME)

clean:
	docker ps -a | grep $(API_IMAGE_NAME) | awk  '{print $$1}' | xargs docker stop
	docker ps -a | grep $(API_IMAGE_NAME) | awk  '{print $$1}' | xargs docker rm
	docker ps -a | grep $(ADMIN_IMAGE_NAME) | awk  '{print $$1}' | xargs docker stop
	docker ps -a | grep $(ADMIN_IMAGE_NAME) | awk  '{print $$1}' | xargs docker rm