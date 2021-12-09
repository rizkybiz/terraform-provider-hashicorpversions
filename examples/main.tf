terraform {
  required_providers {
    hashicorpversions = {
      version = "0.0.1"
      source = "local/rizkybiz/hashicorpversions"
    }
  }
}

provider "hashicorpversions" {}

data "version" "consul_version" {
  product = "consul"
}

output "consul_version" {
  value = data.version.consul_version.version
}