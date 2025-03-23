package structs

import "testing"

func TestInitializationByNew(t *testing.T){
	p := InitializationByNew()

	if p.Name != "John" {
		t.Error("Name should be 'John'")
	}

	if p.LastName != "Wick" {
		t.Error("Last name should be 'Wick'")
	}

	if p.Age != 41 {
		t.Error("Age should be 41")
	}
}

