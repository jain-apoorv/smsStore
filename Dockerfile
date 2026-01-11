# ---- build stage ----
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o sms-store

# ---- run stage ----
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/sms-store .
EXPOSE 8081
CMD ["./sms-store"]
