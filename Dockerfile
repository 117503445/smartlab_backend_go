FROM golang:1.15 as build
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
FROM alpine as prod
EXPOSE 80
COPY --from=build /go/release/app /app
COPY --from=build /wait /wait
WORKDIR /
ENTRYPOINT /wait && /app