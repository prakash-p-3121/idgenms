package impl

import "math/big"

func getNextID(idBytes []byte) []byte {

	lastIDBigInt := new(big.Int)
	lastIDBigInt = lastIDBigInt.SetBytes(idBytes)

	addBigInt := new(big.Int)
	addBigInt = addBigInt.SetUint64(1)

	nextIDBigInt := lastIDBigInt.Add(lastIDBigInt, addBigInt)
	return nextIDBigInt.Bytes()
}
