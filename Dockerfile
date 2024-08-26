FROM golang:1.23-alpine3.19 as BASE

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/
COPY migrations/ migrations/
COPY pkg/ pkg/


FROM base as api
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/*.go

FROM base as migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o migration ./cmd/migration/*.go


FROM gcr.io/distroless/static-debian12:latest AS final

ENV TZ=Asia/Jakarta

COPY --from=api /build/api /api
COPY --from=migrate /build/migration /migration
COPY --from=base /build/migrations/*.sql /data/migrations/