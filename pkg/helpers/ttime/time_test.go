package ttime

import (
	"fmt"
	"testing"
)

func TestRealTime_Now(t *testing.T) {
	rt := &RealTime{}
	time := rt.Now()
	fmt.Println(time)
	fmt.Println(rt.Timestamp())
	fmt.Println(rt.Pretty(true))

	time = rt.Now(true)
	fmt.Println(time)
}
