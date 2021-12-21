package hashicorpversions

import (
	"fmt"
	"regexp"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVersion_Consulbasic(t *testing.T) {
	datasourceName := "data.hashicorpversions_version.test"
	regexStr := `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`

	resource.Test(t, resource.TestCase{
		PreCheck: func() { sdkacctest.PreCheck(t) },
		// ErrorCheck: acctest.ErrorCheck(t, connect.EndpointsID),
		Providers: sdkacctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVersion_ConsulBaseConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "version", regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`)),
					resource.TestCheckResourceAttr(datasourceName, "lex_bot", resourceName, "lex_bot"),
				),
			},
		},
	})
}

func testAccDataSourceVersion_ConsulBaseConfig() string {
	return fmt.Sprintf(`
data "hashicorpversions_version" "product_version" {
  product = "consul"
}
`)
}

func testAccDataSourceVersion_VersionBasic(rProduct string) string {
	return fmt.Sprintf(testAccDataSourceVersion_ConsulBaseConfig() + `
output "product_version" {
  value = data.hashicorpversions_version.product_version.version
}
`)
}
