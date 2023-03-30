# AInsight-Server
- Web Service
- Data API Service
- ***Not Include AI Service***

## Build
- `make build`
- or `go build -o bin`
- output folder: bin

## Execute
- `make execute`
- or `cd ./bin; ./AInsight-Server.exe`

## .env Example
```env
# Web 服務
WEB_HOST=
WEB_PORT=80
WEB_PORT_TLS=443
WEB_FOLDER=web

# DB 資料庫
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_DATABASE=database

# AI 伺服器
AI_PORT=:9000

# JWT Key
SECRET=0123456789ABCDEFG
```