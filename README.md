```bash
swag init -g cmd/main.go
docker volume create onelab_psql
docker compose up -d
docker image build -t onelab .
```