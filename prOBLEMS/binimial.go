package kata

import (
    "math/big"
)
func EasyLine(n int) string {
    m := int64(n)
    return new(big.Int).Binomial(2 * m, m).String()
}