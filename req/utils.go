package req

import (
	"math/rand"
	"strings"
)

func SpawnUUID(length int) string {
	characters := "ABCDEFGHJKMNP9gqQRSToOLVvI1lWXYZabcdefhijkmnprstwxyz2345678"
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(characters))
		sb.WriteByte(characters[index])
	}

	return sb.String()
}
