package layers

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestLRN(t *testing.T) {
	lrn := LRN{
		N:     5,
		K:     2,
		Alpha: 0.0005,
		Beta:  0.75,
	}

	input := make([][][]float32, 96)
	for i := range input {
		m := mat.Random(64, 64)
		input[i] = m.M
	}

	res, err := lrn.Forward(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
