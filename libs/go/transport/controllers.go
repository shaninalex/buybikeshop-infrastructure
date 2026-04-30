package transport

type Controllable interface {
	// GetNamespace returns a name of an object that controller operates
	GetNamespace() string
}
