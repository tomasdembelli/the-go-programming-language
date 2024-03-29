package scratch

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// It is mandatory to start with Test.. in the function name.
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

func f() *int {
	v := 1
	return &v
}

func inc(v *int) {
	*v++
}

func update(s []int, i, with int) {
	s[i] = with
}

func TestPointer(t *testing.T) {

	var pz *int // declaring pz as a pointer to an int value
	if pz != nil {
		t.Errorf("expected nil, but got %v", pz)
	}

	a := 5
	p := &a
	*p = 6
	if a != 6 {
		t.Errorf("expected 6, but got %d", a)
	}

	p2 := &a
	if p != p2 {
		t.Errorf("expected %v == %v, but got false", p, p2)
	}

	c1 := f()
	c2 := f()
	if c1 == c2 {
		t.Errorf("expected different values, but got the same value %v", c1)
	}

	v := 5
	inc(&v)
	if v != 6 {
		t.Errorf("expected 6, got %d", v)
	}

	s := []int{1, 2}
	update(s, 0, 999)
	if s[0] != 999 {
		t.Errorf("expected 999, but got %d", s[0])
	}
}

func TestStrings(t *testing.T) {
	engWord := "hello"

	if len(engWord) != 5 {
		t.Errorf("expected 4, got %d", len(engWord))
	}

	euroSign := "€" // the non-ASCII `€` character takes 3 bytes, hence the length is 3 as opposed to 1 (number of characters)

	if len(euroSign) != 3 {
		t.Errorf("expected 3, got %d", len(euroSign))
	}

	fmt.Println("\xE2\x82\xAC") // "€"

	if euroSign != "\xE2\x82\xAC" {
		t.Errorf("expected %v, but got %v", euroSign, "\xE2\x82\xAC")
	}

	// for i, r := range "€"  {
	// 	fmt.Printf("%d, %q, %d\n",i, r, r)
	// }

	// for i:=0; i < len(euroSign); i++  {
	//     r := euroSign[i]
	// 	fmt.Printf("%d, %q, %d\n",i, r, r)
	// }

	var buf bytes.Buffer
	buf.WriteByte('a')
	buf.WriteByte('n')
	buf.WriteRune('€')
	buf.WriteString("a")
	if buf.String() != "an€a" {
		t.Errorf("expected an€, got %v", buf.String())
	}
}

func TestArray(t *testing.T) {
	a := [...]int{1, 2}

	// length is determined by the number of initializers.
	if len(a) != 2 {
		t.Errorf("expected 2, got %d", len(a))
	}

	q := [...]int{99: 5}

	if len(q) != 100 {
		t.Errorf("expected 100, got %d", len(q))
	}

	months := [...]string{"January", "February", "March", "April"} // type is array

	s1 := months[1:3] // type is slice: len 2, cap 3 (cap starts from 1 - February, but ends at the end of the array - April)

	if cap(s1) != 3 {
		t.Errorf("expected 3, got %d", cap(s1))
	}

	s1[1] = "hello" // this will mutate the underlying array - months[2]

	if months[2] != "hello" {
		t.Errorf("expected hello, got %s", months[2])
	}

	s1 = append(s1, "test")

	if months[3] != s1[2] {

		t.Errorf("expected %s, got %s", s1[2], months[3])
	}

}
