package golang

import "testing"

func TestGetSize(t *testing.T) {
	if size, _ := getSize("../a.txt"); size <= 0 {
		t.Error("a.txt has size > 0")
	}
}

func TestLocateNameColum(t *testing.T) {
	if colIdx := locateNameColumn("#a,b,name,d"); colIdx != 2 {
		t.Errorf("expected name col index to be 2, got %v", colIdx)
	}
}

func TestDecideReadingOrderForJoin(t *testing.T) {
	order, _ := decideReadingOrderForJoin("../a.txt", "../b.txt")
	if order[0] != "../b.txt" {
		t.Errorf("expected smaller file: %v to read first but got %v", "../b.txt", order[0])
	}
}
func TestPrepareDataMap(t *testing.T) {
	if dataMap, err := prepareDataMap("../a.txt"); err != nil {
		t.Error(err)
	} else {
		if len(dataMap) != 3 {
			t.Errorf("expected datamap size to be 3, got %v", len(dataMap))
		}
		if dataMap["vishnu"] != "1" {
			t.Errorf("expected datamap value for key: vishnu to be 1, got %v", dataMap["vishnu"])
		}
	}
}
