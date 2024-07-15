FROM golang:alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/cmd/sso
RUN go build -o /src/app

FROM alpine
WORKDIR /app
COPY --from=build /src/app ./
COPY ./config/prod.yaml ./
ENV CONFIG_PATH="prod.yaml"
CMD ["./app"]