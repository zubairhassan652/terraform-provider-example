terraform {
  required_providers {
    example = {
      source = "registry.terraform.io/zubairhassan652/example"
      version = "~> 1.0"
    }
  }
}

variable "api_key" {
  type = string
}

provider "example" {
  api_key = var.api_key
  provider_schema_variable = "any value"
}

data "example_data" "local_name" {}


output "some_output" {
  value = [
    data.example_data.local_name.data_source_schema_variable,
  ]
}
