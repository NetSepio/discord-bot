FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN apk add --no-cache git && go build -o discord-sepio . && apk del git

FROM alpine
WORKDIR /app
COPY --from=builder /app/discord-sepio .
RUN apk add --no-cache chromium
CMD [ "./discord-sepio" ]