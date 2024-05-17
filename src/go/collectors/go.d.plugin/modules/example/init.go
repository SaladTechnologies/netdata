package example

import (
	"errors"

	"github.com/netdata/netdata/go/go.d.plugin/agent/module"
)

func (e *SaladGateway) validateConfig() error {
	if e.Config.Charts.Num > 0 && e.Config.Charts.Dims <= 0 {
		return errors.New("'charts->dimensions' must be > 0")
	}
	return nil
}

func (e *SaladGateway) initCharts() (*module.Charts, error) {
	charts := &module.Charts{}

	var ctx int
	v := calcContextEvery(e.Config.Charts.Num, e.Config.Charts.Contexts)
	for i := 0; i < e.Config.Charts.Num; i++ {
		if i != 0 && v != 0 && ctx < (e.Config.Charts.Contexts-1) && i%v == 0 {
			ctx++
		}
		chart := newChart(i, ctx, e.Config.Charts.Labels, module.ChartType(e.Config.Charts.Type))

		if err := charts.Add(chart); err != nil {
			return nil, err
		}
	}
	return charts, nil
}

func calcContextEvery(charts, contexts int) int {
	if contexts <= 1 {
		return 0
	}
	if contexts > charts {
		return 1
	}
	return charts / contexts
}
