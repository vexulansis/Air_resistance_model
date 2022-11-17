package main

type Data struct { 
	mass                       float64
	initial_speed              float64
	air_resistance_coefficient float64
	time_resistance_asc        float64
	time_resistance_desc       float64
	height_resistance          float64
	time_no_resistance         float64
	height_no_resistance       float64
	accuracy                   int
}
type StartData struct {
	weight         float64
	initial_speed  float64
	air_resistance float64
	accuracy       int
}
type CurvesData struct {
	time                 []float64
	height_resistance    []float64
	height_no_resistance []float64
}
