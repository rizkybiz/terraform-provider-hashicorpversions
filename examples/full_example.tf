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

variable "arch" {
  description = "the architecture of the HashiCorp product build"
  default = "amd64"
}

variable "os" {
  description = "the operating system you're deploying the HashiCorp product to"
  default = "linux"
}

locals {
  version = data.hashicorpversions_product.product_version.version
  builds = data.hashicorpversions_product.product_version.builds
  arch_specific_builds = [for build in local.builds : build if build.arch == "${var.arch}"]
}

output "product_version" {
  value = data.hashicorpversions_product.product_version.version
}

output "product_builds" {
  value = data.hashicorpversions_product.product_version.builds
}

output "specific_arch_and_os" {
  value = [for build in local.arch_specific_builds : build if build.os == "${var.os}" ]
}

output "specific_arc_and_os_url" {
  value = element([for build in local.arch_specific_builds : build.url if build.os == "${var.os}" ],0)
}