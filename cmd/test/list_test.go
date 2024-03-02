package test

import (
	"bytes"
	"testing"

	"github.com/EggsyOnCode/task/cmd"
	"github.com/EggsyOnCode/task/db"
)

func TestList(t *testing.T) {
	//create mock db
	mockDb := db.GetMockDB()
	//create the root command
	cmd := cmd.TestingList(mockDb)

	// redirect the stdout to buffer to capture the output
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	//set the flags as args

	//execute teh root.Execute
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expected := "number : 1 \t task: Task 1 \nnumber : 2 \t task: Task 2 \n"

	out := stdout.String()
	if out != expected {
		t.Errorf("Expected output %v \n output received %v", expected, out)
	}

}
