package commitchain

import "time"

type EvictionPolicy struct {
	// lazy eviction is proceed
	lazy bool

	atLeastSegments int
	leastDuration   time.Duration
}

func (ep *EvictionPolicy) DoYourThing() bool {

}
