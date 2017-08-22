package main

import (
	"testing"

	"github.com/dgraph-io/dgraph/protos"
)

func TestDecodeMsg(t *testing.T) {

	var msg = []byte{
		0x0a, 0x14, 0x09, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x12, 0x05, 0x50, 0x65, 0x74,
		0x65, 0x72, 0x20, 0x01, 0x60, 0x03, 0x22, 0x23, 0x00, 0x00, 0x00, 0x01, 0x80, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	var list protos.PostingList
	if err := list.Unmarshal(msg); err != nil {
		t.Fatal(err)
	}

	t.Logf("#%v", list)
}