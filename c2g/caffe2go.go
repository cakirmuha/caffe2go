package c2g

import (
	"errors"
	"image"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/gkostadinov/caffe2go/caffe"
	"github.com/gkostadinov/caffe2go/layers"
	"github.com/gkostadinov/caffe2go/network"
	"github.com/golang/protobuf/proto"
	"github.com/nfnt/resize"
)

// Caffe2Go is interface of caffe2go.
type Caffe2Go struct {
	Network *network.Network
}

// NewCaffe2Go is constructor.
func NewCaffe2Go(modelPath string) (*Caffe2Go, error) {
	data, err := ioutil.ReadFile(modelPath)
	if err != nil {
		log.Fatalln(err)
	}
	var netParameter caffe.NetParameter
	if err = proto.Unmarshal(data, &netParameter); err != nil {
		log.Fatalln(err)
	}
	var net network.Network
	if len(netParameter.GetLayer()) != 0 {
		//showLayers(netParameter.GetLayer())
		for i := range netParameter.GetLayer() {
			switch netParameter.GetLayer()[i].GetType() {
			case layers.InnerProduct:
				fcLayer, err := SetupFullconnect(netParameter.GetLayer()[i])
				if err != nil {
					return nil, err
				}
				net.Add(fcLayer)
			case layers.Convolution:
				convLayer := SetupConvolution(netParameter.GetLayer()[i])
				net.Add(convLayer)
			case layers.Pooling:
				poolLayer := SetupPooling(netParameter.GetLayer()[i])
				net.Add(poolLayer)
			case layers.Dropout:
				dropoutLayer := SetupDropout(netParameter.GetLayer()[i])
				net.Add(dropoutLayer)
			case layers.Softmax:
				softMaxLayer := SetupSoftmaxLoss(netParameter.GetLayer()[i])
				net.Add(softMaxLayer)
			case layers.ReLU:
				reluLayer := SetupReLU(netParameter.GetLayer()[i])
				net.Add(reluLayer)
			case layers.SoftmaxWithLoss:
				softmaxLossLayer := SetupSoftmaxLoss(netParameter.GetLayer()[i])
				net.Add(softmaxLossLayer)
			case layers.Lrn:
				lrnLayer := SetupLRN(netParameter.GetLayer()[i])
				net.Add(lrnLayer)
			}
		}
	} else {
		//showV1Layers(netParameter.GetLayers())
		for i := range netParameter.GetLayers() {
			switch netParameter.GetLayers()[i].GetType() {
			case caffe.V1LayerParameter_INNER_PRODUCT:
				fcLayer, err := SetupFullconnect(netParameter.GetLayers()[i])
				if err != nil {
					return nil, err
				}
				net.Add(fcLayer)
			case caffe.V1LayerParameter_CONVOLUTION:
				convLayer := SetupConvolution(netParameter.GetLayers()[i])
				net.Add(convLayer)
			case caffe.V1LayerParameter_POOLING:
				poolLayer := SetupPooling(netParameter.GetLayers()[i])
				net.Add(poolLayer)
			case caffe.V1LayerParameter_DROPOUT:
				dropoutLayer := SetupDropout(netParameter.GetLayers()[i])
				net.Add(dropoutLayer)
			case caffe.V1LayerParameter_SOFTMAX:
				softMaxLayer := SetupSoftmaxLoss(netParameter.GetLayers()[i])
				net.Add(softMaxLayer)
			case caffe.V1LayerParameter_RELU:
				reluLayer := SetupReLU(netParameter.GetLayers()[i])
				net.Add(reluLayer)
			case caffe.V1LayerParameter_SOFTMAX_LOSS:
				softmaxLossLayer := SetupSoftmaxLoss(netParameter.GetLayers()[i])
				net.Add(softmaxLossLayer)
			case caffe.V1LayerParameter_LRN:
				lrnLayer := SetupLRN(netParameter.GetLayers()[i])
				net.Add(lrnLayer)
			}
		}
	}
	return &Caffe2Go{
		Network: &net,
	}, nil
}

// Predict start network.
func (c2g *Caffe2Go) Predict(img image.Image, size uint, means []float32, endLayer string) ([][][]float32, error) {
	img = resize.Resize(size, size, img, resize.Lanczos3)
	input := im2vec(img, means)
	return c2g.Network.Predict(input, endLayer)
}

func im2vec(img image.Image, means []float32) [][][]float32 {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	var res [][][]float32
	if img.ColorModel() == color.GrayModel {
		res = make([][][]float32, 1)
		res[0] = make([][]float32, height)
	} else {
		res = make([][][]float32, 3)
		for i := 0; i < 3; i++ {
			res[i] = make([][]float32, height)
		}
	}
	for y := 0; y < height; y++ {
		for i := 0; i < len(res); i++ {
			res[i][y] = make([]float32, width)
		}
		for x := 0; x < width; x++ {
			c := img.At(x, y)
			if img.ColorModel() == color.GrayModel {
				grayColor := img.ColorModel().Convert(c)
				if means != nil {
					res[0][y][x] = float32(grayColor.(color.Gray).Y) - means[0]
				} else {
					res[0][y][x] = float32(grayColor.(color.Gray).Y)
				}
			} else {
				r, g, b, _ := c.RGBA()
				if means != nil {
					res[0][y][x] = (float32(r)/255 - means[0])
					res[1][y][x] = (float32(g)/255 - means[1])
					res[2][y][x] = (float32(b)/255 - means[2])
				} else {
					res[0][y][x] = (float32(r) / 255)
					res[1][y][x] = (float32(g) / 255)
					res[2][y][x] = (float32(b) / 255)
				}
			}
		}
	}
	return res
}

func crop(tensor [][][]float32, l int) ([][][]float32, error) {
	w := len(tensor[0][0])
	h := len(tensor[0])
	if h < l || w < l {
		return nil, errors.New("Length is mismatched")
	}
	var w1, h1 int
	if w > h {
		w1 = l * w / h
		h1 = l
	} else {
		w1 = l
		h1 = l * h / w
	}
	sx := (w1 - l) / 2
	sy := (h1 - l) / 2
	res := make([][][]float32, len(tensor))
	for i := range tensor {
		res[i] = make([][]float32, l)
		for j, s1 := range tensor[i][sy : sy+l] {
			res[i][j] = s1[sx : sx+l]
		}
	}
	return res, nil
}
