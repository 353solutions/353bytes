package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	urlTemplate = "https://api.stocktwits.com/api/2/streams/symbol/%s.json"
)

// RelatedSymbols returns list of related symbols and their counts
func RelatedSymbols(symbol string) (map[string]int, error) {
	url := fmt.Sprintf(urlTemplate, symbol)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	// Anonymous structure to get only list of symbols per message
	var reply struct {
		Messages []struct {
			Symbols []struct {
				Symbol string
			}
		}
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	for _, msg := range reply.Messages {
		for _, sym := range msg.Symbols {
			if sym.Symbol != symbol {
				counts[sym.Symbol] += 1
			}
		}
	}

	return counts, nil
}

func main() {
	syms, err := RelatedSymbols("NFLX")
	if err != nil {
		log.Fatal(err)
	}

	for sym, n := range syms {
		fmt.Printf("%-5s %d\n", sym, n)
	}
}
