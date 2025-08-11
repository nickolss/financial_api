FROM golang:1.24.5 AS builder

WORKDIR /app

COPY . .

RUN go build -o /financial_api

FROM debian:stable-slim

ARG PORT=8080
ENV PORT=${PORT}

WORKDIR /

RUN useradd -m app_user

USER app_user

COPY --from=builder /financial_api /financial_api
RUN command
EXPOSE ${PORT}

ENTRYPOINT ["/financial_api"]
