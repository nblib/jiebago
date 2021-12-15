package finalseg

var probTransNew = make([]map[int]float64, 4, 4)

func init() {
	probTransNew[State_B] = map[int]float64{State_E: -0.510825623765990,
		State_M: -0.916290731874155}
	probTransNew[State_E] = map[int]float64{State_B: -0.5897149736854513,
		State_S: -0.8085250474669937}
	probTransNew[State_M] = map[int]float64{State_E: -0.33344856811948514,
		State_M: -1.2603623820268226}
	probTransNew[State_S] = map[int]float64{State_B: -0.7211965654669841,
		State_S: -0.6658631448798212}
}
