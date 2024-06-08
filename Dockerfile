FROM golang:1.22.3-alpine3.20 AS server_builder

RUN apk add gcc libc-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . go-backend/.

WORKDIR go-backend
RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -a cmd/server/main.go

FROM scratch
EXPOSE 8080
COPY --from=server_builder /app/go-backend/main .
COPY --from=server_builder /app/go-backend/config/config.yaml .
CMD ["./main"]