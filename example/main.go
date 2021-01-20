package main

import (
	"github.com/fale/sit"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	p, err := plot.New()
	if err != nil {
		return
	}

	line := plotter.XYs{
		plotter.XY{X: float64(1), Y: float64(1)},
		plotter.XY{X: float64(2), Y: float64(11)},
		plotter.XY{X: float64(3), Y: float64(5)},
		plotter.XY{X: float64(4), Y: float64(2)},
		plotter.XY{X: float64(5), Y: float64(7)},
	}

	p.Add(plotter.NewGrid())
	err = plotutil.AddLinePoints(p,
		"First", line,
	)
	if err != nil {
		return
	}

	p.Y.Tick.Marker = sit.Ticker{}
	p.Y.Min = sit.Min(p.Y.Min, p.Y.Max)
	p.Y.Max = sit.Max(p.Y.Min, p.Y.Max)

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "points.png"); err != nil {
		return
	}
	return
}
