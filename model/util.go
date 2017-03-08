package model

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CommaSepStringToIntArray(commaSep string) []int {
	stringArray := strings.Split(commaSep, ",")
	intArray := make([]int, len(stringArray), len(stringArray))
	for key, value := range stringArray {
		tempInt64, _ := strconv.ParseInt(value, 10, 32)
		intArray[key] = int(tempInt64)
	}
	return intArray
}

func generateRandomString(length int) string {
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rs := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	r := ""
	for i, _ := range b {
		b[i] = pool[rs.Intn(len(pool))]
	}
	r = string(b)
	return r
}
