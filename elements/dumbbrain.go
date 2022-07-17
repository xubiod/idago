package idago

// Note absolutely no experience with neural networks or machine learning has
// gone into this and don't feel like doing any as well :)

type Neuron float32

type Axon struct {
	Start      *Neuron
	End        *Neuron
	Multiplier float32
}

type Layer []Axon

type DumbBrain []Layer

func (layer *Layer) Passdown() {
	var fromVal float32
	for _, axon := range []Axon(*layer) {
		*axon.End = 1.0
	}

	for _, axon := range []Axon(*layer) {
		fromVal = float32(*axon.Start)
		*axon.End *= Neuron(fromVal * axon.Multiplier)
	}
}

// func (into *DumbBrain) Merge(with *DumbBrain, preference float32) {
// 	for layerIndex, layer := range []Layer(*with) {
// 		for axonIndex, axon := range []Axon(layer) {

// 		}
// 	}
// }
