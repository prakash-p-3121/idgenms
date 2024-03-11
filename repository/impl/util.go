package impl

import (
	"fmt"
	"log"
	"math/big"
)

func getNextID(idBytes []byte) ([]byte, int) {

	lastIDBigInt := new(big.Int)
	lastIDBigInt = lastIDBigInt.SetBytes(idBytes)

	addBigInt := new(big.Int)
	addBigInt = addBigInt.SetUint64(1)

	nextIDBigInt := lastIDBigInt.Add(lastIDBigInt, addBigInt)
	log.Println(nextIDBigInt.String())
	binaryString := fmt.Sprintf("%b", nextIDBigInt)
	return nextIDBigInt.Bytes(), len(binaryString)
}
