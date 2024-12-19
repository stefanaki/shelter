FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api cmd/api/*.go

FROM scratch

WORKDIR /app

# COPY --from=builder /etc/ssl/ca-certificates.crt /etc/ssl/certs

COPY --from=builder /app/api .

EXPOSE 8080

CMD [ "./api" ]