package logic

import (
	"math/big"
)

//PrintBigInt  must provide an int64 value and will return a big.Int Pointer
func PrintBigInt(data int64)  (d *big.Int) {
 return big.NewInt(data)
}