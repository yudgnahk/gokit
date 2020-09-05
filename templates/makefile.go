package templates

const MakefileTemplate = `prepare:
	@go get -u github.com/pressly/goose
	@go get -u github.com/swaggo/swag
	@go mod download

migration:
	@cd $(PWD)/migrations && read -p "Enter migration name: " migration_name; \
    goose create $${migration_name} sql

migrate:
	@goose up

run:
	@cd $(PWD)/cmd && go run *.go

build:
	@cd $(PWD)/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

gen_docs:
	@cp $(PWD)/cmd/main.go $(PWD)/src
	@cd $(PWD) && swag init
	@rm -rf $(PWD)/main.go
	@rm -rf $(PWD)/cmd/docs
	@mv $(PWD)/docs $(PWD)/cmd/mobile
`
