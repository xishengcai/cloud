package images

//
//import (
//	"encoding/json"
//	"os"
//
//	"k8s.io/klog/v2"
//)
//
//var (
//	pullQueue chan *Pull
//)
//
//func init() {
//	b, err := os.ReadFile("cache.json")
//	if err != nil {
//		klog.Warningf("read cache.json failed, %s", err.Error())
//	} else {
//		err = json.Unmarshal(b, &cache)
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	start()
//}
//
//func start() {
//	pullQueue = make(chan *Pull, 100)
//	go func() {
//		for {
//			select {
//			case p, ok := <-pullQueue:
//				if !ok {
//					panic("pull channel get  pop failed")
//				}
//				imageInfo, err := p.download()
//				if err != nil {
//					klog.Errorf("imageInfo: %v, err: %v", imageInfo, err)
//					cache.setReason(imageInfo, err.Error())
//				}
//
//			}
//		}
//	}()
//}
//
///* 文件操作 */
//// 有个坑，Python、Java的写文件默认函数操作默认是覆盖的，而是Golang的OpenFile函数写入默认是追加的
//// os.O_TRUNC 覆盖写入，不加则追加写入
//func saveToLocal() {
//	f, err := os.OpenFile("cache.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
//
//	if err != nil {
//		panic(err)
//	}
//	b, _ := json.Marshal(cache)
//	defer f.Close()
//	n, _ := f.Seek(0, os.SEEK_END)
//	_, err = f.WriteAt(b, n)
//	if err != nil {
//		panic(err)
//	}
//}
