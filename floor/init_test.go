package floor 

import (
	"testing"
	"reflect"
)

func Test_file(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	if reflect.TypeOf(tab).Kind() == reflect.Slice {
		elemType := reflect.TypeOf(tab).Elem()
		if elemType.Kind() != reflect.Slice || elemType.Elem().Kind() != reflect.Int {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func Test_empty_alpha(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	for i:=0; i<len(tab); i++ {
		if len(tab[i]) == 0 {
			t.Fail()
		}
	}
}

func Test_line_size(t *testing.T) {
	tab := readFloorFromFile("../floor-files/test/test_line_size")
	size := len(tab[0])
	for i:=0; i<len(tab); i++ {
		if len(tab[i]) != size {
			t.Fail()
		}
	} 
}