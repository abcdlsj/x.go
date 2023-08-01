package xutil

type Batch struct {
	start, end int
	step       int
}

func NewBatch(start, end, step int) *Batch {
	return &Batch{
		start: start,
		end:   end,
		step:  step,
	}
}

func (b *Batch) Next() (start, end int) {
	start = b.start
	end = b.start + b.step
	if end > b.end {
		end = b.end
	}
	b.start = end
	return
}

func (b *Batch) HasNext() bool {
	return b.start < b.end
}
