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

func TestGetAllStudentAccuracyInfo(t *testing.T) {
	allAccuracy, err := GetAllStudentAccuracyInfo()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(allAccuracy); i++ {
		fmt.Println(allAccuracy[i])
	}

	fmt.Println(len(allAccuracy))
}

func TestGetClassKnowledgeAccuracyInfo(t *testing.T) {
	accuracy, err := GetClassKnowledgeAccuracyInfo("1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(accuracy)
}
