package sgs

import "github.com/netdata/netdata/go/go.d.plugin/agent/module"

var nodesChart = module.Chart{
	ID:    "nodes",
	Title: "Nodes count",
	Units: "nodes",
	Fam:   "nodes",
	Ctx:   "sgs.nodes",
}
