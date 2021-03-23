package dupfilter

import (
	"fmt"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	filter := &Filter{}
	filter.Run()
	go filter.CheckIn("hi")

	for {
		s := <-filter.out
		fmt.Printf("Get: %s\n", s)
		// t.Errorf("Get: %s\n", s)
		if s != "" {
			go filter.CheckIn(s + "i")
		}
		time.Sleep(1 * time.Second)
	}
}
