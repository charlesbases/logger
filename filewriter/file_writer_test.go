package filewriter

import "testing"

func TestNewFileWriter(t *testing.T) {
	New(OutputPath("_test.log"), MaxRolls(7))
}
