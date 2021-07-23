FROM golang:alpine AS builder
ARG TAGGED 
LABEL maintainer="dedy styawan <dedy.styawan@gmail.com>"
LABEL tagged=$TAGGED
WORKDIR /src
RUN go get github.com/cespare/reflex
# Allow for caching
COPY go.mod go.sum ./
RUN go mod download

COPY / .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -ldflags "-s -w " \
  -o binary ./cmd

FROM alpine
LABEL maintainer="dedy styawan <dedy.styawan@gmail.com>"
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/binary /binary
COPY --from=builder /src/app.env /src/app.env

EXPOSE 8000

CMD [ "/binary" ]