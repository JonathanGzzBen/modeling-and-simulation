package pkg

import "testing"

func TestZ0(t *testing.T) {
	cases := []struct {
		average  float64
		expected float64
	}{{0.330173, 1.86036}}
	for _, c := range cases {
		actual := z0(c.average)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}

func TestZAlphaHalf(t *testing.T) {
	cases := []struct {
		alpha    float64
		expected float64
	}{{0.05, 0.475}}
	for _, c := range cases {
		actual := zAlphaHalf(c.alpha)
		if actual != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, actual)
		}
	}
}
