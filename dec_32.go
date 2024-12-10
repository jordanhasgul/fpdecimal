package fpdecimal

type Dec32 struct {
	bits uint32
}

func NewDec32(f float32) *Dec32 {
	panic("unimplemented")
}

func (d *Dec32) Bits() uint32 {
	return d.bits
}
