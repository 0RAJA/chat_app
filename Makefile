.PHONY: test docker_net migrate_up migrate_up1 migrate_down migrate_down1 sqlc format swag pull init server_init docker_build docker_run run run_back
test: # 运行所有的测试程序
	go test -v -cover ./... -count=1
docker_net: # 创建docker网络
	docker network create chat_net
postgres_zr_init: # 初始化postgres数据库
	docker run --name chat_postgres_zr --network chat_net -v chat_postgres_zr_data:/var/lib/postgresql/data -v 项目路径/config/postgres/my_postgres.conf:/etc/postgresql/postgresql.conf -p 5432:5432 -e ALLOW_IP_RANGE=0.0.0.0/0 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=chat -d chenxinaz/zhparser -c 'config_file=/etc/postgresql/postgresql.conf'
redis_init: # redis初始化
	docker run --name chat_redis_62 --network chat_net --privileged=true -p 7963:7963 -v chat_redis_data:/redis/data -v 项目路径/config/redis:/etc/redis -d redis:6.2 redis-server /etc/redis/redis.conf
sqlc_install:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.13.0
migrate_install: # 安装migrate
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1
migrate_init_db: # 初始化数据库
	migrate create -ext sql -dir src/dao/postgres/migration -seq init_schema
migrate_up: # 向上迁移数据库
	migrate -path src/dao/postgres/migration -database "postgresql://root:123456@localhost:5432/chat?sslmode=disable" -verbose up
migrate_up1: # 向上迁移一级数据库
	migrate -path src/dao/postgres/migration -database "postgresql://root:123456@localhost:5432/chat?sslmode=disable" -verbose up 1
migrate_down: # 向下迁移数据库
	migrate -path src/dao/postgres/migration -database "postgresql://root:123456@localhost:5432/chat?sslmode=disable" -verbose down
migrate_down1: # 向下迁移一级数据库
	migrate -path src/dao/postgres/migration -database "postgresql://root:123456@localhost:5432/chat?sslmode=disable" -verbose down 1
sqlc: # sqlc生成go代码
	sqlc generate
goimports_install: # goimports安装
	go install golang.org/x/tools/cmd/goimports@latest
format: # 格式化并检查代码
	goimports -w . && gofmt -w . && golangci-lint run
golang-cli_install: # 安装golang-cli工具，用于静态检查代码质量
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
swag_install: # 安装swag工具，用于生成swagger文档
	go install github.com/swaggo/swag/cmd/swag@v1.8.0
swag: # swag生成文档
	swag fmt --d src/ && swag init -d src/
pull: # 拉取并变基代码
	git fetch origin master && git rebase origin/master
init: migrate_install goimports_install golang-cli_install swag_install # 安装工具包
docker_build: # 构建docker镜像
	docker build -t chat:app .
docker_run: # docker运行镜像
	docker run -d --name chat_app --network chat_net -p 8080:8080 chat:app
build:
	go build -ldflags="-s -w" -o bin/chat src/main.go
run: build# 运行server
	./bin/chat
run_back: build# 后台运行
	nohup ./bin/chat > nohup.out &
