package adalight

// Effect is the interface for generic lighting effects.
//
// Init performs Effect initialization, passing the underlying Strip
// object to the Effect, as well as the approximate number of frames
// per second the Effect will be run at. The Effect should return its
// name for logging/debugging purposes
//
// Frame commands the Effect to update the Strip before it is written
// out to the hardware device. A frame count is provided.
type Effect interface {
	Init(strip *Strip, fps int) (name string)
	Frame(n int) (done bool)
}
