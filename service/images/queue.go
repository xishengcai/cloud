package images

import "k8s.io/klog/v2"

var (
	pullQueue chan *Pull
)

func init() {
	start()
}

func start() {
	pullQueue = make(chan *Pull, 100)
	go func() {
		for {
			select {
			case p, ok := <-pullQueue:
				if !ok {
					panic("pull channel get  pop failed")
				}
				imageInfo, err := p.download()
				if err != nil {
					klog.Errorf("imageInfo: %v, err: %v", imageInfo, err)
					cache.setReason(imageInfo, err.Error())
				}

			}
		}
	}()

}
