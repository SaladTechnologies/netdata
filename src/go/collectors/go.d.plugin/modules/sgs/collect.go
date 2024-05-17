// SPDX-License-Identifier: GPL-3.0-or-later

package sgs

import (
	"fmt"

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
	num := e.Config.Charts.Dims

	for i := 0; i < num; i++ {
		name := fmt.Sprintf("random%d", i)
		id := fmt.Sprintf("%s_%s", chart.ID, name)

		if !e.collectedDims[id] {
			e.collectedDims[id] = true

			dim := &module.Dim{ID: id, Name: name}
			if err := chart.AddDim(dim); err != nil {
				e.Warning(err)
			}
			chart.MarkNotCreated()
		}
		if i%2 == 0 {
			collected[id] = e.randInt()
		} else {
			collected[id] = -e.randInt()
		}
	}
}
