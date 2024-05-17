package sgs

import (
	_ "embed"
	"math/rand"

	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

//go:embed "config_schema.json"
var configSchema string

func init() {
	module.Register("example", module.Creator{
		JobConfigSchema: configSchema,
		Defaults: module.Defaults{
			UpdateEvery: module.UpdateEvery,
			Priority:    module.Priority,
			Disabled:    true,
		},
		Create: func() module.Module { return New() },
	})
}

func New() *SaladGateway {
	return &SaladGateway{
		Config: Config{
			Charts: ConfigCharts{
				Num:  1,
				Dims: 1,
			},
		},

		randInt:       func() int64 { return rand.Int63n(100) },
		collectedDims: make(map[string]bool),
	}
}

type (
	Config struct {
		UpdateEvery int          `yaml:"update_every" json:"update_every"`
		Charts      ConfigCharts `yaml:"charts" json:"charts"`
	}
	ConfigCharts struct {
		Type     string `yaml:"type" json:"type"`
		Num      int    `yaml:"num" json:"num"`
		Contexts int    `yaml:"contexts" json:"contexts"`
		Dims     int    `yaml:"dimensions" json:"dimensions"`
		Labels   int    `yaml:"labels" json:"labels"`
	}
)

type SaladGateway struct {
	module.Base // should be embedded by every module
	Config      `yaml:",inline"`

	randInt       func() int64
	charts        *module.Charts
	collectedDims map[string]bool
}

func (e *SaladGateway) Configuration() any {
	return e.Config
}

func (e *SaladGateway) Init() error {
	err := e.validateConfig()
	if err != nil {
		e.Errorf("config validation: %v", err)
		return err
	}

	charts, err := e.initCharts()
	if err != nil {
		e.Errorf("charts init: %v", err)
		return err
	}
	e.charts = charts
	return nil
}

func (e *SaladGateway) Check() error {
	return nil
}

func (e *SaladGateway) Charts() *module.Charts {
	return e.charts
}

func (e *SaladGateway) Collect() map[string]int64 {
	mx, err := e.collect()
	if err != nil {
		e.Error(err)
	}

	if len(mx) == 0 {
		return nil
	}
	return mx
}

func (e *SaladGateway) Cleanup() {}
