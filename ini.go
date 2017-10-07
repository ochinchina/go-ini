package ini

type Ini struct {
	sections map[string]*Section
}

func NewIni() *Ini {
	return &Ini{sections: make(map[string]*Section)}
}

func (ini *Ini) NewSection(name string) *Section {
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
func (ini *Ini) Sections() []string {
	r := make([]string, 0)
	for k := range ini.sections {
		r = append(r, k)
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
func (ini *Ini) Get(sectionName, key string, defValue string) string {
	if section, ok := ini.sections[sectionName]; ok {
		return section.Get(key, defValue)
	}
	return defValue
}

func (ini *Ini) GetInt(sectionName, key string, defValue int) int {
	if section, ok := ini.sections[sectionName]; ok {
		return section.GetInt(key, defValue)
	}
	return defValue
}
