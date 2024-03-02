package test

import (
	"bytes"
	"testing"

	"github.com/EggsyOnCode/task/cmd"
	"github.com/EggsyOnCode/task/db"
)

func TestAdd(t *testing.T) {
	//create mock db
	mockDb := db.GetMockDB()
	//create the root command
	cmd := cmd.TestingAdd(mockDb)

	// redirect the stdout to buffer to capture the output
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	//set the flags as args
	args := []string{"do", "test", "prep"}
	cmd.SetArgs(args)

	//execute teh root.Execute
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expected := "the task \"do test prep\" with id 1 has been added to the task list"
	out := stdout.String()
	if out != expected {
		t.Errorf("Expected output %v output received %v", expected, out)
	}

}
