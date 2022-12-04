package naive_ed25519

import (
	"fmt"
	"math/big"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// G-generator receiving
func BasePointGGet() (point ECPoint) {
	x := new(big.Int)
	x.SetString("216936D3CD6E53FEC0A4E231FDD6DC5C692CC7609525A7B2C9562D608F25D51A", 16)
	y := new(big.Int)
	y.SetString("6666666666666666666666666666666666666666666666666666666666666658", 16)
	return ECPoint{x, y}
}

// ECPoint creation with pre-defined parameters
func ECPointGen(x, y *big.Int) (point ECPoint) {
	return ECPoint{
		X: x,
		Y: y,
	}
}

// P ∈ CURVE?
// −x^2 + y^2 = 1 − (121665/121666) * x^2 * y^2
// (-x*x + y*y - 1 - d*x*x*y*y) % q == 0
func IsOnCurveCheck(a ECPoint) (c bool) {
	zero := big.NewInt(0)
	negative_one := big.NewInt(-1)
	xx := new(big.Int).Mul(a.X, a.X)
	negative_xx := new(big.Int).Neg(xx)
	yy := new(big.Int).Mul(a.Y, a.Y)
	negative_dxxyy := new(big.Int).Neg(new(big.Int).Mul(
		new(big.Int).Mul(constantdGet(), xx),
		yy))
	res := new(big.Int).Add(
		new(big.Int).Add(negative_xx, yy),
		new(big.Int).Add(negative_one, negative_dxxyy),
	)
	return new(big.Int).Mod(res, fieldCharacteristicGet()).Cmp(zero) == 0
}

// P + Q
func AddECPoints(a, b ECPoint) (c ECPoint) {
	q := fieldCharacteristicGet()
	x1x2 := new(big.Int).Mul(a.X, b.X)
	y1y2 := new(big.Int).Mul(a.Y, b.Y)
	dx1x2y1y2 := new(big.Int).Mul(constantdGet(), new(big.Int).Mul(x1x2, y1y2))
	one := new(big.Int).SetUint64(1)
	c.X = new(big.Int).Mul(
		new(big.Int).Add(
			new(big.Int).Mul(a.X, b.Y),
			new(big.Int).Mul(b.X, a.Y),
		),
		inverse(new(big.Int).Add(one, dx1x2y1y2)))
	// y3 = (y1*y2+x1*x2) * inv(1-d*x1*x2*y1*y2)
	c.Y = new(big.Int).Mul(
		new(big.Int).Add(
			new(big.Int).Mul(a.Y, b.Y),
			new(big.Int).Mul(a.X, b.X),
		),
		inverse(new(big.Int).Sub(one, dx1x2y1y2)))
	// return [x3 % q,y3 % q]
	c.X = c.X.Mod(c.X, q)
	c.Y = c.Y.Mod(c.Y, q)
	return c
}

// 2 * P
func DoubleECPoints(a ECPoint) (c ECPoint) {
	return AddECPoints(a, a)
}

// k * P
func ScalarMult(a ECPoint, k big.Int) (c ECPoint) {
	zero := new(big.Int).SetInt64(0)
	one := new(big.Int).SetInt64(1)
	two := new(big.Int).SetInt64(2)
	if k.Cmp(zero) == 0 {
		return ECPointGen(new(big.Int).SetInt64(0), new(big.Int).SetInt64(1))
	}
	Q := ScalarMult(a, *new(big.Int).Div(&k, two))
	Q = AddECPoints(Q, Q)
	if new(big.Int).And(&k, one).Cmp(one) == 0 {
		Q = AddECPoints(Q, a)
	}
	return Q
}

// Convert point to string
func ECPointToString(point ECPoint) (s string) {
	return fmt.Sprintf("ECPoint(%d,%d)", point.X, point.X)
}

// Print point
func PrintECPoint(point ECPoint) {
	fmt.Print(ECPointToString(point))
}

// 2**255 - 19
func fieldCharacteristicGet() (p *big.Int) {
	p, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed", 16)
	return p
}

// -121665 * inv(121666)
func constantdGet() (d *big.Int) {
	d, _ = new(big.Int).SetString("-4513249062541557337682894930092624173785641285191125241628941591882900924598840740", 10)
	return d
}

func inverse(a *big.Int) (b *big.Int) {
	q := fieldCharacteristicGet()
	return new(big.Int).Exp(a, new(big.Int).Sub(q, new(big.Int).SetInt64(2)), q)
}
