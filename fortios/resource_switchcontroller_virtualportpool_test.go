// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSwitchControllerVirtualPortPool_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSwitchControllerVirtualPortPool_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSwitchControllerVirtualPortPoolConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSwitchControllerVirtualPortPoolExists("fortios_switchcontroller_virtualportpool.trname"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_virtualportpool.trname", "description", "virtualport"),
					resource.TestCheckResourceAttr("fortios_switchcontroller_virtualportpool.trname", "name", rname),
				),
			},
		},
	})
}

func testAccCheckFortiOSSwitchControllerVirtualPortPoolExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SwitchControllerVirtualPortPool: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SwitchControllerVirtualPortPool is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerVirtualPortPool(i)

		if err != nil {
			return fmt.Errorf("Error reading SwitchControllerVirtualPortPool: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SwitchControllerVirtualPortPool: %s", n)
		}

		return nil
	}
}

func testAccCheckSwitchControllerVirtualPortPoolDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_switchcontroller_virtualportpool" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSwitchControllerVirtualPortPool(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SwitchControllerVirtualPortPool %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSwitchControllerVirtualPortPoolConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_switchcontroller_virtualportpool" "trname" {
  description = "virtualport"
  name        = "%[1]s"
}
`, name)
}
