package e

import (
	"fmt"
	"runtime"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"
)

func RecoverGoPanic() {
	if err := recover(); err != nil {
		printStack()
		klog.Errorf("panic recover from err: %v", err)
	}
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	klog.V(4).Infof("==> %s", string(buf[:n]))
}

func ISAlreadyError(err error) error {
	if errors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func IsExist(err error) bool {
	return err != nil
}

func IsNoExist(err error) bool {
	return err == nil
}

func AssertError(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}

func MergeError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	var msg strings.Builder
	for index, item := range errs {
		msg.Write([]byte(fmt.Sprintf("%d :, %v", index, item)))
	}
	return fmt.Errorf(msg.String())
}
