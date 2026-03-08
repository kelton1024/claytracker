FROM golang:1.26-trixie AS build

WORKDIR /app
COPY backend /app/backend
COPY frontend /app/frontend
COPY json /app/json
COPY ddl /app/ddl

WORKDIR /app/backend
RUN go build -o /app/backend/main .
RUN ls -l /app/backend
RUN export PATH=$PATH:/app/backend
RUN echo $PATH
CMD echo 'hello' && ls /app/backend && /app/backend/main

#FROM debian:trixie AS run
#WORKDIR /app
#COPY --from=build /app/backend/main /app/backend/main
#COPY json /app/json
#COPY ddl /app/ddl
#
#WORKDIR /app/backend
#ENTRYPOINT ["/app/backend/main"]

