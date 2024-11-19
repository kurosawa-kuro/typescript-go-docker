.PHONY: up down rebuild logs ps clean help test-frontend test-frontend-watch test-backend exec-frontend exec-backend exec-db

# デフォルトのターゲット
.DEFAULT_GOAL := help

# 変数定義
DC := docker-compose

###################
# 基本コマンド
###################
up: ## コンテナを起動
	$(DC) up -d

down: ## コンテナを停止
	$(DC) down

ps: ## 実行中のコンテナを表示
	$(DC) ps

logs: ## コンテナのログを表示
	$(DC) logs -f

clean:
	-docker-compose down -v --rmi all --remove-orphans 2>/dev/null || true
	-docker system prune -a -f

###################
# ビルド関連
###################
rebuild: down ## コンテナを再ビルドして起動
	$(DC) up --build

rebuild-frontend: down ## フロントエンドのコンテナを再ビルドして起動
	$(DC) up --build frontend

rebuild-backend: down ## バックエンドのコンテナを再ビルドして起動
	$(DC) up --build backend

###################
# コンテナ接続
###################
exec-frontend: ## フロントエンドのコンテナに入る
	$(DC) exec frontend sh

exec-backend: ## バックエンドのコンテナに入る
	$(DC) exec backend sh

exec-db: ## DBのコンテナに入る
	$(DC) exec db sh

###################
# テスト関連
###################
test-frontend: ## フロントエンドのテストを実行
	$(DC) exec frontend npm test

test-frontend-watch: ## フロントエンドのテストをウォッチモードで実行
	$(DC) exec frontend npm run test:watch

test-backend: ## バックエンドのテストを実行
	$(DC) exec backend go test ./...

###################
# ヘルプ
###################
help: ## このヘルプを表示
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'