package machinery

type Environments struct {
	Order           int    `yaml:"order"`
	Name            string `yaml:"name"`
	Hash            string `yaml:"hash"`
	IsDefaultBranch bool   `yaml:"isDefaultBranch"`
}
