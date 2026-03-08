FROM golang:1.26-trixie AS build

WORKDIR /app
COPY backend /app/backend
COPY frontend /app/frontend
COPY json /app/json
COPY ddl /app/ddl

WORKDIR /app/backend
RUN go build -o main main.go

FROM node:25-trixie AS run
COPY frontend /frontend
WORKDIR /frontend
RUN npm install

WORKDIR /app
COPY --from=build /app/backend/main /app/backend/main
COPY json /app/json
COPY ddl /app/ddl

CMD ["./main"]

