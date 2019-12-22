FROM golang:1.13 AS build

WORKDIR /app

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o /go/bin/client helloworld/client/main.go

RUN go build -o /go/bin/server helloworld/server/main.go

FROM golang:1.13 AS client

COPY --from=build /go/bin/client /go/bin/client

ENTRYPOINT [ "/go/bin/client" ]

FROM golang:1.13 AS server

COPY --from=build /go/bin/server /go/bin/server

ENTRYPOINT [ "/go/bin/server" ]
