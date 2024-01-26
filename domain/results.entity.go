package domain

type Result struct {
	Id       string         `json:"id"`
	UserId   string         `json:"userId"`
	Cloud    string         `json:"cloud"`
	MockExam string         `json:"mockExam"`
	Result   float64        `json:"result"`
	Wrong    int            `json:"wrong"`
	Correct  int            `json:"correct"`
	Details  []WrongAnswers `json:"details"`
}

type WrongAnswers struct {
	Id            string `json:"id"`
	ResultId      string `json:"ResultId"`
	ServiceType   string `json:"serviceType"`
	Topic         string `json:"topic"`
	WrongsAnswers int64  `json:"wrongAnswers"`
	Total         int64  `json:"total"`
}

type ExamTypeBody struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

type ExamReq struct {
	ExamName string `json:"examName"`
}
