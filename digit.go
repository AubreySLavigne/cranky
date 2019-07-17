package main

type digit struct {
	value int
	known bool
}

func newDigit() digit {
	return digit{
		value: 0,
		known: false,
	}
}

func (d digit) set(v int) digit {
	d.known = true
	d.value = v
	return d
}
