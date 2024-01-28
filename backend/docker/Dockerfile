# BUILD
FROM golang:1.21.6-alpine3.19 AS builder
WORKDIR /src
COPY ./backend .
RUN go build -o formunity_backend main.go

# RUNTIME
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /src/formunity_backend /
EXPOSE 8081
USER nonroot:nonroot
CMD ["/formunity_backend"]