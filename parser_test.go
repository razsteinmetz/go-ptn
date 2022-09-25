package ptn

import (
	"encoding/json"
	"fmt"
	"github.com/sanity-io/litter"
	"os"
	"reflect"
	"testing"
)

// test routine using the testdata.json file
const testdata = "testdata.json"

type tsdata struct {
	Fname  string      `json:"fname"`
	Wanted TorrentInfo `json:"wanted"`
}

type tsall []tsdata

func TestParser(t *testing.T) {
	var ts tsall
	//if true {
	//	for i, fname := range testData {
	//		ts[i].Fname = fname
	//		goldenFilename := filepath.Join("testdata", fmt.Sprintf("golden_file_%03d.json", i))
	//		buf, err := os.ReadFile(goldenFilename)
	//		if err != nil {
	//			t.Fatalf("error reading golden filke: %v", err)
	//		}
	//		err = json.Unmarshal(buf, &ts[i].Wanted)
	//		if err != nil {
	//			t.Fatalf("error unmarshal on golden file: %v", err)
	//		}
	//
	//	}
	//}
	//tsjson, err := json.MarshalIndent(ts, "", "   ")
	////tsjson, err := json.Marshal(&y)
	//
	//if err != nil {
	//	t.Fatalf("Cant marshal test dataq: %v", err)
	//}
	//
	//os.Stdout.Write(tsjson)
	//err = os.WriteFile("testdata.json", tsjson, 0644)
	buf, err := os.ReadFile(testdata)
	if err != nil {
		t.Fatalf("error reading golden filke: %v", err)
	}
	err = json.Unmarshal(buf, &ts)
	if err != nil {
		t.Fatalf("error trying to unmarshal the test data: %v", err)
	}
	for i, tsline := range ts {
		fname := tsline.Fname
		wanted := tsline.Wanted
		t.Run(fmt.Sprintf("Testing Name: %s", fname), func(t *testing.T) {
			tor, err := Parse(fname)
			if err != nil {
				fmt.Println(fname)
				t.Fatalf("test %v: parser error:\n  %v", i, err)
			}
			if !reflect.DeepEqual(*tor, wanted) {
				t.Fatalf("test %v: wrong result for %q\nwant:\n  %v\ngot:\n  %v", i, fname, litter.Sdump(wanted), litter.Sdump(*tor))
			}
		})
	}
}
