package hello

import "crypto/rand"

const (
	// base58
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	size     = 11
)

// NewID generates a random base-58 ID.
func NewID() string {

	var id = make([]byte, size)
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}

	for i, p := range id {
		id[i] = alphabet[int(p)%len(alphabet)] // discard everything but the least significant bits
	}

	return string(id)
}
