package policy

type Policy struct {
	NRule struct {
		Enable           bool `yaml:"enable"`
		Keep             bool `yaml:"keep"` // Not implemented yet.
		Size             int  `yaml:"size"`
		IncludeLatestTag bool `yaml:"includeLatestTag"`
	} `yaml:"n"`

	RegexRule struct {
		Enable  bool     `yaml:"enable"`
		Keep    bool     `yaml:"keep"` // Not implemented yet.
		Pattern []string `yaml:"pattern"`
	} `yaml:"regex"`

	UntilDateRule struct {
		Enable bool   `yaml:"enable"`
		Keep   bool   `yaml:"keep"` // Not implemented yet.
		Date   string `yaml:"date"`
	} `yaml:"until-date"`
}
