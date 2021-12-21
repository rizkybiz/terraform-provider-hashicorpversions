package hashicorpversions

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVersion_Consulbasic(t *testing.T) {
	datasourceName := "data.hashicorp_products_version.consul_test"

	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVersion_ConsulBaseConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(datasourceName, "version", regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`)),
				),
			},
		},
	})
}

func testAccDataSourceVersion_ConsulBaseConfig() string {
	return fmt.Sprintf(`
data "hashicorp_versions_product" "consul_test" {
  name = "consul"
}
`)
}
