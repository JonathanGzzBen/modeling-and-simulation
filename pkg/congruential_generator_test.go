package pkg

import "testing"

func TestX1(t *testing.T) {
	cases := []struct {
		a         int
		x0        int
		c         int
		m         int
		operation generationOperation
		expected  int
	}{{8, 2, 3, 10, Addition, 9},
		{8, 9, 3, 10, Addition, 5},
		{5, 5, 0, 32, Multiplication, 25},
		{5, 25, 0, 32, Multiplication, 29},
	}
	for _, c := range cases {
		actual := X1(c.a, c.x0, c.c, c.m, c.operation)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}

func TestRectangularNumber(t *testing.T) {
	cases := []struct {
		a         int
		x0        int
		c         int
		m         int
		operation generationOperation
		expected  float64
	}{{8, 2, 3, 10, Addition, 0.9},
		{8, 9, 3, 10, Addition, 0.5},
		{5, 5, 0, 32, Multiplication, 0.78125},
		{5, 25, 0, 32, Multiplication, 0.90625},
	}
	for _, c := range cases {
		actual := RectangularNumber(c.a, c.x0, c.c, c.m, c.operation)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}
