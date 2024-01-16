package images

import (
	"github.com/docker/docker/client"

	"github.com/xishengcai/cloud/pkg/app"
)

type Repo struct {
	Addr     string  `json:"addr" default:"docker.io"`
	User     string  `json:"user"`
	Password string  `json:"password"`
	Org      string  `json:"org" default:"library"`
	Images   []Image `json:"images"`
}
type Image struct {
	Name    string `json:"name" default:"nginx"`
	Version string `json:"version" default:"latest"`
}

type Rule struct {
	All     bool
	Latest  bool
	Regular string
}

type Sync struct {
	Source Repo `json:"Source"`
	Target Repo `json:"Target"`
	Rule   Rule `json:"Rule"`
}

func (s Sync) Validate() error {
	return nil
}

func (s Sync) Run() app.ResultRaw {
	return nil
}

func (r Repo) getClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.WithVersionFromEnv())
}
