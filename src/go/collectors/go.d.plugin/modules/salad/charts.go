package salad

import "github.com/netdata/netdata/go/go.d.plugin/agent/module"

var nodesChart = module.Chart{
	ID:    "nodes",
	Title: "Nodes count",
	Units: "nodes",
	Fam:   "nodes",
	Ctx:   "salad.nodes",
}

func initCharts() *module.Charts {
	charts := module.Charts{}
	_ = nodesChart.AddDim(&module.Dim{
		ID:   "all",
		Name: "all",
	})
	_ = charts.Add(&nodesChart)
	return &charts
}
