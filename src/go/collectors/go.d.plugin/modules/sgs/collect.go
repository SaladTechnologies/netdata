// SPDX-License-Identifier: GPL-3.0-or-later

package sgs

import (
	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

func (e *SaladGateway) collect() (map[string]int64, error) {
	collected := make(map[string]int64)

	for _, chart := range *e.Charts() {
		e.collectChart(collected, chart)
	}
	return collected, nil
}

func (e *SaladGateway) collectChart(collected map[string]int64, chart *module.Chart) {

	nodes, err := e.client.GetNodeCount()
	if err != nil {
		e.Error(err)
	}

	id := "nodes"
	name := "nodes"
	e.collectedDims[id] = err != nil
	collected[id] = int64(nodes)
	dim := &module.Dim{
		ID:   id,
		Name: name,
	}
	if err := chart.AddDim(dim); err != nil {
		e.Warning(err)
	}
	chart.MarkNotCreated()
	collected[id] = int64(nodes)

}
