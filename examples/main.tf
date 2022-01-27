terraform {
  required_providers {
    hashicorpversions = {
      source  = "terraform.example.com/local/hashicorpversions"
      version = "0.1.0"
    }
  }
}

provider "hashicorpversions" {}

data "hashicorpversions_product" "product_version" {
  name = "consul"
}

output "product_version" {
  value = data.hashicorpversions_product.product_version.version
}

output "product_builds" {
  value = data.hashicorpversions_product.product_version.builds
}