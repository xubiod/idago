package idago

import "errors"

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

func (into *Axon) Merge(with *Axon, preference float32) {
	(*into).Multiplier = (((*into).Multiplier * (1 - preference)) + ((*with).Multiplier * preference))
}

func (into *Layer) Merge(with *Layer, preference float32) error {
	if len(*into) != len(*with) {
		return errors.New("dumbbrain: axon count in layer mismatch, no changes done")
	}

	for intoIndex, intoAxon := range *into {
		intoAxon.Merge(&(*with)[intoIndex], preference)
	}

	return nil
}

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

func (network *DumbBrain) Runthrough() {
	for _, layer := range []Layer(*network) {
		layer.Passdown()
	}
}

func (into *DumbBrain) Merge(with *DumbBrain, preference float32) error {
	if len(*into) != len(*with) {
		return errors.New("dumbbrain: layer count mismatch, no changes done")
	}

	for intoIndex, _ := range *into {
		(*into)[intoIndex].Merge(&(*with)[intoIndex], preference)
	}
	return nil
}
