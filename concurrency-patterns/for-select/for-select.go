for {
	select {
	case <-done:
		return
	default:
		// Do non-preemptable work
	}
}
