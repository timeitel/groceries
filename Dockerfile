FROM alpine:3.20 as builder-tw
WORKDIR /out
COPY /static/css/ tailwind.config.js ./
COPY /internal/views/ ./internal/views/
RUN apk add --no-cache curl gcompat build-base
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN chmod +x tailwindcss-linux-arm64
RUN ./tailwindcss-linux-arm64 -i ./input.css -o ./output.css --minify

FROM golang:alpine as builder-go
WORKDIR /out
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN go build -o ./main ./cmd

FROM alpine:3.20 as runner
WORKDIR /app
COPY /internal/views/ ./internal/views/
COPY --from=builder-go /out/main .
COPY --from=builder-tw /out/output.css ./static/css/output.css
COPY /static/favicon.ico ./static/favicon.ico
CMD ./main
