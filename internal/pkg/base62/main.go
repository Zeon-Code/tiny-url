package base62

const base = 62

var charset = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

var index = func() map[byte]int {
	m := make(map[byte]int, len(charset))
	for i, c := range charset {
		m[c] = i
	}
	return m
}()

// Encode converts a positive int64 into a base62 string.
func Encode(n int64) string {
	if n == 0 {
		return string(charset[0])
	}

	var buf [11]byte
	i := len(buf)

	for n > 0 {
		i--
		buf[i] = charset[n%base]
		n /= base
	}

	return string(buf[i:])
}

// Decode converts a base62 string back into an int64.
// It panics on invalid characters.
func Decode(s string) int64 {
	var n int64

	for i := 0; i < len(s); i++ {
		val, ok := index[s[i]]
		if !ok {
			panic("invalid base62 character")
		}
		n = n*base + int64(val)
	}

	return n
}
