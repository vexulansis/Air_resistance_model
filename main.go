package main

import (
	"fmt"
	"time"
)

const g = 9.80665

func main() {
	close := false
	for close != true {
		log := []string{}
		var data Data
		startData, err := InputStartData()
		if err == nil {
			data = CalculateData(startData)
			log = validationLog(startData, &data)
		}
		if len(log) == 0 && err == nil {
			data.PrintData(int(startData.accuracy))
			data.CreateResistanceCurve()
			close = true
		}
	}
	time.Sleep(time.Minute * 3)
}
func validationLog(startdata *StartData, data *Data) []string {
	log := []string{}
	if startdata.air_resistance <= 0 || startdata.initial_speed <= 0 || startdata.weight <= 0 {
		log = append(log, "Ошибка: Входные данные должны быть больше нуля\n")
	}
	if data.height_no_resistance >= 6400 || data.height_resistance >= 6400 {
		log = append(log, "Ошибка: Получившиеся значения выходят за границы применимости g\n")
	}
	if len(log) > 0 {
		for _, l := range log {
			fmt.Printf(l)
		}
	}
	return log
}
