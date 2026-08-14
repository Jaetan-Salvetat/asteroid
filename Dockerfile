FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=js GOARCH=wasm go build -o /out/game.wasm . \
 && cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /out/

FROM caddy:2-alpine

COPY --from=builder /out/ /srv/
COPY web/index.html /srv/
COPY Caddyfile /etc/caddy/Caddyfile
