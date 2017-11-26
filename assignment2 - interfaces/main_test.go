package main

import "testing"

func TestTriangleGetArea(t *testing.T) {
	tr := triangle{
		height: 2,
		base:   4,
	}

	if tr.getArea() != 4 {
		t.Errorf("Expected triangle area to be 4, but got: %v", tr.getArea())
	}

	tr.height = 10
	tr.base = 20

	if tr.getArea() != 100 {
		t.Errorf("Expected triangle area to be 100, but got: %v", tr.getArea())
	}
}

func TestRectangleGetArea(t *testing.T) {
	r := rectangle{
		sideLength: 2,
	}

	if r.getArea() != 4 {
		t.Errorf("Expected rectangle area to be 4, but got: %v", r.getArea())
	}

	r.sideLength = 10

	if r.getArea() != 100 {
		t.Errorf("Expected rectangle area to be 100, but got: %v", r.getArea())
	}
}
