FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY go.* .
COPY . .
RUN  go build -o main main.go

FROM alpine:3.6
LABEL Authors="Asem&Dias" Project="ASCII-ART-WEB-DOCKEIZE" Date="21.06.2022"
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/internal/utils/banner /app/internal/utils/banner/
COPY --from=builder /app/templates /app/templates

CMD ["/app/main"]

