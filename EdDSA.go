package naive_ed25519

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"math/big"
)

func reverseBytes(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func hash(b []byte) []byte {
	h := sha512.New()
	h.Write(b)
	return h.Sum(nil)
}

func hashModQ(b []byte) *big.Int {
	return new(big.Int).Mod(new(big.Int).SetBytes(reverseBytes(hash(b))), GroupOrderGet())
}

func secretExpand(secret []byte) (*big.Int, []byte, error) {
	if len(secret) != 32 {
		return big.NewInt(0), []byte{}, errors.New("invalid private key")
	}
	h := hash(secret)
	a := new(big.Int).SetBytes(reverseBytes(h[:32]))
	a = new(big.Int).And(a,
		new(big.Int).Sub(new(big.Int).Lsh(new(big.Int).SetInt64(1), 254),
			new(big.Int).SetInt64(8)))
	a = new(big.Int).Or(a,
		new(big.Int).Lsh(new(big.Int).SetInt64(1), 254),
	)
	return a, h[32:], nil
}

func Signify(secret []byte, msg []byte) []byte {
	a, prefix, err := secretExpand(secret)
	if err != nil {
		fmt.Println(err)
	}
	A := PointCompress(ScalarMult(BasePointGGet(), *a))
	toHash := []byte{}
	toHash = append(toHash, prefix...)
	toHash = append(toHash, msg...)
	r := hashModQ(toHash)
	R := ScalarMult(BasePointGGet(), *r)
	Rs := PointCompress(R)
	toHash = []byte{}
	toHash = append(toHash, Rs[0:]...)
	toHash = append(toHash, A[0:]...)
	toHash = append(toHash, msg...)
	h := hashModQ(toHash)
	s := new(big.Int).Mod(new(big.Int).Add(r, new(big.Int).Mul(h, a)), GroupOrderGet())
	res := []byte{}
	res = append(res, Rs[:]...)
	b := [32]byte{}
	s.FillBytes(b[:])
	res = append(res, reverseBytes(b[:])...)
	return res
}

func Verify(public []byte, msg []byte, signature []byte) bool {
	R := PointDecompress(signature[:32])
	_ = R
	A := PointDecompress(public)
	_ = A
	S := new(big.Int).SetBytes(reverseBytes(signature[32:64]))
	_ = S
	toHash := []byte{}
	compressedR := PointCompress(R)
	toHash = append(toHash, compressedR[:]...)
	toHash = append(toHash, public...)
	toHash = append(toHash, msg...)
	first := ScalarMult(BasePointGGet(), *S)
	h := new(big.Int).SetBytes(reverseBytes(hash(toHash)))
	second := AddECPoints(R, ScalarMult(A, *h))
	if first.X.Cmp(second.X) != 0 || first.Y.Cmp(second.Y) != 0 {
		return false
	}
	return true
}
