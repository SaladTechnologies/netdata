package salad

import "github.com/netdata/netdata/go/go.d.plugin/agent/module"

var nodesChart = module.Chart{
	ID: "nodes",
}

func initCharts() *module.Charts {
	charts := module.Charts{}
	charts.Add(&nodesChart)
	return &charts
}
