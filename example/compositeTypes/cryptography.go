package main;

import("fmt"
	"crypto/sha256")

func main() {
	fmt.Println("welcome to cryptography hash using SHA256 package")

	firstHash := sha256.Sum256([]byte("x"))
	secondHash := sha256.Sum256([]byte("X"))
	numberOfBitsThatAreDifferent, err := arrayBitComparisonCounter(&firstHash, &secondHash)
	fmt.Println(numberOfBitsThatAreDifferent)
	fmt.Println(err)

}
	
func arrayBitComparisonCounter(firstArray *[32]uint8, secondArray *[32]uint8) (numberOfBits int, error string) {
	numberOfBits = 0
	for index := range firstArray {
		if firstArray[index] != secondArray[index] {
				numberOfBits++
		}
	}
	if len(firstArray) != 32 || len(secondArray) != 32 {
		return 0, "error, invalid hash"
	}
	return numberOfBits, ""
}