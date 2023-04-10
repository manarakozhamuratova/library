```bash
swag init -g cmd/main.go
docker volume create onelab_psql
docker compose up -d
docker image build -t onelab .
```

## migrations

```
export POSTGRES_URL='postgres://demo:postgres@localhost:5432/onelab?sslmode=disable'

migrate create -ext sql -dir internal/storage/postgre/migrations -seq create_users

```
