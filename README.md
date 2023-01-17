# api-demo
Purpose of this is for Go practice. Basic CRUD app with Go, Postgres, and Docker.

## Docker
Build the image by running this command in the root directory container the Dockerfile, supplement `<image-name>` for the name of the docker image.

`docker build -t api-demo .`

Spin up the docker container. 

```
docker run -d --rm -p 8080:8080 \
    -e ANIMALS_API_USERNAME=<api_username> \
    -e ANIMALS_API_PASSWORD=<api_password> \
    -e POSTGRES_HOST=docker.host.internal \
    -e POSTGRES_PORT=5432 \
    -e POSTGRES_USER=<postgres_username> \
    -e POSTGRES_PASSWORD=<postgres_password> \
    -e POSTGRES_DB=<postgres_db> \
    api-demo
```
