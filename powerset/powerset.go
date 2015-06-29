package powerset

import "math/big"

func addToAll(elem interface{}, arr [][]interface{}) [][]interface{} {

	newarr := make([][]interface{}, len(arr))

	for i := range newarr {
		newelem := make([]interface{}, 0, 1+len(arr[i]))
		newelem = append(newelem, elem)
		newelem = append(newelem, arr[i]...)
		newarr[i] = newelem
	}
	return newarr
}

func CreatePowerSet(set []interface{}) [][]interface{} {

	if len(set) == 0 {
		// return set containing an empty set
		result := make([][]interface{}, 1)
		result[0] = make([]interface{}, 0)
		return result
	}

	// Take the first element, and then the set without it:
	head := set[0]
	tail := set[1:]

	// get sub sets for the tails
	tailSubSets := CreatePowerSet(tail)

	// creates the sub sets that include the head
	subSetsWithHead := addToAll(head, tailSubSets)

	// return all sub sets
	return append(subSetsWithHead, tailSubSets...)

}


func StreamPowerSet(set []interface{}) <- chan []interface{} {
	state := big.NewInt(1)
	one := big.NewInt(1)
	two := big.NewInt(2)
	for i := 0 ; i < len(set); i++ {
		state.Mul(state, two)
	}
	
	c := make(chan []interface{})
	go func() {	
		for state.Sign() > 0 {
			state.Sub(state, one)
			currentSubSet := make([]interface{}, 0)
			bitLen := state.BitLen()
			for i := 0; i < bitLen; i++ {
				if state.Bit(i) != 0 {
					currentSubSet = append(currentSubSet, set[i])
				}
			}
			c <- currentSubSet
		}
		close(c)
	}()
	
	return c
}
