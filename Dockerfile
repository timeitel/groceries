FROM node:lts-slim AS node
WORKDIR /src
COPY /static/css/input.css /input.css
RUN npm i -g tailwindcss
RUN tailwindcss -i /input.css -o /output.css --minify

FROM golang:1.22
WORKDIR /app
COPY --from=node /output.css /static/css/output.css
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd
CMD ["./main"]
