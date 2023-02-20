FROM golang:1.18 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0
CMD ["air"]
