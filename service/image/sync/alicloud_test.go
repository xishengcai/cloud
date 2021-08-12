package sync

import (
	"os"
	"testing"

	"github.com/alibabacloud-go/cr-20181201/client"
)

func TestListInstance(t *testing.T) {
	a := NewAliImageRegistry("cr.cn-hangzhou.aliyuncs.com", os.Getenv("accessKeyId"), os.Getenv("accessKeySecret"))
	if err := a.setClient(); err != nil {
		t.Fatal(err)
	}
	resp, err := a.ListInstance(&client.ListInstanceRequest{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %v", resp)
}
