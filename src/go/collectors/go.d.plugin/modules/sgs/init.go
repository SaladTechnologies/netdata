package sgs

import (
	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

func (e *SaladGateway) initCharts() (*module.Charts, error) {
	charts := &module.Charts{}
	charts.Add(&nodesChart)
	return charts, nil
}
