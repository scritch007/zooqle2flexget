package main

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestZooqleExtract(t *testing.T) {
	f, err := os.Open("test_examples/zooqle.xml")
	toFind := "https://zooqle.com/download/vlp1u.torrent"
	if err != nil {
		t.Fatalf("Failed to open xml test file with error %v", err)
	}

	dec := xml.NewDecoder(f)
	res := ZXML{}
	err = dec.Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode xml test file with error %v", err)
	}

	if len(res.Channel.Items) == 0 {
		t.Fatalf("Didn't find any items in the list")
	}

	for _, i := range res.Channel.Items {
		if i.GUID.URL == toFind {
			return
		}

	}
	t.Fatalf("Failed to find the expected URL %s", toFind)
}
