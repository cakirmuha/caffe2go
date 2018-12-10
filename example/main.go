package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"

	"github.com/cakirmuha/caffe2go/c2g"
	"os"
	"image"
)

func main() {
	caffe2go, err := c2g.NewCaffe2Go("lenet.caffemodel")
	if err != nil {
		panic(err)
	}
	reader, err := os.Open("mnist_zero.png")
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	output, err := caffe2go.Predict(img, 28, nil, "")
	if err != nil {
		panic(err)
	}

	for i := range output {
		fmt.Printf("%d: %f\n", i, output[i][0][0])
	}

}
