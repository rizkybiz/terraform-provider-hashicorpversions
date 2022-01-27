package hashicorpversions

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProduct_Consulbasic(t *testing.T) {
	datasourceName := "data.hashicorpversions_product.consul_test"

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProduct_ConsulBaseConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "version", regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`)),
					resource.TestCheckResourceAttr(datasourceName, "name", "consul"),
					resource.TestCheckResourceAttr(datasourceName, "builds.#", "11"),
				),
			},
		},
	})
}

func testAccDataSourceProduct_ConsulBaseConfig() string {
	return `
data "hashicorpversions_product" "consul_test" {
  name = "consul"
}
`
}
