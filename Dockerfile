FROM golang:1.22 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN chmod +x /usr/local/bin/tailwindcss
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ["tailwindcss", "-i", "./static/css/input.css", "-o", "./static/css/output.css", "--minify"]
CMD ["air"]

FROM golang:alpine AS builder
WORKDIR /out
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN go build -o ./main ./cmd

FROM alpine:3.20 AS runner
WORKDIR /app
COPY /internal/views/ ./internal/views/
COPY --from=builder /out/main ./
COPY --from=dev /app/static/css/output.css ./static/css/output.css
COPY /static/favicon.ico ./static/favicon.ico
CMD ["./main"]
