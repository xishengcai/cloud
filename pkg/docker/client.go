package docker

import "github.com/docker/docker/client"

var Client *client.Client

func NewDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	return cli
}

func init() {
	Client = NewDockerClient()
}
