package hash

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type seeds struct {
	Seeds []string `json:"seeds"`
}

func hash(s string) [32]byte {
	sum := sha256.Sum256([]byte(s))
	return sum
}

func DecodeValidateHashJSON(w http.ResponseWriter, r *http.Request) {
	var seeds seeds
	err := json.NewDecoder(r.Body).Decode(&seeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashSeedsAndPrintToChannel(seeds)

}

func hashSeedsAndPrintToChannel(seeds seeds) {

	channel := make(chan [32]byte, 3)

	var wg sync.WaitGroup

	for _, seed := range seeds.Seeds {
		wg.Add(1)

		seed2 := seed

		go func() {
			defer wg.Done()

			go func() { channel <- hash(seed2) }()

		}()

	}

	wg.Wait()

	seed3 := <-channel
	fmt.Println(seed3)
	seed2 := <-channel
	fmt.Println(seed2)
	seed1 := <-channel
	fmt.Println(seed1)

}
