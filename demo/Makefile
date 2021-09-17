all: run_demo

run_demo:
	mkdir -p registry.terraform.io/<github-user-name>/example/1.0.0/darwin_amd64
	go mod tidy
	go build -o ./registry.terraform.io/<github-user-name>/example/1.0.0/darwin_amd64/terraform-provider-example ../main.go
	terraform init -plugin-dir=.
	terraform apply