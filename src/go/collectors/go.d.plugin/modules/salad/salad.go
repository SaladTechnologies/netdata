package salad

import (
	_ "embed"

	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

//go:embed "config_schema.json"
var configSchema string

type Config struct {
	UpdateEvery int `yaml:"update_every,omitempty" json:"update_every"`
}

type Salad struct {
	module.Base
	Config `yaml:",inline"`
	charts *module.Charts
}

func init() {
	module.Register("salad", module.Creator{
		JobConfigSchema: configSchema,
		Defaults: module.Defaults{
			UpdateEvery: module.UpdateEvery,
			Priority:    module.Priority,
		},
		Create: func() module.Module { return &Salad{} },
	})
}

func (s *Salad) Init() error {
	s.charts = initCharts()
	return nil
}

func (s *Salad) Charts() *module.Charts {
	return s.charts
}

func (s *Salad) Check() error {
	return nil
}

func (s *Salad) Cleanup() {}

func (s *Salad) Collect() map[string]int64 {
	mx := map[string]int64{
		"foo": 42,
	}
	return mx
}

func (s *Salad) Configuration() any {
	return s.Config
}
