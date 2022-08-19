package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"sync"
)

type myJSON struct {
	Hashes []string
}

type Seeds struct {
	Seeds []string `json:"seeds"`
}

type Seed struct {
	Seed            *string `json:"seed"`
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

	hashSeedsAndPrintToChannel(w, Seeds)

}

func hashSeedsAndPrintToChannel(w http.ResponseWriter, seeds Seeds) {

	channel := make(chan Seed)

	var wg sync.WaitGroup

	for _, stringseed := range seeds.Seeds {
		var seed Seed
		wg.Add(1)
		seed2 := stringseed

		go func() {
			defer wg.Done()
			seed.hexadecimalhash = hexSha256(seed2)
			go func() { channel <- seed }()

		}()

	}

	wg.Wait()

	s3 := <-channel
	s2 := <-channel
	s1 := <-channel

	hashedseeds := []string{}
	hashedseeds = append(hashedseeds, s1.hexadecimalhash, s2.hexadecimalhash, s3.hexadecimalhash)

	jsondat := &myJSON{Hashes: hashedseeds}
	err := json.NewEncoder(w).Encode(jsondat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
