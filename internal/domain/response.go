package domain

type AnswerType string

const (
	AnswerTypeOnlyAnswer AnswerType = "only-answer"
	AnswerTypeFindTours  AnswerType = "find-tours"
)

type FindTourCriteria struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type FindToursCriteriesResponse struct {
	AnswerType  AnswerType         `json:"answerType"`
	ShortAnswer string             `json:"shortAnswer"`
	Positional  []FindTourCriteria `json:"positional"`
	KeyWords    []string           `json:"keyWords"`
}
