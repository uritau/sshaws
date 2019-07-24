package helpers_test

import (
	"sshaws/helpers"
	"testing"
)

func TestConfiguration(t *testing.T) {
	expRegion, expEnv, expApp, expName, expUser := "eu-west-1", "test_env", "test_app", "test_name", "test_user"
	newConfig := helpers.NewConfiguration("eu-west-1", "test_env", "test_app", "test_name", "test_user")
	retEnv, retApp, retName, retRegion, retUser := helpers.ReturnConfiguration(*newConfig)

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
	if retName != expName {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retUser, expUser)
	}
}
func TestNewInstance(t *testing.T) {}
