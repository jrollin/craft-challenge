package domain

type Story struct {
	ID             string
	Title          string
	Description    string
	Specifications Specifications
}

type Stories []*Story

func (s Stories) GetFirstStory() (*Story, error) {
	return s[0], nil
}

type Specifications []*Specification

type Specification struct {
	Description string
	Rules       Rules
}

type Rules []*Rule

type Assertion struct {
	RequestPart string
	Matcher     string
	Param       string
	Expected    string
}

type Request struct {
	Method       string
	URL          string
	QueryParams  map[string]string
	BodyParams   map[string]string
	HeaderParams map[string]string
}

type Rule struct {
	Description string
	Query       Request
	Assertion   Assertion
}
