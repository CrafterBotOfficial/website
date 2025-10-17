# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.1

FROM --platform=$BUILDPLATFORM alpine:latest AS build
RUN apk update
RUN apk add --no-cache go gcc musl-dev
WORKDIR /src
ARG TARGETPLATFORM
COPY . .

ENV CGO_ENABLED=1
RUN go build -o /bin/server main.go

ARG TARGETARCH

RUN apk add --no-cache git openssh-client
RUN --mount=type=ssh \
    mkdir -p ~/.ssh && \
    ssh-keyscan github.com >> ~/.ssh/known_hosts && \
    git clone git@github.com:CrafterBotOfficial/website-info.git --depth 1 website-info/

COPY . .

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk apk --update add ca-certificates tzdata && update-ca-certificates

RUN apk add --no-cache git openssh-client

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/home/website" \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    website 
USER website
WORKDIR /home/website

COPY --from=build /src/. .
COPY --from=build /bin/server /bin/

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]
