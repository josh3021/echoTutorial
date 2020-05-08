package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	if v := Sum(1, 2, 3); v != 6 {
		t.Fatal("1 + 2 + 3 == 6")
	}

	assert.Equal(t, Sum(1, 2, 3), 6, "Equal")

	if v := Sum(6, 5); v != 11 {
		t.Fatal("6 + 5 == 11")
	}
}

func Test_Div(t *testing.T) {
	if v, err := Div(6.5, 5.7); v != 1.1403508771929824 || err != nil {
		t.Fatal("6.5 / 5.7 == 1.1403508771929824")
	}

	if v, err := Div(0, 2); v != 0.0 || err != nil {
		t.Fatal("0 / 2 == 0")
	}

	if v, err := Div(1, 0); v != 0.0 || err == nil {
		t.Fatal("it should return error")
	}
}

func Test_Fibo(t *testing.T) {
	if v := Fibo(6); v != 8 {
		t.Fatal("Fibo(6) must returns 8")
	}
}
