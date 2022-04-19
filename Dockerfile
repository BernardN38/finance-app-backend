FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o financefirst financefirst.go



FROM alpine

COPY --from=builder /app/financefirst /app/financefirst
COPY --from=builder /app/build ./build

EXPOSE 8080

CMD ["/app/financeapp"]