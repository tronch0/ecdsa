package curve

import (
	ec_point "github.com/tronch0/ecdsa/curve/ecdsa/point"

	"github.com/tronch0/ecdsa/field"
	"math/big"
)

// y² = x³ + ax + b
type EcdsaCurve interface {
	GetA() *field.Element
	GetB() *field.Element
	GetP() *big.Int
	GetN() *big.Int
	GetG() *ec_point.Point
}

// Edward Curve - x² + y² = 1 − d ·x²·y²
//type EdwardsCurve interface {
//	GetA() *big.Int
//	GetD() *big.Int
//	GetP() *big.Int
//	GetN() *big.Int
//	GetG() *ed_point.Point
//}
