package models

import "k8s.io/klog/v2"

type TaskRunner interface {
	Run() error
}

type Watcher interface {
	SaveTaskStatus(st status) error
}

type Task struct {
	ID      string `json:"id" bson:"id" gorm:"primaryKey" `
	Name    string `json:"name" bson:"name" gorm:"name"`
	Status  status `json:"status" bson:"status" gorm:"status"`
	Steps   []Step `json:"steps" bson:"steps" gorm:"steps"`
	Watcher `json:"-"`
}

type status string
type Log string

const (
	statusRunning status = "running"
	statusSuccess status = "success"
	statusFailed  status = "failed"
)

type Step struct {
	Name        string      `json:"name" bson:"name" gorm:"name"`
	Skip        bool        `json:"skip" bson:"skip" gorm:"skip"`
	Status      status      `json:"status" bson:"status" gorm:"status"`
	Input       interface{} `json:"input" bson:"input" gorm:"input"`
	IgnoreError bool        `json:"ignoreError" bson:"ignoreError"  gorm:"ignoreError"`
	// 任务创建者实现定制功能
	TaskRunner TaskRunner
}

func (s Step) Run() error {
	if s.Skip {
		klog.Infof("step %s skip ", s.Name)
		return nil
	}
	err := s.TaskRunner.Run()
	return err
}

func (s Step) setStatus(err error) {
	if err != nil {
		s.Status = statusFailed
	} else {
		s.Status = statusSuccess
	}
}

func (t Task) Start() error {
	for _, step := range t.Steps {
		err := step.Run()
		if err != nil && !step.IgnoreError {
			return err
		}
	}
	return nil
}

func (t Task) Notify() error {
	return t.Watcher.SaveTaskStatus(t.Status)
}
