terraform {
  required_providers {
    hashicorpversions = {
      version = "0.0.1"
      source = "terraform.example.com/local/hashicorpversions"
    }
  }
}

provider "hashicorpversions" {}

module "version" {
  source = "./version"
  product_name = "consul"
}

output "version" {
  value = module.version.product_version
}