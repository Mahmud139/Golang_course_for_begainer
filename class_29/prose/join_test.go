package prose

import (
	//"fmt"
	"testing"
)

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


// func TestTwoElements(t *testing.T) {
// 	list := []string{"apple", "orange"}
// 	want := "apple and orange"
// 	got := JoinWithCommas(list)
// 	if  got != want {
// 		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want)
// 	}
// }

// func TestThreeElement(t *testing.T) {
// 	list := []string{"apple", "orange", "pear"}
// 	want := "apple, orange, and pear"
// 	got := JoinWithCommas(list)
// 	if got != want {
// 		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want)
// 	}
// }

// func TestOneElement(t *testing.T) {
// 	list := []string{"apple"}
// 	want := "apple"
// 	got := JoinWithCommas(list)
// 	if got != want {
// 		t.Error(errorString(list, got, want))
// 	}
// }

// func TestTwoElements(t *testing.T) {
// 	list := []string{"apple", "orange"}
// 	want := "apple and orange"
// 	got := JoinWithCommas(list)
// 	if  got != want {
// 		t.Error(errorString(list, got, want))
// 	}
// }

// func TestThreeElements(t *testing.T) {
// 	list := []string{"apple", "orange", "pear"}
// 	want := "apple, orange, and pear"
// 	got := JoinWithCommas(list)
// 	if got != want {
// 		t.Error(errorString(list, got, want))
// 	}
// }

// func errorString(list []string, got string, want string) string {
// 	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want)
// }

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData {
		{list: []string{}, want: ""},
		{list:[]string{"apple"},want:"apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "banana"}, want: "apple, orange, and banana"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",test.list, got, test.want)
		}
	}
}
