package list

import "testing"

func TestList(t *testing.T) {
	l := &List{}
	values := []string{"null", "eins", "zwei", "drei"}
	for i, v := range values {
		err := l.AddAt(i, v)
		if err != nil {
			t.Error("no error expected", err)
		}
	}
	for i, v := range values {
		got, err := l.Get(i)
		if err != nil {
			t.Error("no error expected", err)
		}
		if got.Value != v {
			t.Errorf("at index %d exp: %s got %s", i, v, got.Value)
		}
	}
}
