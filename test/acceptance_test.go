package acceptance_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type instanceDetails struct {
	name         string
	ipAddress    string
	instanceId   string
	instanceType string
}

func TestSshAws(t *testing.T) {
	tfOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "terraform/",
		//TODO: check if we need more data
	})

	defer terraform.Destroy(t, tfOpts)
	terraform.InitAndApply(t, tfOpts)

	t.Run("Existing instance is listed correctly", func(t *testing.T) {
		expectedInstanceName := terraform.Output(t, tfOpts, "instance_name")

		command := shell.Command{
			Command:    "go",
			Args:       []string{"run", "cmd/sshaws/main.go", expectedInstanceName},
			WorkingDir: "../",
		}
		output, _ := shell.RunCommandAndGetOutputE(t, command)

		expectedInstanceID := terraform.Output(t, tfOpts, "instance_id")
		expectedIpAddress := terraform.Output(t, tfOpts, "instance_ip_address")
		expectedInstanceType := terraform.Output(t, tfOpts, "instance_type")

		instance := searchInstanceInOutput(output, expectedInstanceName)
		assert.NotNil(t, instance)
		assert.Equal(t, expectedInstanceID, instance.instanceId)
		assert.Equal(t, expectedIpAddress, instance.ipAddress)
		assert.Equal(t, expectedInstanceType, instance.instanceType)
	})

	t.Run("Listing a nonexisting instance", func(t *testing.T) {
		command := shell.Command{
			Command:    "go",
			Args:       []string{"run", "cmd/sshaws/main.go", "this-name-should-never-exist"},
			WorkingDir: "../",
		}
		output, _ := shell.RunCommandAndGetOutputE(t, command)
		expectedOutput := "\nName: this-name-should-never-exist   Region: eu-west-1\n---------------------------------------------------------\n\nThere are no instances matching your request."
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("Show instance with -silent flag", func(t *testing.T) {
		expectedInstanceName := terraform.Output(t, tfOpts, "instance_name")

		command := shell.Command{
			Command:    "go",
			Args:       []string{"run", "cmd/sshaws/main.go", "-silent", expectedInstanceName},
			WorkingDir: "../",
		}
		output, _ := shell.RunCommandAndGetOutputE(t, command)

		expectedIpAddress := terraform.Output(t, tfOpts, "instance_ip_address")
		assert.Equal(t, expectedIpAddress, output)
	})
}

func searchInstanceInOutput(output string, searchName string) *instanceDetails {
	reg_expr := fmt.Sprintf("[0-9]+ %s [0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}", searchName)
	r, _ := regexp.Compile(reg_expr)

	for _, line := range strings.Split(output, "\n") {
		lineFields := strings.Split(strings.TrimSpace(line), " ")
		if r.MatchString(line) {
			return &instanceDetails{
				name:         strings.TrimSpace(lineFields[1]),
				ipAddress:    strings.TrimSpace(lineFields[2]),
				instanceId:   strings.TrimSpace(lineFields[3]),
				instanceType: strings.TrimSpace(lineFields[4]),
			}
		}
	}

	return nil
}
