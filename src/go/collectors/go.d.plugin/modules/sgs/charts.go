package sgs

import "github.com/netdata/netdata/go/go.d.plugin/agent/module"

var nodesChart = module.Chart{
	ID:    "random_%d",
	Title: "A Random Number",
	Units: "random",
	Fam:   "random",
	Ctx:   "example.random",
}
