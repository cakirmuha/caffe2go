package layers

// Im2Col converts image 3D tensor to matrix.
func Im2Col(img [][][]float32, kernelSize, stride int) [][]float32 {
	colSize := kernelSize * kernelSize * len(img)
	rows := (len(img[0])-kernelSize)/stride + 1
	cols := (len(img[0][0])-kernelSize)/stride + 1
	res := make([][]float32, rows*cols)
	idx1 := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			col := make([]float32, colSize)
			idx2 := 0
			sy := y * stride
			sx := x * stride
			for c := range img {
				for i := sy; i < sy+kernelSize; i++ {
					for j := sx; j < sx+kernelSize; j++ {
						col[idx2] = img[c][i][j]
						idx2++
					}
				}
			}
			res[idx1] = col
			idx1++
		}
	}
	return res
}