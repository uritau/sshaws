package params_test

import (
	"os"
	"reflect"
	"sshaws/helpers"
	"sshaws/params"
	"testing"
)

func TestValidateReadParamsEmpty(t *testing.T) {
	os.Args = []string{"sshaws"}
	retConfig := params.Read()
	expConfig := *helpers.NewConfiguration("eu-west-1", "*", "*", "*")
	if !reflect.DeepEqual(expConfig, retConfig) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retConfig, expConfig)
	}
}

func TestValidateReadName(t *testing.T) {
	os.Args = []string{"sshaws", "-n", "1234"}
	retConfig := params.Read()
	expConfig := *helpers.NewConfiguration("eu-west-1", "*", "*", "1234")
	if !reflect.DeepEqual(expConfig, retConfig) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retConfig, expConfig)
	}

	os.Args = []string{"sshaws", "--name", "test_1"}
	retConfig = params.Read()
	expConfig = *helpers.NewConfiguration("eu-west-1", "*", "*", "test_1")
	if !reflect.DeepEqual(expConfig, retConfig) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retConfig, expConfig)
	}
}

func TestValidateMultipleParams(t *testing.T) {
	os.Args = []string{"sshaws", "-n", "1234", "-a", "application_test", "-e", "environment_test"}
	retConfig := params.Read()
	expConfig := *helpers.NewConfiguration("eu-west-1", "environment_test", "application_test", "1234")
	if !reflect.DeepEqual(expConfig, retConfig) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retConfig, expConfig)
	}

	os.Args = []string{"sshaws", "--name", "1234", "--app", "application_test", "--env", "environment_test"}
	retConfig = params.Read()
	expConfig = *helpers.NewConfiguration("eu-west-1", "environment_test", "application_test", "1234")
	if !reflect.DeepEqual(expConfig, retConfig) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retConfig, expConfig)
	}
}
