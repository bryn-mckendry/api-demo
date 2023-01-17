FROM golang:1.19.5

RUN mkdir /animals_api

COPY . /animals_api

WORKDIR /animals_api

RUN go build

CMD ["./apidemo"]