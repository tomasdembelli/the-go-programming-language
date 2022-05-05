package scratch

import (
	"reflect"
	"testing"
)
 
func TestTypes(t *testing.T) {

	a := 5
	if reflect.TypeOf(a).String() != "int" {
		t.Errorf("expected int, got %v", reflect.TypeOf(a).String())
	}

	var b float64 = 5
	if reflect.TypeOf(b).String() != "float64" {
		t.Errorf("expected int, got %v", reflect.TypeOf(b).String())
	}
}