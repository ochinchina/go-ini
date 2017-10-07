package ini

import (
	"bytes"
	"testing"
)

func TestSection(t *testing.T) {
	data := `[section1]
    key1= value1
    [section2]
    key2=value2`

	ini := NewIni()
	ini.Load(bytes.NewBufferString(data))
	if !ini.HasSection("section1") || !ini.HasSection("section2") {
		t.Error("Fail to load ini file")
	}
}

func TestNormalKey(t *testing.T) {
	data := `[section1]
    key1 = """this is one line"""
    [section2]
    key2 = value2`
	ini := NewIni()
	ini.Load(bytes.NewBufferString(data))
	if ini.Get("section1", "key1", "") != "this is one line" || ini.Get("section2", "key2", "") != "value2" {
		t.Error("Fail to get key")
	}

}

func TestMultiLine(t *testing.T) {
	data := `[section1]
key1 = """this is a
multi line
test"""
[section2]
key2 = value2
`

	ini := NewIni()
	ini.Load(bytes.NewBufferString(data))
	key1_value := `this is a
multi line
test`
	if ini.Get("section1", "key1", "") != key1_value {
		t.Error("Fail to load ini with multi line keys")
	}
}

func TestContinuationLine(t *testing.T) {
	data := "[section1]\nkey1 = this is a \\\nmulti line \\\ntest\nkey2= this is key2\n[section2]\nkey2=value2"

	ini := NewIni()
	ini.Load(bytes.NewBufferString(data))
	if ini.Get("section1", "key1", "") != "this is a multi line test" {
		t.Error("Fail to load ini with Continuation char")
	}

}

func TestValueWithEscapeChar(t *testing.T) {
	data := "[section1]\nkey1 = this is a \\nmulti line\\ttest\nkey2= this is key2\n[section2]\nkey2=value2"
	ini := NewIni()
	ini.Load(bytes.NewBufferString(data))
	if ini.Get("section1", "key1", "") != "this is a \nmulti line\ttest" {
		t.Error("Fail to load ini with escape char")
	}
}
