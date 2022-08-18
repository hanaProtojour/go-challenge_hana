package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// type Seed struct {
// 	Seed string `json:"seed"`
// }

// func (s Seed) String() string {
// 	return fmt.Sprintf(s.Seed)
// }

type Seeds struct {
	Seeds []string `json:"seeds"`
}

type Seed struct {
	Seed            string `json:"seed"`
	hash            [32]byte
	hexadecimalhash string
}

func hash(s string) [32]byte {
	sum := sha256.Sum256([]byte(s))
	return sum
}

func sha256ToHexString(hashedvalue [32]byte) string {
	sliced := hashedvalue[:]
	mdStr := hex.EncodeToString(sliced)
	return mdStr
}

func hexSha256(s string) string {
	hash3 := hash(s)
	hex := sha256ToHexString(hash3)
	return hex
}

func DecodeValidateHashJSON(w http.ResponseWriter, r *http.Request) {
	var Seeds Seeds
	err := json.NewDecoder(r.Body).Decode(&Seeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashSeedsAndPrintToChannel(Seeds)

}

func hashSeedsAndPrintToChannel(seeds Seeds) {

	channel := make(chan string)

	var wg sync.WaitGroup

	for _, stringseed := range seeds.Seeds {
		// var seed Seed
		wg.Add(1)

		seed2 := stringseed

		go func() {
			defer wg.Done()

			go func() { channel <- hexSha256(seed2) }()

		}()

		// seed.Seed = &seed2
		// seed.hexadecimalhash = hexSha256(seed2)

	}

	wg.Wait()

	s3 := <-channel
	fmt.Println(s3)
	s2 := <-channel
	fmt.Println(s2)
	s1 := <-channel
	fmt.Println(s1)
	// s0 := <-channel
	// fmt.Println(s0)

}
