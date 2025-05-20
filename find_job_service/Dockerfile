FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o find-service-app ./cmd/app.go

CMD ["./find-service-app"]