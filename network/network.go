package network

import (
	"github.com/cakirmuha/caffe2go/layers"
)

// Network have netword definition.
type Network struct {
	layers []layers.Layer
}

// Add adds layer to network.
func (n *Network) Add(layer layers.Layer) {
	n.layers = append(n.layers, layer)
}

// Predict forwards network.
func (n *Network) Predict(input [][][]float32, endLayer string) (output [][][]float32, err error) {
	for i := range n.layers {
		input, err = n.layers[i].Forward(input)
		if err != nil {
			return
		}
		if n.layers[i].GetName() == endLayer {
			break
		}
	}
	output = input
	return
}
