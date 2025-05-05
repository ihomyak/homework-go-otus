package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}
	for _, stage := range stages {
		if stage != nil {
			in = stage(doneStage(in, done))
		}
	}
	return in
}

func doneStage(in In, done In) Out {
	out := make(Bi)

	go func() {
		defer func() {
			close(out)
			//nolint:all
			for range in {
				// waiting when all channels will be close
			}
		}()

		for {
			select {
			case <-done:
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				out <- val
			}
		}
	}()

	return out
}
