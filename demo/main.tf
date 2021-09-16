terraform {
  required_providers {
    custom-provider = {
      source = "registry.terraform.io/zubairhassan652/example"
      version = "~> 1.0"
    }
  }
}

variable "api_key" {
  type = string
}

provider "custom-provider" {
  api_key = var.api_key
#   provider_schema_variable = "any value"
}

data "plugin_data" "local_name" {}


output "some_output" {
  value = data.plugin_data.local_name.data_source_schema_variable
}
