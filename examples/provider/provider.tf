terraform {
  required_providers {
    hashicorpversions = {
      source  = "terraform.example.com/local/hashicorpversions"
      version = "0.1.0"
    }
  }
}

provider "hashicorpversions" {}