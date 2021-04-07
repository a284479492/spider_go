package dupfilter

import (
	"crawler/types"
	"fmt"
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	filter := &Filter{}
	filter.Run()
	go filter.CheckIn(types.Request{
		URL: "abc.com",
	})

	for {
		s := <-filter.out
		fmt.Printf("Get: %s\n", s)
		// t.Errorf("Get: %s\n", s)
		if s.URL != "" {
			go filter.CheckIn(types.Request{
				URL: s.URL + "i",
			})
		}
		time.Sleep(1 * time.Second)
	}
}
