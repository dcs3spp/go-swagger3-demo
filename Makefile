##
# Build
##

build-swagger:
	go-swagger3 --module-path . --output oas.yml --schema-without-pkg --generate-yaml true

build-api:
	docker-compose build api

docs:
	docker run --rm -p 3000:8080 -e SWAGGER_JSON=/tmp/oas.yml -v `pwd`:/tmp swaggerapi/swagger-ui

run:
	./create-rest-api-in-go-tutorial

##
# Docker
##

down:
	docker-compose down

up:
	docker-compose up -d

logs:
	docker-compose logs

