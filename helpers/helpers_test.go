package helpers_test

import (
	"testing"

	"sshaws/helpers"
)

func TestConfiguration(t *testing.T) {
	expRegion, expEnv, expApp, expName, expUser, expSilent, expSSH := "eu-west-1", "test_env", "test_app", "test_name", "test_user", false, false
	newConfig := helpers.NewConfiguration("eu-west-1", "test_env", "test_app", "test_name", "test_user", false, false)
	retEnv, retApp, retName, retRegion, retUser, retSilent, retSSH := helpers.ReturnConfiguration(*newConfig)

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
	if retUser != expUser {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retUser, expUser)
	}
	if retSilent != expSilent {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retSilent, expSilent)
	}
	if retSSH != expSSH {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", retSSH, expSSH)
	}
}
func TestNewInstance(t *testing.T) {}
