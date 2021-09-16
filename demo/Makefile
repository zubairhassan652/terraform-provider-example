all: build

build:
    mkdir -p registry.terraform.io/zubairhassan652/example/1.0.0/darwin_amd64
	go mod tidy
	go build -o ./registry.terraform.io/zubairhassan652/example/1.0.0/darwin_amd64/terraform-provider-example ../main.go
	terraform init -plugin-dir=.