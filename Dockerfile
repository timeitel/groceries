FROM alpine:3.20 AS builder-tw
WORKDIR /out
COPY /static/css/ tailwind.config.js ./
COPY /internal/views/ ./internal/views/
RUN apk add --no-cache curl gcompat build-base
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN chmod +x tailwindcss-linux-arm64
RUN mv tailwindcss-linux-arm64 ./tailwindcss
RUN ./tailwindcss -i ./input.css -o ./output.css --minify

FROM golang:1.22 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN mv ./tailwindcss-linux-arm64 ./tailwindcss
RUN chmod +x tailwindcss
CMD ["air"]

FROM golang:alpine AS builder-go
WORKDIR /out
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN go build -o ./main ./cmd

FROM alpine:3.20 AS runner
WORKDIR /app
COPY /internal/views/ ./internal/views/
COPY --from=builder-go /out/main .
COPY --from=builder-tw /out/output.css ./static/css/output.css
COPY /static/favicon.ico ./static/favicon.ico
CMD ["./main"]
