package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

const alphabet = " abcdefghijklmnopqrstuvxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomNullString(n int) sql.NullString {
	return sql.NullString{String: RandomString(n), Valid: true}
}

func RandomTimestamp() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000

	randomNow := time.Unix(randomTime, 0)

	return randomNow
}
