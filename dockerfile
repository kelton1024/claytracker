FROM golang:1.26-trixie AS build

WORKDIR /app
COPY go /app/go
COPY frontend /app/frontend
COPY glue /app/glue
COPY json /app/json
COPY ddl /app/ddl

WORKDIR /app/go
RUN go build -o main main.go

FROM node:25-trixie AS run
COPY frontend /frontend
WORKDIR /frontend
RUN npm install

WORKDIR /app
COPY --from=build /app/go/main /app/go/main
COPY glue /app/glue
COPY json /app/json
COPY ddl /app/ddl

CMD ["./main"]
