package dataAccess

import "testing"

func TestGetStuName(t *testing.T) {
	res, err := GetStuName("2019213860")
	if err != nil {
		t.Fatal(err)
	}

	if res != "吴治霖" {
		t.Fatal("get wrong name")
	}
}
