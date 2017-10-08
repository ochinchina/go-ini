package ini

import (
	"bytes"
	"fmt"
	"io"
)

type Ini struct {
	sections map[string]*Section
}

func NewIni() *Ini {
	return &Ini{sections: make(map[string]*Section)}
}

func (ini *Ini) NewSection(name string) *Section {
	if section, ok := ini.sections[name]; ok {
		return section
	}
	section := NewSection(name)
	ini.sections[name] = section
	return section
}

func (ini *Ini) AddSection(section *Section) {
	ini.sections[section.Name] = section
}

// Get all the section name in the ini
//
// return all the section names
func (ini *Ini) Sections() []*Section {
	r := make([]*Section, 0)
	for _, section := range ini.sections {
		r = append(r, section)
	}
	return r
}

// check if a key exists or not in the Ini
//
// return true if the key in section exists
func (ini *Ini) HasKey(sectionName, key string) bool {
	if section, ok := ini.sections[sectionName]; ok {
		return section.HasKey(key)
	}
	return false
}

// get section by section name
//
// return: section or nil
func (ini *Ini) GetSection(name string) *Section {
	if section, ok := ini.sections[name]; ok {
		return section
	}
	return nil
}

func (ini *Ini) HasSection(name string) bool {
	return ini.GetSection(name) != nil
}

func (ini *Ini) Get(sectionName, key string) (string, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetValue(key)
	}
	return "", fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetWithDefault(sectionName, key string, defValue string) string {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetValueWithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetInt(sectionName, key string) (int, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetInt(key)
	}
	return 0, fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetIntWithDefault(sectionName, key string, defValue int) int {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetIntWithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetInt64(sectionName, key string) (int64, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetInt64(key)
	}
	return 0, fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetInt64WithDefault(sectionName, key string, defValue int64) int64 {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetInt64WithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetUint64(sectionName, key string) (uint64, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetUint64(key)
	}
	return 0, fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetUint64WithDefault(sectionName, key string, defValue uint64) uint64 {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetUint64WithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetFloat32(sectionName, key string) (float32, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetFloat32(key)
	}
	return 0, fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetFloat32WithDefault(sectionName, key string, defValue float32) float32 {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetFloat32WithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetFloat64(sectionName, key string) (float64, error) {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetFloat64(key)
	}
	return 0, fmt.Errorf("no such section:%s", sectionName)
}

func (ini *Ini) GetFloat64WithDefault(sectionName, key string, defValue float64) float64 {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetFloat64WithDefault(key, defValue)
	}
	return defValue
}

func (ini *Ini) String() string {
	buf := bytes.NewBuffer(make([]byte, 0))
	ini.Write(buf)
	return buf.String()
}

func (ini *Ini) Write(writer io.Writer) error {
	for _, section := range ini.sections {
		err := section.Write(writer)
		if err != nil {
			return err
		}
	}
	return nil
}
