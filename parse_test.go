package mdconfig

import "testing"

func TestParseText_Invalid(t *testing.T) {
	_, err := Parse([]byte("hello world"))
	if err == nil {
		t.Log(err)
	}
}

const table1 = `
|title1|title2|
|------|------|
|value1|value2|
`

func TestParseTable1_Invalid(t *testing.T) {
	_, err := Parse([]byte(table1))
	if err != nil {
		t.Log(err)
	}
}
