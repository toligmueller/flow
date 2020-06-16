package flow

func (bs Buckets) Start() bool {
	for _, b := range bs {
		if !b.Start() {
			return false
		}
	}
	return true
}

func (bs Buckets) Stop() bool {
	for _, b := range bs {
		if !b.Stop() {
			return false
		}
	}
	return true
}

func (bs Buckets) Consume(amt int) bool {
	for _, b := range bs {
		if !b.Consume(amt) {
			return false
		}
	}
	return true
}

func (bs Buckets) ConsumeWithTimeout(amt int) bool {
	for _, b := range bs {
		b.ConsumeWithTimeout(amt)
	}
	return true
}
