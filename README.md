## Go to the project directory
```
cd secure-file-vault
```

## Run docker compose to initialize MySQL, Redis and minIO
```
docker compose up -d
```

## Make migrations
```
make upse
```

## Update config filename
```
mv config/production.yaml.example config/production.yaml
```

## Build
```
go build -o secure-file-value cmd/server/main.go
```

## Create bucket in minIO
```
Create bucket name "uploaded-files" by using minIO GUI :)
```

## Run
```
./secure-file-value
```