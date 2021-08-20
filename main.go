package main

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/go-spatial/geom/encoding/geojson"
	"github.com/go-spatial/geom/encoding/wkt"
	"github.com/pkg/errors"
)

func main() {
	r := bufio.NewReader(os.Stdin)

ReadLoop:
	for {
		line, err := r.ReadBytes(0x0A)
		if err == io.EOF {
			break ReadLoop
		}
		if err != nil {
			panic(err)
		}
		out, err := convert(line)
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(out)
	}
}

func convert(in []byte) ([]byte, error) {
	g, err := wkt.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, errors.Wrap(err, "decoding wkt")
	}
	out, err := geojson.Marshal(g)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling geojson")
	}
	return out, nil
}
