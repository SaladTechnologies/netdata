package salad

import (
	"fmt"

	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

var destinations = []string{
	"nflx",
	"dsnp",
	"bbc",
	"iitv",
}

var nodesChart = module.Chart{
	ID:    "nodes",
	Title: "Nodes",
	Units: "nodes",
	Fam:   "nodes",
	Ctx:   "salad.nodes",
}

var destinationsChart = module.Chart{
	ID:    "destinations",
	Title: "Destinations",
	Units: "destinations",
	Ctx:   "salad.destinations",
}

func initCharts() *module.Charts {
	charts := module.Charts{}

	_ = nodesChart.AddDim(&module.Dim{
		ID:   "nodes.active",
		Name: "active",
	})
	_ = nodesChart.AddDim(&module.Dim{
		ID:   "nodes.quarantined",
		Name: "quarantined",
	})
	_ = nodesChart.AddDim(&module.Dim{
		ID:   "nodes.zombied",
		Name: "zombied",
	})

	for _, dest := range destinations {
		_ = destinationsChart.AddDim((&module.Dim{
			ID:   fmt.Sprintf("destinations.%s", dest),
			Name: dest,
		}))
	}

	_ = charts.Add(&nodesChart)
	_ = charts.Add(&destinationsChart)

	return &charts
}
