package domain

type Result struct {
	Id       string         `json:"id"`
	UserId   string         `json:"userId"`
	Cloud    string         `json:"cloud"`
	MockExam string         `json:"mockExam"`
	Result   string         `json:"result"`
	Wrong    string         `json:"wrong"`
	Details  []WrongAnswers `json:"details"`
}

type WrongAnswers struct {
	ServiceType string              `json:"serviceType"`
	Topics      []TopicWrongAnswers `json:"topics"`
	Total       int                 `json:"total"`
}

type TopicWrongAnswers struct {
	Topic         string `json:"topics"`
	WrongsAnswers int    `json:"wrongAnswers"`
}
