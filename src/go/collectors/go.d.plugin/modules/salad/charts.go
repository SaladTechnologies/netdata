package salad

import (
	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

var nodesChart = module.Chart{
	ID:    "nodes_statuses",
	Title: "Nodes Per Status",
	Units: "nodes",
	Fam:   "nodes",
	Ctx:   "salad.nodes_status",
}

var destinationsChart = module.Chart{
	ID:    "nodes_destinations",
	Title: "Nodes Per Destination",
	Units: "nodes",
	Ctx:   "salad.nodes_destinations",
}

var streamsChart = module.Chart{
	ID:    "streams",
	Title: "Active Streams",
	Units: "streams",
	Ctx:   "salad.streams",
	Dims: []*module.Dim{
		{
			ID:   "streams.active",
			Name: "active",
		},
	},
}

func initCharts() *module.Charts {
	charts := module.Charts{}

	for _, status := range knownStatuses {
		nodesChart.AddDim(&module.Dim{
			ID: status,
		})
	}

	for _, dest := range knownDestinations {
		_ = destinationsChart.AddDim((&module.Dim{
			ID: dest,
		}))
	}

	_ = charts.Add(&nodesChart)
	_ = charts.Add(&destinationsChart)
	_ = charts.Add(&streamsChart)

	return &charts
}
