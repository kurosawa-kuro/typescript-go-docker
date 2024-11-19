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


Storybook
http://localhost:6006/

docker-compose down -v&& docker-compose up --build

	// SSH エージェントソケットをマウント
	"mounts": [
		"source=/run/host-services/ssh-auth.sock,target=/run/host-services/ssh-auth.sock,type=bind"
	],

	"remoteEnv": {
		"SSH_AUTH_SOCK": "/run/host-services/ssh-auth.sock"
	}


	docker compose -f .devcontainer/docker-compose.yml build --no-cache