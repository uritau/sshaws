package acceptance_test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestSshAws(t *testing.T) {
	tfOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "terraform/",
		//TODO: check if we need more data
	})
	defer terraform.Destroy(t, tfOpts)
	terraform.InitAndApply(t, tfOpts)

}
