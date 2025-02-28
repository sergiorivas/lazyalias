package types

type Project struct {
	Name     string    `yaml:"name"`
	Commands []Command `yaml:"commands"`
	Folder   string    `yaml:"folder,omitempty"`
	Key      string    `yaml:"-"` // This will store the map key
}
