package http

type FindTourCriteria struct {
	Name  string
	Value string
}

type FindToursCriteriesResponse struct {
	AnswerType  string
	ShortAnswer string
	Positional  []FindTourCriteria
	KeyWords    []string
}
