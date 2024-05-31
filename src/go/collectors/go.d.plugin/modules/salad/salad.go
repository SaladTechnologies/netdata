package salad

import (
	_ "embed"

	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

var (
	knownDestinations = []string{
		"nflx",
		"dsnp",
		"bbc",
		"iitv",
	}
	knownStatuses = []string{
		"active",
		"quarantined",
		"zombied",
	}
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
	client *Client
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
	client, err := NewClient()
	if err != nil {
		return err
	}
	s.client = client
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
	mx := map[string]int64{}
	err := s.client.CollectHealth(mx)
	if err != nil {
		s.Error(err)
	}
	err = s.client.CollectCounters(mx)
	if err != nil {
		s.Error(err)
	}

	return mx
}

func (s *Salad) Configuration() any {
	return s.Config
}
