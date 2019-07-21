package main

import (
	"os"
	"reflect"
	"sshaws/helpers"
	"testing"
)

func TestValidateReadParamsEmpty(t *testing.T) {
	os.Args = []string{"sshaws"}
	returnedConfiguration := readParams()
	generatedConfiguration := *helpers.NewConfiguration("eu-west-1", "*", "*", "*")
	if !reflect.DeepEqual(generatedConfiguration, returnedConfiguration) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", returnedConfiguration, generatedConfiguration)
	}

	os.Args = []string{"sshaws", "-n", "1234"}
	returnedConfiguration = readParams()
	generatedConfiguration = *NewConfiguration("eu-west-1", "*", "*", "*")
	if !reflect.DeepEqual(generatedConfiguration, returnedConfiguration) {
		t.Errorf("\nGot:\n %+v\nWant:\n %+v\n", returnedConfiguration, generatedConfiguration)
	}
}
