package ini

import (
    "bytes"
    "fmt"
    "io"
)
type Section struct {
	Name      string
	KeyValues map[string]Key
}

func NewSection(name string) *Section {
	return &Section{Name: name,
		KeyValues: make(map[string]Key)}
}

func (section *Section) Add(key, value string) {
    section.KeyValues[key] = NewNormalKey(key, value )
}

func (section *Section) HasKey(key string) bool {
	_, ok := section.KeyValues[key]
	return ok
}

// Get all the keys in the section
//
// return: all keys in the section
func (section *Section) Keys() []Key{
	r := make([]Key, 0)
	for _, v := range section.KeyValues {
		r = append(r, v)
	}
	return r
}

// Get the key by name
//
func (section *Section)Key(key string) Key {
    if v, ok := section.KeyValues[key]; ok {
        return v
    }
    return NewNonExistKey( key )
}

// Get value of key
func (section *Section)Value(key string)(string, error ) {
    return section.Key( key ).Value()
}

func (section *Section) ValueWithDefault(key string, defValue string) string {
    return section.Key( key ).ValueWithDefault( defValue )
}

func (section *Section)Int( key string )(int, error ) {
    return section.Key(key).Int()
}

func (section *Section) IntWithDefault(key string, defValue int) int {
    return section.Key( key ).IntWithDefault( defValue )
}

func (section *Section) Int64( key string )(int64, error ) {
    return section.Key( key ).Int64()
}
func (section *Section) Int64WithDefault(key string, defValue int64) int64 {
    return section.Key( key ).Int64WithDefault( defValue )
}

func (section *Section) Uint64( key string )(uint64, error ) {
    return section.Key( key ).Uint64()
}

func (section *Section) Uint64WithDefault(key string, defValue uint64) uint64 {
    return section.Key( key ).Uint64WithDefault( defValue )
}

func (section *Section) Float32( key string) (float32, error ) {
    return section.Key( key ).Float32()
}

func (section *Section) Float32WithDefault(key string, defValue float32) float32 {
    return section.Key( key ).Float32WithDefault( defValue )
}

func (section *Section) Float64( key string) (float64, error ) {
    return section.Key( key ).Float64()
}

func (section *Section) Float64WithDefault(key string, defValue float64) float64 {
    return section.Key( key ).Float64WithDefault( defValue )
}

func (section *Section)String() string {
    buf := bytes.NewBuffer( make([]byte, 0 ) )
    section.Write( buf )
    return buf.String()
}

func (section *Section) Write( writer io.Writer ) error {
    _, err := fmt.Fprintf( writer, "[%s]\n", section.Name )
    if err != nil {
        return err
    }
    for _, v := range section.KeyValues {
        _, err = fmt.Fprintf( writer, "%s\n", v.String() )
        if err != nil {
            return err
        }
    }
    return nil
}
