package images

import (
	"sync"
	"time"

	"github.com/xishengcai/cloud/pkg/app"
)

type Info struct {
}

func (l Info) Validate() error {
	return nil
}

func (l Info) Run() app.ResultRaw {
	return app.NewServiceResult(cache, nil)
}

var (
	cache = Cache{
		Entry: map[string][]ImageInfo{},
		Lock:  sync.Mutex{},
		Index: map[string]int{},
	}
)

type ImageInfo struct {
	FullName string
	Name     string
	Version  string
	Status   status
	Host     string
	Org      string
	Reason   string
	Updated  time.Time
}

type status string

const (
	downloading  status = "downloading"
	saving       status = "saving"
	pushingToOSS status = "pushingToOSS"
	waiting      status = "waiting"
	success      status = "success"
)

type Cache struct {
	Entry map[string][]ImageInfo
	Index map[string]int
	Lock  sync.Mutex
}

func (c Cache) set(name string, i ImageInfo) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	list, ok := c.Entry[name]
	if ok {
		c.Entry[name] = append(list, i)
		c.Index[i.FullName] = len(list) - 1
	} else {
		c.Entry[name] = []ImageInfo{i}
		c.Index[i.FullName] = 0
	}
}

func (c Cache) setStatus(i ImageInfo, stat status) {
	index := c.Index[i.FullName]
	c.Entry[i.Name][index].Status = stat
}

func (c Cache) setReason(i ImageInfo, reason string) {
	index := c.Index[i.FullName]
	c.Entry[i.Name][index].Reason = reason
}
