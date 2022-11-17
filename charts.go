package main

import (
	"fmt"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func (data *Data) CreateResistanceCurve() {
	curvesData := data.GenerateCurvesData()
	line := charts.NewLine()
	XLabel := new(opts.AxisLabel)
	XLabel.Interval = fmt.Sprintf("%d", len(curvesData.time)/100)
	XLabel.Rotate = 90
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeRoma}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Зависимость высоты тела от времени",
			Subtitle: "Шаг сетки соответствует выбранной точности",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:      "Время [c]",
			Show:      false,
			Scale:     false,
			Min:       0.0,
			AxisLabel: XLabel,
		}))

	res, no_res := curvesData.GenerateLinesData()
	line.SetXAxis(TimeConv(data, curvesData.time)).
		AddSeries("With Air Resistance", res).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	line.SetXAxis(TimeConv(data, curvesData.time)).
		AddSeries("With No Air Resistance", no_res).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	f, _ := os.Create("curves.html")
	line.Render(f)
}
func (data *Data) GenerateCurvesData() CurvesData {
	res := new(CurvesData)
	time := []float64{}
	height_resistance := []float64{}
	height_no_resistance := []float64{}
	limit := data.time_resistance_asc + data.time_resistance_desc
	step := math.Pow(10, -float64(data.accuracy))
	for t := 0.0; t <= data.time_no_resistance*2; t += step {
		time = append(time, t)
		if t <= data.time_resistance_asc {
			height_resistance = append(height_resistance, HeightResistanceAsc(data, t))
		} else if t > data.time_resistance_asc && t < limit {
			height_resistance = append(height_resistance, HeightResistanceDesc(data, t))
		} else {
			height_resistance = append(height_resistance, 0.0)
		}
		height_no_resistance = append(height_no_resistance, HeightNoResistance(data, t))
	}
	res.time = time
	res.height_resistance = height_resistance
	res.height_no_resistance = height_no_resistance
	return *res
}
func (curvesData *CurvesData) GenerateLinesData() ([]opts.LineData, []opts.LineData) {
	res := []opts.LineData{}
	no_res := []opts.LineData{}
	for i := range curvesData.time {
		res = append(res, opts.LineData{Value: curvesData.height_resistance[i], Symbol: "none", SymbolSize: 10})
		no_res = append(no_res, opts.LineData{Value: curvesData.height_no_resistance[i], Symbol: "none", SymbolSize: 10})
	}
	return res, no_res
}
func TimeConv(data *Data, time []float64) []string {
	res := []string{}
	for _, t := range time {
		res = append(res, Format(t, data.accuracy))
	}
	return res
}
