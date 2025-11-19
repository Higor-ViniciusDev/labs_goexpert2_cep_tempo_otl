FROM golang:1.24-alpine as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o servicoB ./cmd/webserver/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/servicoB .
EXPOSE 8181
ENTRYPOINT [ "./servicoB" ]