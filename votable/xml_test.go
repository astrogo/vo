package votable

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

func TestDecoding(t *testing.T) {
	for i, test := range []struct {
		fname string
	}{
		{
			fname: "testdata/test.vot",
		},
	} {
		f, err := os.Open(test.fname)
		if err != nil {
			t.Fatalf("test #%d: error opening [%s]: %v\n", i, test.fname, err)
		}
		defer f.Close()

		var vo xmlVO
		err = xml.NewDecoder(f).Decode(&vo)
		if err != nil {
			t.Fatalf("test #%d: error decoding [%s]: %v\n", i, test.fname, err)
		}

		fmt.Printf("test #%d: %#v\n", i, vo)

	}
}
