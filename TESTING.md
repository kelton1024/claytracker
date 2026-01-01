# cURL Commands

Insert
```bash
curl -X POST localhost:8080/insert -H "Content-Type: application/json" -d '{"key": "MyKey", "value": "Test"}'
```

Query
```bash
curl -X POST localhost:8080/query -H "Content-Type: application/json" -d '{"key": "test"}'
```

# DDL Commands
Run Postgres Container
```bash
podman run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword postgres:14.19
```

Generate DDL File
```bash
cd ddl
go run main.go generate
```

Create database
```bash
cd ddl
go run main.go create
```

It should be noted that main.go will use "localhost:5432" as the host. 