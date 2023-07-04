package dataAccess

import (
	"errors"
	"fmt"
	"testing"
)

func TestGetStudentAccuracyInfo(t *testing.T) {
	_, err := GetStudentAccuracyInfo("", "111111")
	if err == nil {
		t.Fatal("check not trigger")
	} else {
		if !errors.Is(err, ErrStuNotExist) {
			t.Fatal(err)
		}
	}
	accuracy, err := GetStudentAccuracyInfo("", "2019213860")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(accuracy)
}
