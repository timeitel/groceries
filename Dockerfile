FROM node:lts-slim AS builder-node
WORKDIR /src
COPY /static/css/ tailwind.config.js ./
COPY /internal/views/ ./internal/views/
RUN npm i -g tailwindcss
RUN tailwindcss -i ./input.css -o ./output.css --minify

FROM golang:alpine as builder-go
WORKDIR /app
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN go build -o ./main ./cmd

FROM alpine:3.20.2 as runner
WORKDIR /app
COPY /internal/views/ ./internal/views/
COPY --from=builder-go /app/main .
COPY --from=builder-node /src/output.css ./static/css/output.css
COPY /static/favicon.ico ./static/favicon.ico
CMD ["./main"]
