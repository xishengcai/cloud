package kubernetes

var jobChan chan job

type job interface {
	startJob()
}

func init() {
	jobChan = make(chan job, 100)
	go func() {
		for {
			select {
			case i := <-jobChan:
				i.startJob()
			}
		}
	}()

}
