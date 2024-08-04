# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

# TODO: set the correct architecture for deployment
RUN CGO_ENABLED=0 GOOS=linux go build -o /consumer 

CMD ["/consumer"]