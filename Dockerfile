FROM golang:1.20.5-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache git

RUN go build -o app .

FROM scratch

ENV HOST=0.0.0.0 PORT=8000

COPY --from=builder /app/app .

ENTRYPOINT [ "/app" ]
