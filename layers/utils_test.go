package layers

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestIm2Col(t *testing.T) {
	m := [][][]float32{
		{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 1, 0},
			{0, 1, 2, 0, 0, 1, 0},
			{0, 2, 2, 1, 2, 2, 0},
			{0, 0, 0, 1, 2, 1, 0},
			{0, 2, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 1, 0},
			{0, 1, 2, 0, 0, 1, 0},
			{0, 2, 2, 1, 2, 2, 0},
			{0, 0, 0, 1, 2, 1, 0},
			{0, 2, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 2, 0, 0, 1, 0},
			{0, 1, 2, 0, 0, 1, 0},
			{0, 2, 2, 1, 2, 2, 0},
			{0, 0, 0, 1, 2, 1, 0},
			{0, 2, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	}

	res := mat.NewMatrix(Im2Col(m, 3, 2))

	a := mat.NewMatrix([][]float32{
		{0, 0, 0, 0, 0, 2, 0, 1, 2, 0, 0, 0, 0, 0, 2, 0, 1, 2, 0, 0, 0, 0, 0, 2, 0, 1, 2},
		{0, 0, 0, 2, 0, 0, 2, 0, 0, 0, 0, 0, 2, 0, 0, 2, 0, 0, 0, 0, 0, 2, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0},
		{0, 1, 2, 0, 2, 2, 0, 0, 0, 0, 1, 2, 0, 2, 2, 0, 0, 0, 0, 1, 2, 0, 2, 2, 0, 0, 0},
		{2, 0, 0, 2, 1, 2, 0, 1, 2, 2, 0, 0, 2, 1, 2, 0, 1, 2, 2, 0, 0, 2, 1, 2, 0, 1, 2},
		{0, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 0, 2, 2, 0, 2, 1, 0},
		{0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0},
		{0, 1, 2, 1, 1, 1, 0, 0, 0, 0, 1, 2, 1, 1, 1, 0, 0, 0, 0, 1, 2, 1, 1, 1, 0, 0, 0},
		{2, 1, 0, 1, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 0, 0},
	})

	if !res.Equals(a) {
		t.Error("not same")
		a.Show()
		res.Show()
	}

}

func TestConvertMatrix(t *testing.T) {
	src := [][]float32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	dst := ConvertMatrix(src)
	r, c := dst.Dims()
	if r != len(src) || c != len(src[0]) {
		t.Error("not same")
	}

	for i := range src {
		for j := range src[i] {
			if src[i][j] != float32(dst.At(i, j)) {
				t.Error("not same")
			}
		}
	}
}

func TestConvertMat64(t *testing.T) {
	src := mat.NewDense(3, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
	})
	ans := [][]float32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	res := ConvertMat64(src)
	for i := range ans {
		for j := range ans[i] {
			if res[i][j] != ans[i][j] {
				t.Error("not same")
			}
		}
	}
}
