.PHONY: up down rebuild logs ps clean help

# デフォルトのターゲット
.DEFAULT_GOAL := help

# 変数定義
DC := docker-compose

# コマンド
up: ## コンテナを起動
	$(DC) up -d

down: ## コンテナを停止
	$(DC) down

rebuild: ## コンテナを再ビルドして起動
	$(DC) down -v
	$(DC) up --build

rebuild-frontend: ## フロントエンドのコンテナを再ビルドして起動
	$(DC) down -v
	$(DC) up --build frontend

rebuild-backend: ## バックエンドのコンテナを再ビルドして起動
	$(DC) down -v
	$(DC) up --build backend

# docker compose exec frontend sh でコンテナ内に入る
exec-frontend: ## フロントエンドのコンテナに入る
	$(DC) exec frontend sh

exec-backend: ## バックエンドのコンテナに入る
	$(DC) exec backend sh

exec-db: ## DBのコンテナに入る
	$(DC) exec db sh

logs: ## コンテナのログを表示
	$(DC) logs -f

ps: ## 実行中のコンテナを表示
	$(DC) ps

clean: ## コンテナ、ボリューム、ネットワークを削除
	$(DC) down -v --rmi all --remove-orphans

test-frontend: ## フロントエンドのテストを実行
	$(DC) exec frontend npm test

test-frontend-watch: ## フロントエンドのテストをウォッチモードで実行
	$(DC) exec frontend npm run test:watch

# ヘルプコマンド
help: ## このヘルプを表示
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'