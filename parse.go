package mdconfig

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gomarkdown/markdown"
)

type Config struct {
}

func Parse(text []byte) (*Config, error) {
	ast := markdown.Parse(text, nil)
	fmt.Printf("%#v", ast)
	return nil, nil
}

func ParseFrom(inp io.Reader) (*Config, error) {
	var buf bytes.Buffer
	n, err := buf.ReadFrom(inp)
	if err != nil {
		return nil, fmt.Errorf("i/o error after reading %d bytes: %w", n, err)
	}
	return Parse(buf.Bytes())
}
