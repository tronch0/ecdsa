package point

import (
	"github.com/tronch0/ecdsa/field"
	"math/big"
	"testing"
)

func TestCreation1(t *testing.T) {
	_ = New(
		field.New(big.NewInt(192), big.NewInt(223)), // x
		field.New(big.NewInt(105), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)
}

func TestCreation2(t *testing.T) {
	_ = New(
		field.New(big.NewInt(17), big.NewInt(223)), // x
		field.New(big.NewInt(56), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)
}

func TestCreationNotOnCurve1(t *testing.T) {
	defer func() { recover() }()

	_ = New(
		field.New(big.NewInt(42), big.NewInt(223)), // x
		field.New(big.NewInt(99), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	// Never reaches here if panics (expected panic).
	t.Errorf("did not panic")
}

func TestCreationNotOnCurve2(t *testing.T) {
	defer func() { recover() }()

	_ = New(
		field.New(big.NewInt(200), big.NewInt(223)), // x
		field.New(big.NewInt(119), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	// Never reaches here if panics (expected panic).
	t.Errorf("did not panic")
}

func TestAdd1(t *testing.T) {

	p1 := New(
		field.New(big.NewInt(170), big.NewInt(223)), // x
		field.New(big.NewInt(142), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)
	p2 := New(
		field.New(big.NewInt(60), big.NewInt(223)),  // x
		field.New(big.NewInt(139), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	expectedP := New(
		field.New(big.NewInt(220), big.NewInt(223)), // x
		field.New(big.NewInt(181), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p3 := p1.Add(p2)

	//fmt.Printf("p3: %s \n expectedRes: %s", p3.Print(), expectedP.Print())
	if p3.Equal(expectedP) == false {
		t.Fatal("assert result error")
	}

}

func TestAdd2(t *testing.T) {

	p1 := New(
		field.New(big.NewInt(47), big.NewInt(223)), // x
		field.New(big.NewInt(71), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)
	p2 := New(
		field.New(big.NewInt(17), big.NewInt(223)), // x
		field.New(big.NewInt(56), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	expectedP := New(
		field.New(big.NewInt(215), big.NewInt(223)), // x
		field.New(big.NewInt(68), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p3 := p1.Add(p2)

	//fmt.Printf("p3: %s \n expectedRes: %s", p3.Print(), expectedP.Print())
	if p3.Equal(expectedP) == false {
		t.Fatal("assert result error")
	}
}

func TestAdd3(t *testing.T) {

	p1 := New(
		field.New(big.NewInt(143), big.NewInt(223)), // x
		field.New(big.NewInt(98), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)
	p2 := New(
		field.New(big.NewInt(76), big.NewInt(223)), // x
		field.New(big.NewInt(66), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	expectedP := New(
		field.New(big.NewInt(47), big.NewInt(223)), // x
		field.New(big.NewInt(71), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	p3 := p1.Add(p2)

	if p3.Equal(expectedP) == false {
		t.Fatal("assert result error")
	}

}

func TestAdd4(t *testing.T) {

	p1 := New(
		field.New(big.NewInt(192), big.NewInt(223)), // x
		field.New(big.NewInt(105), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p2 := New(
		field.New(big.NewInt(17), big.NewInt(223)), // x
		field.New(big.NewInt(56), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	expectedP := New(
		field.New(big.NewInt(170), big.NewInt(223)), // x
		field.New(big.NewInt(142), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p3 := p1.Add(p2)

	if p3.Equal(expectedP) == false {
		t.Fatal("assert result error")
	}

}

func TestAdd5(t *testing.T) {

	p1 := New(
		field.New(big.NewInt(143), big.NewInt(223)), // x
		field.New(big.NewInt(98), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p2 := New(
		field.New(big.NewInt(76), big.NewInt(223)), // x
		field.New(big.NewInt(66), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	expectedP := New(
		field.New(big.NewInt(47), big.NewInt(223)), // x
		field.New(big.NewInt(71), big.NewInt(223)), // y
		field.New(big.NewInt(0), big.NewInt(223)),  // a
		field.New(big.NewInt(7), big.NewInt(223)),  // b
	)

	p3 := p1.Add(p2)

	if p3.Equal(expectedP) == false {
		t.Fatal("assert result error")
	}

}

func TestClone1(t *testing.T) {
	p1 := New(
		field.New(big.NewInt(143), big.NewInt(223)), // x
		field.New(big.NewInt(98), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p1Clone := Clone(p1)

	if p1Clone.Equal(p1) == false {
		t.Fatal("assert result error - p1 clone are not equal in value")
	}

	if &p1Clone.x == &p1.x {
		t.Fatal("assert result error - p1 and p1 clone share the same address ")
	}

	if &p1Clone.y == &p1.y {
		t.Fatal("assert result error - p1 and p1 clone share the same address ")
	}

	if &p1Clone.a == &p1.a {
		t.Fatal("assert result error - p1 and p1 clone share the same address ")
	}

	if &p1Clone.b == &p1.b {
		t.Fatal("assert result error - p1 and p1 clone share the same address ")
	}
}

func TestEqual1(t *testing.T) {
	p1 := New(
		field.New(big.NewInt(143), big.NewInt(223)), // x
		field.New(big.NewInt(98), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	p2 := New(
		field.New(big.NewInt(143), big.NewInt(223)), // x
		field.New(big.NewInt(98), big.NewInt(223)),  // y
		field.New(big.NewInt(0), big.NewInt(223)),   // a
		field.New(big.NewInt(7), big.NewInt(223)),   // b
	)

	if p2.Equal(p1) == false {
		t.Fatal("assert result error - p1 and p2 are not equal in value")
	}
}

func TestScalarMul(t *testing.T) {
	p := big.NewInt(223)
	a := big.NewInt(0)
	b := big.NewInt(7)

	testData := []struct {
		coef      *big.Int
		p1X       *big.Int
		p1Y       *big.Int
		expectedX *big.Int
		expectedY *big.Int
	}{
		{coef: big.NewInt(2), p1X: big.NewInt(192), p1Y: big.NewInt(105), expectedX: big.NewInt(49), expectedY: big.NewInt(71)},

		{coef: big.NewInt(2), p1X: big.NewInt(143), p1Y: big.NewInt(98), expectedX: big.NewInt(64), expectedY: big.NewInt(168)},
		{coef: big.NewInt(2), p1X: big.NewInt(47), p1Y: big.NewInt(71), expectedX: big.NewInt(36), expectedY: big.NewInt(111)},
		{coef: big.NewInt(4), p1X: big.NewInt(47), p1Y: big.NewInt(71), expectedX: big.NewInt(194), expectedY: big.NewInt(51)},
		{coef: big.NewInt(8), p1X: big.NewInt(47), p1Y: big.NewInt(71), expectedX: big.NewInt(116), expectedY: big.NewInt(55)},
		{coef: big.NewInt(21), p1X: big.NewInt(47), p1Y: big.NewInt(71), expectedX: nil, expectedY: nil},
	}

	for i := range testData {

		p1 := New(
			field.New(testData[i].p1X, p), // x
			field.New(testData[i].p1Y, p), // y
			field.New(a, p),               // a
			field.New(b, p),               // b
		)

		product := p1.ScalarMul(testData[i].coef)

		var expectedRes *Point
		if testData[i].expectedX == nil {
			expectedRes = New(
				nil,             // x
				nil,             // y
				field.New(a, p), // a
				field.New(b, p), // b
			)
		} else {
			expectedRes = New(
				field.New(testData[i].expectedX, p), // x
				field.New(testData[i].expectedY, p), // y
				field.New(a, p),                     // a
				field.New(b, p),                     // b
			)
		}

		if product.Equal(expectedRes) == false {
			t.Fatal("assert result error")
		}
	}
}

//n, _ := new(big.Int).SetString("0xffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551", 16)
//a, _ := new(big.Int).SetString("0xffffffff00000001000000000000000000000000fffffffffffffffffffffffc", 16)
//b, _ := new(big.Int).SetString("0x5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", 16)
//x, _ := new(big.Int).SetString("0x6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", 16)
//y, _ := new(big.Int).SetString("0x4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5", 16)
