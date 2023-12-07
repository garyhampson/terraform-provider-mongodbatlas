package streaminstance_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/config"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/testutil/acc"
)

func TestAccStreamRSStreamInstance_basic(t *testing.T) {
	var (
		orgID        = os.Getenv("MONGODB_ATLAS_ORG_ID")
		projectName  = acctest.RandomWithPrefix("test-acc-stream")
		instanceName = acctest.RandomWithPrefix("test-acc-name")
		resourceName = "mongodbatlas_stream_instance.test"
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.PreCheckBasic(t) },
		ProtoV6ProviderFactories: acc.TestAccProviderV6Factories,
		CheckDestroy:             checkDestroyStreamInstance,
		Steps: []resource.TestStep{
			{
				Config: streamInstanceConfig(orgID, projectName, instanceName, region, cloudProvider), // as of now there are no values that can be updated because only one region is supported
				Check:  streamInstanceAttributeChecks(resourceName, orgID, projectName, instanceName, region, cloudProvider),
			},
			{
				ResourceName:      resourceName,
				ImportStateIdFunc: checkStreamInstanceImportStateIDFunc(resourceName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func streamInstanceAttributeChecks(resourceName, orgID, projectName, instanceName, region, cloudProvider string) resource.TestCheckFunc {
	resourceChecks := []resource.TestCheckFunc{
		checkSearchInstanceExists(),
		resource.TestCheckResourceAttrSet(resourceName, "id"),
		resource.TestCheckResourceAttrSet(resourceName, "project_id"),
		resource.TestCheckResourceAttr(resourceName, "instance_name", instanceName),
		resource.TestCheckResourceAttr(resourceName, "data_process_region.region", region),
		resource.TestCheckResourceAttr(resourceName, "data_process_region.cloud_provider", cloudProvider),
		resource.TestCheckResourceAttr(resourceName, "hostnames.#", "1"),
	}
	return resource.ComposeTestCheckFunc(resourceChecks...)
}

func streamInstanceConfig(orgID, projectName, instanceName, region, cloudProvider string) string {
	return fmt.Sprintf(`
		resource "mongodbatlas_project" "test" {
			org_id = %[1]q
			name   = %[2]q
		}

		resource "mongodbatlas_stream_instance" "test" {
			project_id = mongodbatlas_project.test.id
			instance_name = %[3]q
			data_process_region = {
				region = %[4]q
				cloud_provider = %[5]q
			}
		}
	`, orgID, projectName, instanceName, region, cloudProvider)
}

func checkStreamInstanceImportStateIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("%s-%s", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["instance_name"]), nil
	}
}

func checkSearchInstanceExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		connV2 := acc.TestMongoDBClient.(*config.MongoDBClient).AtlasV2
		for _, rs := range s.RootModule().Resources {
			if rs.Type == "mongodbatlas_stream_instance" {
				_, _, err := connV2.StreamsApi.GetStreamInstance(context.Background(), rs.Primary.Attributes["project_id"], rs.Primary.Attributes["instance_name"]).Execute()
				if err != nil {
					return fmt.Errorf("stream instance (%s:%s) does not exist", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["instance_name"])
				}
			}
		}
		return nil
	}
}

func checkDestroyStreamInstance(state *terraform.State) error {
	if projectDestroyedErr := acc.CheckDestroyProject(state); projectDestroyedErr != nil {
		return projectDestroyedErr
	}
	connV2 := acc.TestMongoDBClient.(*config.MongoDBClient).AtlasV2
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "mongodbatlas_stream_instance" {
			_, _, err := connV2.StreamsApi.GetStreamInstance(context.Background(), rs.Primary.Attributes["project_id"], rs.Primary.Attributes["instance_name"]).Execute()
			if err == nil {
				return fmt.Errorf("stream instance (%s:%s) still exists", rs.Primary.Attributes["project_id"], rs.Primary.Attributes["instance_name"])
			}
		}
	}
	return nil
}