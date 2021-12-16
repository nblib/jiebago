package finalseg

const (
	State_B = 0
	State_E = 1
	State_M = 2
	State_S = 3
)
const minFloat = -3.14e100

var (
	prevStatusNew = make([][]int, 4, 4)
	probStartNew  = make([]float64, 4, 4)
)

func init() {
	prevStatusNew[State_B] = []int{State_E, State_S}
	prevStatusNew[State_M] = []int{State_M, State_B}
	prevStatusNew[State_S] = []int{State_S, State_E}
	prevStatusNew[State_E] = []int{State_B, State_M}

	probStartNew[State_B] = -0.26268660809250016
	probStartNew[State_E] = -3.14e+100
	probStartNew[State_M] = -3.14e+100
	probStartNew[State_S] = -1.4652633398537678
}

type VbPropState struct {
	prop  float64
	state int
}

func vbGetStartprob(state int, c rune) float64 {
	if val, ok := probEmitNew[state][c]; ok {
		return val + probStartNew[state]
	} else {
		return minFloat + probStartNew[state]
	}
}
func vbGetEmitProb(state int, c rune) float64 {
	if val, ok := probEmitNew[state][c]; ok {
		return val
	} else {
		return minFloat
	}
}
func vbGetTransProb(prev, cur int) float64 {
	if tp, ok := probTransNew[prev][cur]; ok {
		return tp
	} else {
		return minFloat
	}
}
func vbInitPath(n int) [][]int {
	path := make([][]int, 4, 4)
	path[State_B] = make([]int, 0, n)
	path[State_M] = make([]int, 0, n)
	path[State_S] = make([]int, 0, n)
	path[State_E] = make([]int, 0, n)
	path[State_B] = append(path[State_B], State_B)
	path[State_M] = append(path[State_M], State_M)
	path[State_S] = append(path[State_S], State_S)
	path[State_E] = append(path[State_E], State_E)

	return path
}
func viterbiNew(obs []rune) (float64, []int) {
	lenObs := len(obs)

	path := vbInitPath(lenObs)
	V := make([][]float64, len(obs))
	V[0] = make([]float64, 4, 4)

	V[0][State_B] = vbGetStartprob(State_B, obs[0])
	V[0][State_M] = vbGetStartprob(State_M, obs[0])
	V[0][State_S] = vbGetStartprob(State_S, obs[0])
	V[0][State_E] = vbGetStartprob(State_E, obs[0])

	for t := 1; t < len(obs); t++ {
		V[t] = make([]float64, 4, 4)
		newPath := make([][]int, 4, 4)
		for y := 0; y < 4; y++ {
			var maxPropState *VbPropState
			emP := vbGetEmitProb(y, obs[t])
			for _, y0 := range prevStatusNew[y] {
				transP := vbGetTransProb(y0, y)
				prob0 := V[t-1][y0] + transP + emP

				if maxPropState == nil {
					maxPropState = &VbPropState{prop: prob0, state: y0}
					continue
				}
				if prob0 > maxPropState.prop {
					maxPropState.prop, maxPropState.state = prob0, y0
				} else if prob0 == maxPropState.prop && y0 > maxPropState.state {
					maxPropState.prop, maxPropState.state = prob0, y0
				}
			}
			V[t][y] = maxPropState.prop
			newPath[y] = append(path[maxPropState.state], y)
		}
		path = newPath
	}
	finalE := V[len(obs)-1][State_E]
	finalS := V[len(obs)-1][State_S]

	if finalE >= finalS {
		return finalE, path[State_E]
	} else {
		return finalS, path[State_S]
	}
}
