package salad

import "github.com/netdata/netdata/go/go.d.plugin/agent/module"

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
	return map[string]int64{
		"foo": 42,
	}
}

func (s *Salad) Configuration() any {
	return s.Config
}
