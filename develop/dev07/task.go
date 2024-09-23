package dev07

// or сигнализирует о закрытии одного из множества done-каналов закрытием возвращенного канала.
func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		done := make(chan struct{})

		for _, c := range channels {
			go func(c <-chan interface{}) {
				select {
				case <-c:
					close(done)
				}
			}(c)
		}

		// Ждем, пока один из каналов закроется.
		<-done
		close(out)
	}()

	return out
}
