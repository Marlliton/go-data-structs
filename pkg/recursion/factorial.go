package recursion

import "math/big"

func Factorial(n int64) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, Factorial(n-1))
}
