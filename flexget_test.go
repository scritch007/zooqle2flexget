package main

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestFlexGetExtract(t *testing.T) {
	f, err := os.Open("test_examples/flexget.xml")
	toFind := "http://tvunderground.org.ru/torrent.php?tid=264013"
	if err != nil {
		t.Fatalf("Failed to open xml test file with error %v", err)
	}

	dec := xml.NewDecoder(f)
	res := FXML{}
	err = dec.Decode(&res)
	if err != nil {
		t.Fatalf("Failed to decode xml test file with error %v", err)
	}

	if len(res.Channel.Items) == 0 {
		t.Fatalf("Didn't find any items in the list")
	}

	for _, i := range res.Channel.Items {
		if i.GUID == toFind {
			return
		}

	}
	t.Fatalf("Failed to find the expected URL %s", toFind)
}
