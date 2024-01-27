package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	res := in

	for _, s := range stages {
		res = s(checkDone(done, res))
	}

	return res
}

func checkDone(done In, in In) Out {
	out := make(Bi)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case i, ok := <-in:
				if !ok {
					return
				}
				out <- i
			}
		}
	}()

	return out
}
