# api-demo
Purpose of this is for Go practice. Basic CRUD app with Go, Postgres, and Docker.

## Docker
Build the image by running this command in the root directory container the Dockerfile, supplement `<image-name>` for the name of the docker image.

`docker build -t api-demo .`

Spin up the docker container called `animals_api`. Here I am using the default for Postgres on the host machine.

```
docker run -d --rm -p 8080:8080 \
    -e ANIMALS_API_USERNAME=<api_username> \
    -e ANIMALS_API_PASSWORD=<api_password> \
    -e POSTGRES_HOST=host.docker.internal \
    -e POSTGRES_PORT=5432 \
    -e POSTGRES_USER=<postgres_username> \
    -e POSTGRES_PASSWORD=<postgres_password> \
    -e POSTGRES_DB=<postgres_db> \
    --add-host host.docker.internal:host-gateway \
    --name animals_api \
    api-demo
```
