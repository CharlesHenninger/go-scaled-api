FROM golang:latest as base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 1323

FROM base as build
RUN go build -o main . 
CMD ["./main"]