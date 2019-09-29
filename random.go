package strutil

import (
	"crypto/rand"
	"math/big"
)

var randReader = rand.Reader

// Random creates new string based on strSet. It uses crypto/rand as the
// random number generator. error is the one returned by rand.Int
func Random(strSet string, length int) (string, error) {
	if length == 0 || strSet == "" {
		return "", nil
	}
	set := []rune(strSet)
	bigLen := big.NewInt(int64(len(set)))

	res := make([]rune, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(randReader, bigLen)
		if err != nil {
			return "", err
		}
		res[i] = set[n.Int64()]
	}

	return string(res), nil
}
