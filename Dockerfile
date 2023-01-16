FROM golang:1.19.5

ARG username
ENV ANIMALS_API_USERNAME=$username

ARG password
ENV ANIMALS_API_PASSWORD=$password

RUN mkdir /animals_api

COPY . /animals_api

WORKDIR /animals_api

RUN go build

CMD ["./apidemo"]