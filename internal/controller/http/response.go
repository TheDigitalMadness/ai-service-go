package http

type AnswerType string

const (
	AnswerTypeOnlyAnswer AnswerType = "only-answer"
	AnswerTypeFindTours  AnswerType = "find-tours"
)

type FindTourCriteria struct {
	Name  string
	Value string
}

type FindToursCriteriesResponse struct {
	AnswerType  AnswerType
	ShortAnswer string
	Positional  []FindTourCriteria
	KeyWords    []string
}
