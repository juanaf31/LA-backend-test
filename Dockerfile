FROM golang:alpine as build-env
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /app/main .
COPY --from=build-env /app/.env .
COPY --from=build-env /app/.env.dev .
COPY --from=build-env /app/.env.prod .
EXPOSE 8080
CMD ["./main"]