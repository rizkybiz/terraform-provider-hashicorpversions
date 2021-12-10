terraform {
  required_providers {
    hashicorpversions = {
      version = "0.0.1"
      source = "terraform.example.com/local/hashicorpversions"
    }
  }
}

variable "product_name" {
  type = string
}

# variable "os" {
#   type = string
# }

# variable "arch" {
#   type = string
# }

data "hashicorpversions_version" "product_version" {
  product = var.product_name
}

output "product_version" {
  value = data.hashicorpversions_version.product_version.version
}

output "product_builds" {
  value = data.hashicorpversions_version.product_version.builds
}