FROM golang:alpine AS build

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY internal ./internal
COPY cmd ./cmd
COPY pkg ./pkg

RUN go build -o /app/bin ./cmd/main.go

FROM alpine:3.21

RUN apk --no-cache add curl

COPY --from=build /app/bin /app/bin

ENTRYPOINT ["/app/bin"]