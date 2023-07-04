package dataAccess

import (
	"fmt"
	"testing"
)

func TestGetTreeStructure(t *testing.T) {
	data, err := GetTreeStructure()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)
}