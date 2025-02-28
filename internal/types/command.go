package types

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Args    []Arg  `yaml:"args"`
}
