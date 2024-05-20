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
		ID:   "active",
		Name: "active",
	})
	_ = nodesChart.AddDim(&module.Dim{
		ID:   "quarantined",
		Name: "quarantined",
	})
	_ = nodesChart.AddDim(&module.Dim{
		ID:   "zombied",
		Name: "zombied",
	})

	_ = charts.Add(&nodesChart)
	return &charts
}
