package ini

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSection(t *testing.T) {
	data := `[section1]
    key1= value1
    [section2]
    key2=value2`

	ini := Load(data)
	if !ini.HasSection("section1") || !ini.HasSection("section2") {
		t.Error("Fail to load ini file")
	}
}

func TestNormalKey(t *testing.T) {
	data := `[section1]
    key1 = """this is one line"""
    [section2]
    key2 = value2`
	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is one line" || ini.GetValueWithDefault("section2", "key2", "") != "value2" {
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

	ini := Load(data)
	key1_value := `this is a
multi line
test`
	if ini.GetValueWithDefault("section1", "key1", "") != key1_value {
		t.Error("Fail to load ini with multi line keys")
	}
}

func TestContinuationLine(t *testing.T) {
	data := "[section1]\nkey1 = this is a \\\nmulti line \\\ntest\nkey2= this is key2\n[section2]\nkey2=value2"

	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a multi line test" {
		t.Error("Fail to load ini with Continuation char")
	}

	data = "[section1]\nkey1 = this is a line end with \\\\\nkey2= this is key2\n[section2]\nkey2=value2"
	ini = Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a line end with \\" {
		t.Error("Fail to load ini without Continuation char")
	}
}

func TestValueWithEscapeChar(t *testing.T) {
	data := "[section1]\nkey1 = this is a \\nmulti line\\ttest\nkey2= this is key2\n[section2]\nkey2=value2"
	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a \nmulti line\ttest" {
		t.Error("Fail to load ini with escape char")
	}
}

func TestToEscape(t *testing.T) {
	data := "string with escape char \r\n\t;# for testing"
	new_data := "string with escape char \\r\\n\\t\\;\\# for testing"
	if toEscape(data) != new_data {
		t.Error("Fail to convert escape string")
	}
}
func TestInlineComments(t *testing.T) {
	//inline comments must be start with ; or # and a space char before it
	data := "[section1]\nkey1 = this is a inline comment test ; comments ; do you know\nkeys=this is key2\n[section2]\nkey3=value3"
	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a inline comment test" {
		t.Error("Fail to load ini with inline comments")
	}

	data = "[section1]\nkey1 = this is a inline comment test;comments\nkeys=this is key2\n[section2]\nkey3=value3"
	ini = Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a inline comment test;comments" {
		t.Error("Fail to load ini with inline comments")
	}

	data = "[section1]\nkey1 = this is not a inline comment test \\;comments\nkeys=this is key2\n[section2]\nkey3=value3"

	ini = Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is not a inline comment test ;comments" {
		t.Error("Fail to load ini without inline comments")
	}
}

func TestOctInValue(t *testing.T) {
	data := "[section1]\nkey1=this is \\141 oct test"
	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a oct test" {
		t.Error("fail to convert oct to char")
	}
}

func TestUnicodeValue(t *testing.T) {
	data := "[section1]\nkey1=this is \\x0061 unicode test"
	ini := Load(data)
	if ini.GetValueWithDefault("section1", "key1", "") != "this is a unicode test" {
		t.Error("fail to convert unicode to char")
	}
}

func TestIniWriteRead(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0))
	for i := 0; i < 100; i++ {
		sectionName := fmt.Sprintf("section_%d", i)
		fmt.Fprintf(buf, "[%s]\n", sectionName)
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key_%d_%d", i, j)
			value := fmt.Sprintf("value_%d_%d", i, j)
			fmt.Fprintf(buf, "%s=%s\n", key, value)
		}
	}

	ini := Load(buf.String())
	ini = Load(ini.String())
	if len(ini.Sections()) != 100 {
		t.Error("fail to write&load ini")
	}

	for i := 0; i < 100; i++ {
		sectionName := fmt.Sprintf("section_%d", i)
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key_%d_%d", i, j)
			value := fmt.Sprintf("value_%d_%d", i, j)
			if v, err := ini.GetValue(sectionName, key); err != nil || v != value {
				t.Error("fail to write&load ini")
			}
		}
	}
}
