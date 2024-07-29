FROM node:lts-slim AS node
WORKDIR /src
COPY /static/css/input.css /input.css
RUN npm i -g tailwindcss
RUN tailwindcss -i /input.css -o /output.css --minify

FROM golang:alpine as builder
WORKDIR /app
ENV CGO_ENABLED=1
COPY --from=node /output.css /static/css/output.css
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN go build -o /main ./cmd
# CMD ["/main"]
#
FROM alpine:3.20.2 as runner
WORKDIR /app
# RUN apk --no-cache add ca-certificates tzdata libc6-compat libgcc libstdc++
COPY --from=builder /main /
COPY . .
CMD ["/main"]
