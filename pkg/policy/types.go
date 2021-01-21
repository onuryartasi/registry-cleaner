package policy

type Policy struct {
	ImageRule struct {
		Enable  bool     `yaml:"enable"`
		Inverse bool     `yaml:"inverse"`
		Images []string `yaml:"images"`
	} `yaml:"ImageRule"`

	NRule struct {
		Enable  bool `yaml:"enable"`
		Inverse bool `yaml:"inverse"`
		Keep    int  `yaml:"keep"`
		IncludeLatestTag bool `yaml:"includeLatestTag"`
	} `yaml:"NRule"`

	RegexRule struct {
		Enable  bool     `yaml:"enable"`
		Inverse bool     `yaml:"inverse"`
		Pattern []string `yaml:"pattern"`
	} `yaml:"RegexRule"`

	OlderThanGivenDateRule struct {
		Enable  bool   `yaml:"enable"`
		Inverse bool   `yaml:"inverse"`
		Date    string `yaml:"date"`
	} `yaml:"OlderThanGivenDateRule"`

}
