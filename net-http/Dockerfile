# build stage
FROM golang:1.19-buster AS build

WORKDIR /app

RUN mkdir ./build

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/app ./cmd/main.go

# deploy stage
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build ./bin/app ./app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]
