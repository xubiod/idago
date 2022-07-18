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

func (dst *Axon) Merge(src *Axon, preference float32) {
	(*dst).Multiplier = (((*dst).Multiplier * (1 - preference)) + ((*src).Multiplier * preference))
}

func (dst *Layer) Merge(src *Layer, preference float32) error {
	if len(*dst) != len(*src) {
		return errors.New("dumbbrain: axon count in layer mismatch, no changes done")
	}

	for intoIndex, intoAxon := range *dst {
		intoAxon.Merge(&(*src)[intoIndex], preference)
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

func (dst *DumbBrain) Merge(src *DumbBrain, preference float32) error {
	if len(*dst) != len(*src) {
		return errors.New("dumbbrain: layer count mismatch, no changes done")
	}

	for intoIndex, _ := range *dst {
		(*dst)[intoIndex].Merge(&(*src)[intoIndex], preference)
	}
	return nil
}
