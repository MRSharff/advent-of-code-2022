package metrics

import (
	"fmt"
	"time"
)

func Timeit(f func()) {
	start := time.Now()
	defer func(){ fmt.Println(time.Since(start))}()
	f()
}