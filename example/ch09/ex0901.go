package main

import ( //liststart1
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
) //listend1

func main() {
	r := seedRand()
	fmt.Println(r.Int())
}

func seedRand() *rand.Rand { //liststart2
	var b [8]byte
	_, err := crand.Read(b[:]) // Readにはスライスが渡る。詳しくは ex0901x.go参照
	if err != nil {
		panic("エラー：暗号論的擬似乱数生成器のシード")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
} //listend2
