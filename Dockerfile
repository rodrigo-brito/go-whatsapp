FROM golang:1.14 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o graphql cmd/graphql/graphql.go

FROM alpine
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/website/build /app/website/build
COPY --from=builder /app/graphql .
CMD ["/app/graphql"]
