package models

import (
	"fmt"
	"testing"

	"github.com/xishengcai/cloud/pkg/common"
)

type printInfo struct {
}

func (p printInfo) Run(input interface{}) (interface{}, error) {
	fmt.Println(input)
	return nil, nil
}

func TestTask(t *testing.T) {
	taskCase := Task{
		ID:   common.GetUUID(),
		Name: "demo",
		Steps: []Step{
			{
				Name: "pre-check",
				Skip: false,
				SubTasks: []SubTask{
					{
						Name:       "check-host version",
						Input:      "check host",
						TaskRunner: printInfo{},
					},
				},
			},
			{
				Name: "install-docker",
				Skip: false,
				SubTasks: []SubTask{
					{
						Name:       "check docker version",
						Input:      "aa",
						TaskRunner: printInfo{},
					},
					{
						Name:       "install docker",
						Input:      "xx",
						TaskRunner: printInfo{},
					},
				},
			},
		},
	}

	err := taskCase.Start()
	if err != nil {
		t.Fatalf("task start failed, %v", err)
	}
}
