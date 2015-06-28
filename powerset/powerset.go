package powerset

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
