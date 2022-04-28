package prose

import "testing"

func TestTwoElements(t *testing.T) {
	//t.Error("no test written yet")
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElement(t *testing.T) {
	//t.Error("no test here either")
	list := []string{"apple", "orange", "pear"}
	if JoinWithCommas(list) != "apple and orange, and pear" {
		t.Error("didn't match expected value")
	}
}
