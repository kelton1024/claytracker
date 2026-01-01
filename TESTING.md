# cURL Commands

Insert
```bash
curl -X POST localhost:8080/insert -H "Content-Type: application/json" -d '{"key": "MyKey", "value": "Test"}'
```

Query
```bash
curl -X POST localhost:8080/query -H "Content-Type: application/json" -d '{"key": "test"}'
```

Run Postgres Container
```bash
podman run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword postgres:14.19
```