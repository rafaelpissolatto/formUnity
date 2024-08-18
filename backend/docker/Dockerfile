# BUILD
FROM golang:1.21.6-alpine3.19 AS builder
WORKDIR /src
COPY ./backend .
RUN go build -o formunity_backend main.go

# RUNTIME
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder --chown=65532:65532 /src/formunity_backend /
EXPOSE 8081
USER 65532:65532
CMD ["/formunity_backend"]