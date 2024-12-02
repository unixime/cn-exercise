FROM library/golang:1.23.1-alpine3.20 AS builder

RUN mkdir -p /go/src/cn-exercise/
COPY go.mod main.go /go/src/cn-exercise/
COPY cmd/* /go/src/cn-exercise/cmd/
COPY docs /go/src/cn-exercise/docs/
COPY internal /go/src/cn-exercise/internal/

WORKDIR /go/src/cn-exercise

RUN go mod tidy

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix cgo -o /go/bin/server cn-exercise


FROM scratch


#ENV OPENAM_DEPLOY_URL=${OPENAM_DEPLOY_URL:-"http://openam:8080/openam"}
COPY --from=builder /go/bin/server /app/server

USER 1001

ENTRYPOINT ["/app/server"]