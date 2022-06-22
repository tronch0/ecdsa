package point

import (
	"github.com/tronch0/ecdsa/field"
	"math/big"
)

// Point y² = x³ + ax + b
type Point struct {
	a *field.Element
	b *field.Element
	x *field.Element
	y *field.Element
}

func New(x, y, a, b *field.Element) *Point {

	res := &Point{
		a: a,
		b: b,
		x: x,
		y: y,
	}

	// special case - point at infinity
	if x == nil && y == nil {
		return res
	}

	// y² = x³ + ax + b
	yExpo2 := y.Expo(big.NewInt(2))
	xExpo3axb := x.Expo(big.NewInt(3)).Add(a.Mul(x).Add(b))

	if yExpo2.Equal(xExpo3axb) == false {
		panic("point not on the curve")
	}

	return res
}

func (p *Point) Add(p2 *Point) *Point {
	if p.a.Equal(p2.a) == false || p.b.Equal(p2.b) == false {
		panic("points are not on the same curve")
	}

	// check if identity
	if p.isIdentity() {
		return Clone(p2)
	}

	if p2.isIdentity() {
		return Clone(p)
	}

	// P₁ == P₂
	if p.Equal(p2) {
		s := getSlopeEqualPoints(p)
		px3 := getAddX(p.x, p2.x, s)
		py3 := getAddY(p.x, px3, p.y, s)

		return &Point{
			x: px3,
			y: py3,
			a: field.Clone(p.a),
			b: field.Clone(p.b),
		}
	}

	// x₁ == x₂
	if p.x.Equal(p2.x) {
		return &Point{
			x: nil,
			y: nil,
			a: field.Clone(p.a),
			b: field.Clone(p.b),
		}
	}

	// x₁ != x₂
	s := getSlopeUnequalPoints(p, p2)
	px3 := getAddX(p.x, p2.x, s)
	py3 := getAddY(p.x, px3, p.y, s)

	return &Point{
		x: px3,
		y: py3,
		a: field.Clone(p.a),
		b: field.Clone(p.b),
	}
}

func (p *Point) ScalarMul(coef *big.Int) *Point {
	a := field.Clone(p.a)
	b := field.Clone(p.b)

	curr := Clone(p)
	res := New(nil, nil, a, b)

	for coef.Cmp(big.NewInt(0)) > 0 {
		if coef.Bit(0) == 1 {
			res = res.Add(curr)
		}
		curr = curr.Add(curr)
		coef = new(big.Int).Rsh(coef, 1)
	}

	return res
}

func (p *Point) Equal(p2 *Point) bool {
	if p.x == nil || p2.x == nil || p.y == nil || p2.y == nil {
		if p.x != nil || p2.x != nil || p.y != nil || p2.y != nil {
			panic("equal function validation error: invalid infinity point definition")
		}

		return p.a.Equal(p2.a) && p.b.Equal(p2.b)
	}

	return p.a.Equal(p2.a) && p.b.Equal(p2.b) && p.x.Equal(p2.x) && p.y.Equal(p2.y)
}

func Clone(p *Point) *Point {
	newA := field.Clone(p.a)
	newB := field.Clone(p.b)
	newX := field.Clone(p.x)
	newY := field.Clone(p.y)

	return &Point{a: newA, b: newB, x: newX, y: newY}
}

func (p *Point) Print() string {
	res := "{x: "
	if p.x == nil {
		res += "nil, "
	} else {
		res += p.x.ToString() + ", "
	}

	res += "y: "

	if p.y == nil {
		res += "nil"
	} else {
		res += p.y.ToString() + ", "
	}

	res += "a: "

	if p.a == nil {
		res += "nil"
	} else {
		res += p.a.ToString() + ", "
	}

	res += "b: "

	if p.b == nil {
		res += "nil"
	} else {
		res += p.b.ToString()
	}

	res += "}"

	return res
}

func (p *Point) GetX() *field.Element {
	return p.x
}

func (p *Point) GetY() *field.Element {
	return p.y
}

func getSlopeEqualPoints(p1 *Point) *field.Element {
	// s = (3x₁² + a) / 2y₁
	a := p1.x.Expo(big.NewInt(2)).ScalarMul(big.NewInt(3)).Add(p1.a)
	b := p1.y.ScalarMul(big.NewInt(2))

	s := a.Div(b)
	return s
}

func getSlopeUnequalPoints(p1, p2 *Point) *field.Element {
	// s = (y₂ - y₁) / (x₂ - x₁)
	a := p2.y.Sub(p1.y)
	b := p2.x.Sub(p1.x)

	s := a.Div(b)
	return s
}

func getAddX(px1, px2, s *field.Element) *field.Element {
	// x₃ = s² - x₁ - x₂
	x3 := s.Expo(big.NewInt(2)).Sub(px1).Sub(px2)

	return x3
}

func getAddY(px1, px3, py1, s *field.Element) *field.Element {
	// y₃ = s(x₁ - x₃) - y₁
	y3 := px1.Sub(px3).Mul(s).Sub(py1)

	return y3
}

func (p *Point) isIdentity() bool {
	return p.x == nil
}
