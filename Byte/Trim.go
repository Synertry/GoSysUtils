package Byte

import "bytes"

// TrimLeftUntil the first occurence of char.
// May create out of bounds error?
func TrimLeftUntil(b []byte, char byte) []byte {
	return b[bytes.IndexByte(b, char):]
}
