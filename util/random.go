package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a random integer between min ans max
func RandomInt(min,max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i<n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomWoner generator a random owner name
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney genrates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0,100)
}

//RandomCurrency generates a random currecny code
func RandomCurrency() string {
	currencies := []string{"EUR","USD","CAD","LKR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

