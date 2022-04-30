package prose

import "testing"

// func TestTwoElements(t *testing.T) {
// 	//t.Error("no test written yet")
// 	list := []string{"apple", "orange"}
// 	if JoinWithCommas(list) != "apple and orange" {
// 		t.Error("didn't match expected value")
// 	}
// }

// func TestThreeElement(t *testing.T) {
// 	//t.Error("no test here either")
// 	list := []string{"apple", "orange", "pear"}
// 	if JoinWithCommas(list) != "apple, orange, and pear" {
// 		t.Error("didn't match expected value")
// 	}
// }


func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if  got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want)
	}
}

func TestThreeElement(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want)
	}
}
