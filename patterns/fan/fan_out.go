package fan

import "context"

func FanOut[T any](ctx context.Context, input chan T, outputs ...chan T) {
	for {
		select {
		case data := <-input:
			for _, output := range outputs {
				output <- data
			}
		case <-ctx.Done():
			break
		default:
			continue
		}
	}
}
