FROM golang:1.18 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
