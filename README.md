# terraform-provider-hashicorpversions

The purpose of this Terraform provider is to get the latest semantic version
of any of the suite of HashiCorp tools.

## Build Provider

Run the following command to build the provider:
```shell
$ go build -o terraform-provider-hashicorpversions
```

## Local Release Build

```shell
$ go install github.com/goreleaser/goreleaser@latest
```

```shell
$ make release
```

You will find the releases in the `/dist` directory. You will need to rename the provider binary to `terraform-provider-hashicorpversions` and move the binary into `~/.terraform.d/plugins/terraform.example.com/local/hashicorpversions/0.0.1/darwin_amd64`

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```