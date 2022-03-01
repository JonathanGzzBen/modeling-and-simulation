package pkg

import "testing"

func TestX1(t *testing.T) {
	cases := []struct {
		a        int
		x0       int
		c        int
		m        int
		expected int
	}{{8, 2, 3, 10, 9},
		{8, 9, 3, 10, 5}}
	for _, c := range cases {
		actual := X1(c.a, c.x0, c.c, c.m)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}

func TestRectangularNumber(t *testing.T) {
	cases := []struct {
		a        int
		x0       int
		c        int
		m        int
		expected float64
	}{{8, 2, 3, 10, 0.9},
		{8, 9, 3, 10, 0.5}}
	for _, c := range cases {
		actual := RectangularNumber(c.a, c.x0, c.c, c.m)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}
