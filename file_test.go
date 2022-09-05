package gk

import "testing"

func TestStringExistsInFile(t *testing.T) {
	exists, err := StringExistsInFile("./test.txt", "123")
	if err != nil {
		t.Error(err)
		return
	}
	if !exists {
		t.Error("not exists")
	}
}

func TestGetLastLineInFile(t *testing.T) {
	lastLine, err := GetLastLineInFile("./test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(lastLine)
}

func TestInsertOneLineToFile(t *testing.T) {
	if exists, err := StringExistsInFile("./test.txt", "456"); !exists || err != nil {
		err := InsertOneLineToFile("./test.txt", "456", "123")
		if err != nil {
			t.Error(err)
		}
	}
}
