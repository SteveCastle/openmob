#!/bin/bash
# Docker command to run an empty docker db with dev tern settings.
	docker run -p 8081:8080 -e SWAGGER_JSON=/schema/shrike.swagger.json -v $HOME/dev/openmob/packages/shrike/src/api/swagger/v1:/schema swaggerapi/swagger-ui
    docker ps