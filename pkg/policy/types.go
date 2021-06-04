package policy

type Policy struct {
	ImageRule struct {
		Enable bool     `yaml:"enable"`
		Keep   bool     `yaml:"keep"`
		Images []string `yaml:"images"`
	} `yaml:"ImageRule"`

	NRule struct {
		Enable           bool `yaml:"enable"`
		Keep             bool `yaml:"keep"`
		Size             int  `yaml:"size"`
		IncludeLatestTag bool `yaml:"includeLatestTag"`
	} `yaml:"NRule"`

	RegexRule struct {
		Enable  bool     `yaml:"enable"`
		Keep    bool     `yaml:"keep"`
		Pattern []string `yaml:"pattern"`
	} `yaml:"RegexRule"`

	OlderThanGivenDateRule struct {
		Enable bool   `yaml:"enable"`
		Keep   bool   `yaml:"keep"`
		Date   string `yaml:"date"`
	} `yaml:"OlderThanGivenDateRule"`
}
