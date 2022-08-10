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
	{
		var b [8]byte
		fmt.Printf("%v %T\n", b, b)			
		_, err := crand.Read(b[:])
		fmt.Printf("%v %T\n", b[:], b[:])

		if err != nil {
			panic("エラー：暗号論的擬似乱数生成器のシード")
		}
		r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
		fmt.Println(r.Int())
		// return r
	}

	{ // 下のように書くほうが直感的と思われる
		var b = make([]byte, 8)  
		fmt.Printf("%v %T\n", b, b)			
		_, err := crand.Read(b)
		fmt.Printf("%v %T\n", b, b)			

		if err != nil {
			panic("エラー：暗号論的擬似乱数生成器のシード")
		}
		r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
		// fmt.Println(r.Int())
		return r
	}

} //listend2
