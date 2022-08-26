FROM heroiclabs/nakama-pluginbuilder:3.13.1 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend

COPY . .

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.13.1

COPY --from=go-builder /backend/backend.so /nakama/data/modules/
COPY nakama-config.yaml /nakama/data/local.yml