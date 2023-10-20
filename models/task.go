package models

type TaskRunner interface {
	Run(input interface{}) (interface{}, error)
}
type Task struct {
	ID     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Status status `json:"status"`
	Steps  []Step `json:"steps"`
}

type status string

const (
	statusRunning status = "running"
)

type Step struct {
	Name     string    `json:"name" bson:"name"`
	Skip     bool      `json:"skip" bson:"skip"`
	Status   status    `json:"status"`
	SubTasks []SubTask `json:"subTasks" bson:"subTasks"`
}

type SubTask struct {
	Name        string      `json:"name" bson:"name"`
	Input       interface{} `json:"input" bson:"input"`
	OutPut      interface{} `json:"outPut" bson:"outPut"`
	Skip        bool        `json:"skip" bson:"skip"`
	ErrContinue bool        `json:"errContinue" bson:"errContinue"`
	// 任务创建者实现定制功能
	TaskRunner TaskRunner
}

func (s SubTask) Run() (interface{}, error) {
	out, err := s.TaskRunner.Run(s.Input)
	return out, err
}

func (t Task) Start() error {
	for _, step := range t.Steps {
		for _, task := range step.SubTasks {
			_, err := task.Run()
			if err != nil && !task.ErrContinue {
				return err
			}
		}
	}
	return nil
}
