# syntax=docker/dockerfile:1.7

# Etap budowania aplikacji.
# Zmiennych TARGETOS i TARGETARCH dostarcza BuildKit/buildx podczas budowy wieloarchitekturowej.
FROM --platform=$BUILDPLATFORM golang:1.25.10-alpine3.23 AS build

WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY src ./src
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/app ./src

# Obraz wynikowy jest minimalny. Nie zawiera powłoki, menedżera pakietów ani zbędnych bibliotek.
FROM scratch

COPY --from=build /out/app /app
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/app"]
