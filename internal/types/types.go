package types

type Project struct {
	Name     string    `yaml:"name"`
	Commands []Command `yaml:"commands"`
	Folder   string    `yaml:"folder,omitempty"`
	Key      string    `yaml:"-"` // This will store the map key
}

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Args    []Arg  `yaml:"args"`
}

type Arg struct {
	Name    string `yaml:"name"`
	Options string `yaml:"options"`
	Value   string
}
