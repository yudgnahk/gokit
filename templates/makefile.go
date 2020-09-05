package templates

const MakefileTemplate = `prepare:
	@go get -u github.com/beego/bee
	@go get -u github.com/swaggo/swag
	@go mod download

migration: ## Generate new migration
	@read -p "Enter migration name:" migration_name; \
	cd $(PWD); bee generate migration $${migration_name}

migrate: ## Run database migrations
	@cd $(PWD) && bee migrate --conn="$(MYSQL_USER):$(MYSQL_PASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DB)?charset=utf8"

run:
	@cd $(PWD)/cmd && go run *.go

build:
	@cd $(PWD)/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

gen_docs:
	@cp $(PWD)/cmd/main.go $(PWD)
	@cd $(PWD) && swag init
	@rm -rf $(PWD)/main.go
	@rm -rf $(PWD)/cmd/docs
	@mv $(PWD)/docs $(PWD)/cmd
`
