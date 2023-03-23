// Copyright 2018 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache-2.0
// license that can be found in the LICENSE file.

package recordio

import (
	"fmt"
	"strings"
	"sync"

	"github.com/grailbio/base/recordio/recordioiov"
	"github.com/pkg/errors"
)

// TransformerFactory is a function that creates a new TransformerFunc given an
// optional config string.
type TransformerFactory func(config string) (TransformFunc, error)

type registerer struct {
	sync.RWMutex
	transformers   map[string]TransformerFactory
	untransformers map[string]TransformerFactory
}

var registry = &registerer{
	transformers:   make(map[string]TransformerFactory),
	untransformers: make(map[string]TransformerFactory),
}

func idTransform(scratch []byte, in [][]byte) ([]byte, error) {
	out := recordioiov.Slice(scratch, recordioiov.TotalBytes(in))
	n := 0
	for _, b := range in {
		copy(out[n:], b)
		n += len(b)
	}
	return out, nil
}

func (r *registerer) registerTransformer(name string, t TransformerFactory, u TransformerFactory) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.transformers[name]; ok {
		panic(fmt.Sprintf("Transformer %s already registered", name))
	}
	r.transformers[name] = t
	r.untransformers[name] = u
}

func getTransformers(h []string, factory map[string]TransformerFactory) ([]TransformFunc, error) {
	var transformers []TransformFunc
	for _, str := range h {
		toks := strings.SplitN(str, " ", 2)
		name := toks[0]
		f, ok := factory[name]
		if !ok {
			return nil, errors.Errorf("Transformer %s not found", str)
		}

		var config string
		if len(toks) > 1 {
			config = toks[1]
		}
		var tr TransformFunc
		var err error
		if tr, err = f(config); err != nil {
			return nil, err
		}
		transformers = append(transformers, tr)
	}
	return transformers, nil
}

func (r *registerer) getTransformer(h []string) (TransformFunc, error) {
	r.RLock()
	transformers, err := getTransformers(h, r.transformers)
	r.RUnlock()
	if err != nil {
		return nil, err
	}

	switch {
	case len(transformers) == 0:
		return idTransform, nil
	case len(transformers) == 1:
		return transformers[0], nil
	default:
		combined := func(scratch []byte, data [][]byte) ([]byte, error) {
			for _, tr := range transformers {
				out, err := tr(scratch, data)
				if err != nil {
					return nil, err
				}
				if len(data) == 0 {
					data = [][]byte{out}
				} else {
					data = data[:1]
					data[0] = out
				}
				scratch = nil
				// TODO(saito) Maybe reuse one of data[] as scratch?
			}
			if len(data) != 1 { // At least one transformer should have run.
				panic(data)
			}
			return data[0], nil
		}
		return combined, nil
	}
}

func (r *registerer) GetUntransformer(h []string) (TransformFunc, error) {
	r.RLock()
	transformers, err := getTransformers(h, r.untransformers)
	r.RUnlock()
	if err != nil {
		return nil, err
	}

	switch {
	case len(transformers) == 0:
		return idTransform, nil
	case len(transformers) == 1:
		return transformers[0], nil
	default:
		combined := func(scratch []byte, data [][]byte) ([]byte, error) {
			for i := len(transformers) - 1; i >= 0; i-- {
				tr := transformers[i]
				out, err := tr(scratch, data)
				if err != nil {
					return nil, err
				}
				if len(data) == 0 {
					data = [][]byte{out}
				} else {
					data = data[:1]
					data[0] = out
				}
			}
			if len(data) != 1 { // At least one transformer should have run.
				panic(data)
			}
			return data[0], nil
		}
		return combined, nil
	}
}

// RegisterTransformer registers a block transformer. Factory transformer should
// produce a transformer function. The factory is run by NewWriterV2.  The
// transformer function is called by the writer to transform a block just before
// storing it in storage.
//
// The untransformer factory is the reverse of the transformer factory. It is
// run by NewScannerV2. The untransformer function is called by the scanner to
// transform data read from storage into a block.
//
// This function is usually called when the process starts.
//
// The transformer and untransformer factories, as well as the functions
// generated by these factories must be all thread safe.
//
// REQUIRES: A (un)transformer with the same "name" has not been registered
// already.
func RegisterTransformer(name string, transformer TransformerFactory, untransformer TransformerFactory) {
	registry.registerTransformer(name, transformer, untransformer)
}
