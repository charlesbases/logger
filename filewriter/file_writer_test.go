package filewriter

import "testing"

func TestNewFileWriter(t *testing.T) {
	New(FilePath("_test.log"), TTL(7))
}
