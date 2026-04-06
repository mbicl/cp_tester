package server

type ProblemContent struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Group       string `json:"group"`
	Interactive bool   `json:"interactive"`
	MemoryLimit int    `json:"memoryLimit"`
	TimeLimit   int    `json:"timeLimit"`
	TestType    string `json:"testType"`
	Tests       []Test `json:"tests"`
	Input       struct {
		Type string `json:"type"`
	} `json:"input"`
	Output struct {
		Type string `json:"type"`
	} `json:"output"`
	Languages Languages `json:"languages"`
	Batch     Batch     `json:"batch"`
}

type Test struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Languages struct {
	Java struct {
		MainClass string `json:"mainClass"`
		TaskClass string `json:"taskClass"`
	} `json:"java"`
}

type Batch struct {
	Id   string `json:"id"`
	Size int    `json:"size"`
}
