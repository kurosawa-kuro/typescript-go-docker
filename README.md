# react-ts-go-docker
React TypeScript Go in Docker 

docker-compose down


docker-compose down -v
docker-compose up -d --build backend
docker logs -f backend


docker compose exec backend go test ./...

http://localhost:8000/swagger/index.html

http://localhost:3000

Pgadmin
http://localhost:5050/

docker-compose down -v&& docker-compose up --build

Nextjsは作り直し


# コンテナを再ビルドして起動
make rebuild

# コンテナを停止
make down

# コンテナのログを表示
make logs

# 実行中のコンテナを表示
make ps

# コンテナ、ボリューム、ネットワークを完全に削除
make clean

# 利用可能なコマンドの一覧を表示
make help