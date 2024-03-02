package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/EggsyOnCode/task/cmd"
	"github.com/EggsyOnCode/task/db"
)

func TestDone(t *testing.T) {
	//create mock db
	mockDb := db.GetMockDB()
	//create the root command
	cmd := cmd.TestingDone(mockDb)

	// redirect the stdout to buffer to capture the output
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	//set the flags as args
	args := []string{"1"}
	cmd.SetArgs(args)

	//execute teh root.Execute
	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	expected := "Marked \"1\" as completed.\n"
	fmt.Print(args)
	out := stdout.String()
	if out != expected {
		t.Errorf("Expected output %v \n output received %v", expected, out)
	}

}
