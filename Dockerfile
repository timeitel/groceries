FROM golang:1.23 AS dev
WORKDIR /app
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest
RUN curl -sLo /usr/local/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.7/tailwindcss-linux-x64
RUN chmod +x /usr/local/bin/tailwindcss
COPY go.mod go.sum ./
RUN go mod download
RUN apt update && apt install sqlite3
COPY . .
RUN mkdir /data
RUN sqlite3 /data/groceries.db < ./internal/infrastructure/data/schema.sql
RUN tailwindcss -i ./internal/web/static/css/input.css -o ./internal/web/static/css/output.css --minify
CMD ["air"]

FROM golang:1.23 AS builder
WORKDIR /out
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o ./main ./cmd/web

FROM alpine:3.20 AS runner
WORKDIR /app
COPY /internal/web/views/ ./internal/web/views/
COPY --from=builder /out/main ./
COPY --from=dev /app/internal/web/static/css/output.css ./web/static/css/output.css
COPY ./internal/web/static/favicon.ico ./static/favicon.ico
CMD ["./main"]
