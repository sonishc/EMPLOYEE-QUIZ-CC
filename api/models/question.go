package models

type Question struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Answer string `json:"answer"`
}

var Questions []Question

func SeedQuestions() {
	Questions = append(Questions, Question{ID: "1", Text: "5+6", Answer: "11"})
	Questions = append(Questions, Question{ID: "2", Text: "2+7", Answer: "9"})
	Questions = append(Questions, Question{ID: "3", Text: "1+6", Answer: "7"})
	Questions = append(Questions, Question{ID: "4", Text: "7+9", Answer: "16"})
	Questions = append(Questions, Question{ID: "5", Text: "3+5", Answer: "8"})
}
