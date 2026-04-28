package transport

type Controllable interface {
	// GetObject returns a name of an object that controller operates
	GetObject() string
}
