package core

// GoString represents object that can give their Go internals as a String
type GoString interface {
	// GoString returns a GO representation of this
	GoString() string
}
