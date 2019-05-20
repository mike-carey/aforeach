package async

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Input generic.Type
type Output generic.Type

func ForEachInputToOutput(these []Input, do func(Input) (Output, error)) ([]Output, []error) {
	pool := make([]Output, 0)
	errs := make([]error, 0)

	if len(these) > 0 {
		var wg sync.WaitGroup

		wg.Add(len(these))

		poolCh := make(chan Output, len(these))
		errsCh := make(chan error, len(these))
		for _, this := range these {
			go func(this Input) {
				defer wg.Done()

				t, e := do(this)
				if e != nil {
					errsCh <- e
				} else {
					poolCh <- t
				}
			}(this)
		}

		wg.Wait()

		for _ = range these {
			select {
			case this := <-poolCh:
				pool = append(pool, this)
			case err := <-errsCh:
				errs = append(errs, err)
			}
		}
	}

	return pool, errs
}
