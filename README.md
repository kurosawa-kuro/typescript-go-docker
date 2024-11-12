# react-ts-go-docker
React TypeScript Go in Docker 

docker-compose down

# frontendディレクトリをクリーンアップ
Remove-Item -Recurse -Force frontend -ErrorAction SilentlyContinue
New-Item -ItemType Directory -Path frontend
icacls "frontend" /grant Everyone:F