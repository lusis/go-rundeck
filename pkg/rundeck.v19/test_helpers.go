package rundeck

import (
	"testing"
)

func strexpects(a string, e string, t *testing.T) {
	if a != e {
		t.Errorf("Expected: %s != actual: %s", e, a)
	}
}

func intexpects(a int64, e int64, t *testing.T) {
	if a != e {
		t.Errorf("Expected: %d != actual: %d", e, a)
	}
}

func f64expects(a float64, e float64, t *testing.T) {
	if a != e {
		t.Errorf("Expected: %v != actual: %v", e, a)
	}
}
