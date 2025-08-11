FROM golang:1.24.6-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /financial_api

FROM scratch

ARG PORT=8080
ENV PORT=${PORT}

WORKDIR /

COPY --from=builder /financial_api /financial_api

EXPOSE ${PORT}

ENTRYPOINT ["/financial_api"]
