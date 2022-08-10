FROM golang:1.14-alpine as builder

RUN apk update && apk add --no-cache make && apk --update add ca-certificates && apk add --no-cache git

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /math_server

COPY go.mod go.sum /math_server/

RUN go mod download

COPY . /math_server/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable
COPY --from=builder /math_server/math_server /
# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["/math_server"]