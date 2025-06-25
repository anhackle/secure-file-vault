## How to use
### Go to the project directory
```
cd secure-file-vault
```

### Run docker compose to initialize MySQL, Redis and minIO
```
docker compose up -d
```

### Make migrations
```
make upse
```

### Update config filename
```
mv config/production.yaml.example config/production.yaml
```

### Build
```
go build -o secure-file-value cmd/server/main.go
```

### Create bucket in minIO
```
Create bucket name "uploaded-files" by using minIO GUI :)
```

### Run
```
./secure-file-value
```

## 📝 Tech Stack
- **Language**: Go 1.20+
- **Cache**: Redis
- **Queue**: Apache Kafka (KRaft mode)
- **Storage**: MinIO
- **Database**: MySQL 5.7
- **Migration**: goose
- **Containerization**: Docker + Docker Compose