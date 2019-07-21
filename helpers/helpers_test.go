package helpers_test

import (
	"sshaws/helpers"
	"testing"
)

func TestConfiguration(t *testing.T) {
	expRegion, expEnv, expApp, expName := "eu-west-1", "test_env", "test_app", "test_name"
	newConfig := helpers.NewConfiguration("eu-west-1", "test_env", "test_app", "test_name")
	retEnv, retApp, retName, retRegion := helpers.ReturnConfiguration(*newConfig)

	if retRegion != expRegion {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retRegion, expRegion)
	}
	if retEnv != expEnv {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retEnv, expEnv)
	}
	if retApp != expApp {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retApp, expApp)
	}
	if retName != expName {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retName, expName)
	}

}
func TestNewInstance(t *testing.T) {}
