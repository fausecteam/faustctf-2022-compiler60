FROM golang:1.18.3 AS build

WORKDIR /app

RUN go install mvdan.cc/garble@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/

# use garbe to obfuscate go binary
RUN CGO_ENABLED=0 GOGARBLE=* garble -tiny -literals build executor/cmd/executor

FROM scratch

COPY --from=build /app/executor /
COPY LICENSE /

EXPOSE 6062

ENTRYPOINT ["/executor"]
