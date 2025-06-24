FROM golang:1.24.4-alpine3.22

COPY ./go.mod ./go.mod
COPY ./main.go ./main.go

CMD ["go", "build", "-o", "/healthcheck/healthcheck"]
