package idago

import (
	"errors"
	"fmt"
	"math/rand"
)

// Note absolutely no experience with neural networks or machine learning has
// gone into this and don't feel like doing any as well :)

type Neuron float32

type Axon struct {
	Start      *Neuron
	End        *Neuron
	Multiplier float32
}

type Layer []*Axon

type DumbBrain []*Layer

func (dst *Axon) Merge(src *Axon, preference float32) {
	(*dst).Multiplier = (((*dst).Multiplier * (1 - preference)) + ((*src).Multiplier * preference))
}

func (dst *Layer) Merge(src *Layer, preference float32) error {
	if len(*dst) != len(*src) {
		return errors.New("dumbbrain/merge: axon count in layer mismatch, no changes done")
	}

	for dstIndex := range *dst {
		(*dst)[dstIndex].Merge((*src)[dstIndex], preference)
	}

	return nil
}

func (layer *Layer) Passdown() {
	var fromVal float32
	for _, axon := range []*Axon(*layer) {
		*axon.End = 1.0
	}

	for _, axon := range []*Axon(*layer) {
		fromVal = float32(*axon.Start)
		*axon.End *= Neuron(fromVal * axon.Multiplier)
	}
}

func (network *DumbBrain) Runthrough() {
	for _, layer := range []*Layer(*network) {
		layer.Passdown()
	}
}

func (dst *DumbBrain) Merge(src *DumbBrain, preference float32) error {
	if len(*dst) != len(*src) {
		return errors.New("dumbbrain/merge: layer count mismatch, no changes done")
	}

	for dstIndex := range *dst {
		(*dst)[dstIndex].Merge((*src)[dstIndex], preference)
	}
	return nil
}

func Stork(templateStructure []int, axonMin float32, axonMax float32, into *DumbBrain) error {
	if *into != nil {
		return errors.New("dumbbrain/stork: refusal to change non-nil network")
	}

	*into = DumbBrain(make([]*Layer, len(templateStructure)))

	for layerIndex := range *into {
		*(*into)[layerIndex] = Layer(make([]*Axon, templateStructure[layerIndex]))

		for axonIndex := range *(*into)[layerIndex] {
			(*(*into)[layerIndex])[axonIndex].Multiplier = (rand.Float32() * (axonMax - axonMin)) + axonMin
		}
	}

	return nil
}

func StorkMany(templateStructure []int, axonMin float32, axonMax float32, into *[]*DumbBrain) error {
	if len(*into) <= 0 {
		return errors.New("dumbbrain/storks: into is empty")
	}

	var result error
	for networkIndex := range *into {
		result = Stork(templateStructure, axonMin, axonMax, (*into)[networkIndex])
		if result != nil {
			return fmt.Errorf("dumbbrain/storks: issue at index %d:\n\t%s", networkIndex, result.Error())
		}
	}

	return nil
}
