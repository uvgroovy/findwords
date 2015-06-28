package powerset

import (
	"testing"
)

func pow2(exp int) int {
	result := 1
	for i :=0; i < exp; i++ {
		result *= 2
	}
	return result
} 

func TestPowerSetSize(t *testing.T) {
	
	in := make([]interface{}, 0)
	for i :=0; i < 10; i++ {
		want := pow2(i)
		got := len(CreatePowerSet(in))
		if got != want {
			
			t.Errorf("len(CreatePowerSet(%q)) == %q, want %q", in, got, want)
		}
		in = append(in, 1)
	}
	
}
