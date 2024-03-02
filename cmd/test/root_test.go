package test

import (
	"bytes"
	"testing"

	"github.com/EggsyOnCode/task/cmd"
)

func TestRoot(t *testing.T) {
	//create the root command
	cmd := cmd.TestingRoot()

	// redirect the stdout to buffer to capture the output
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	//set the flags as args

	//execute teh root.Execute
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expected := "Task is a cli tool to manage ur tasks"

	if cmd.Short != expected {
		t.Errorf("Expected output %v output received %v", expected, cmd.Short)
	}

}
