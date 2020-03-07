type ModInt struct {
	x, m int
}

func NewModInt(x, m int) *ModInt {
	mi := ModInt{x, m}
	mi.x %= m
	return &mi
}

func (m *ModInt) add(a *ModInt) *ModInt {
	x := (m.x + a.x) % m.m
	return &ModInt{x, m.m}
}