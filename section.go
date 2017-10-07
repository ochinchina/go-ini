package ini

import (
	"strconv"
)

type Section struct {
	Name      string
	KeyValues map[string]string
}

func NewSection(name string) *Section {
	return &Section{Name: name,
		KeyValues: make(map[string]string)}
}

func (section *Section) Add(key, value string) {
	section.KeyValues[key] = value
}

func (section *Section) HasKey(key string) bool {
	_, ok := section.KeyValues[key]
	return ok
}

// Get all the keys in the section
//
// return: all keys in the section
func (section *Section) Keys() []string {
	r := make([]string, 0)
	for k := range section.KeyValues {
		r = append(r, k)
	}
	return r
}

func (section *Section) Get(key string, defValue string) string {
	if v, ok := section.KeyValues[key]; ok {
		return v
	}
	return defValue
}

func (section *Section) GetInt(key string, defValue int) int {
	if v, ok := section.KeyValues[key]; ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return defValue
}

func (section *Section) GetInt64(key string, defValue int64) int64 {
	if v, ok := section.KeyValues[key]; ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return i
		}
	}
	return defValue
}

func (section *Section) GetUInt64(key string, defValue uint64) uint64 {
	if v, ok := section.KeyValues[key]; ok {
		i, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return i
		}
	}
	return defValue
}

func (section *Section) GetFloat32(key string, defValue float32) float32 {
	if v, ok := section.KeyValues[key]; ok {
		f, err := strconv.ParseFloat(v, 32)
		if err == nil {
			return float32(f)
		}
	}
	return defValue
}

func (section *Section) GetFloat64(key string, defValue float64) float64 {
	if v, ok := section.KeyValues[key]; ok {
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f
		}
	}
	return defValue
}
