package policy

import "time"

type Policy struct {
	imageRule ImageRule
	regexRule RegexRule
	olderThanGivenDateRule OlderThanGivenDateRule
	nRule NRule
}


type ImageRule struct {
	enable bool
	inverse bool
	images []string
}

type RegexRule struct {
	enable bool
	inverse bool
	pattern []string
}

type OlderThanGivenDateRule struct {
	enable bool
	inverse bool
	date time.Duration
}

type NRule struct {
	enable bool
	inverse bool
	keep int
}