package policy

type Policy struct {
	NRule struct {
		Enable           bool `yaml:"enable"`
		Keep             bool `yaml:"keep"`
		Size             int  `yaml:"size"`
		IncludeLatestTag bool `yaml:"includeLatestTag"`
	} `yaml:"n"`

	RegexRule struct {
		Enable  bool     `yaml:"enable"`
		Keep    bool     `yaml:"keep"`
		Pattern []string `yaml:"pattern"`
	} `yaml:"regex"`

	UntilDateRule struct {
		Enable bool   `yaml:"enable"`
		Keep   bool   `yaml:"keep"`
		Date   string `yaml:"date"`
	} `yaml:"until-date"`
}
