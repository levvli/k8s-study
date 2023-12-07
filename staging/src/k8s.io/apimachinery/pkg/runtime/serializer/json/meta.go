package json

import "k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/schema"

// MetaFactory is used to store and retrieve the version and kind
// information for JSON objects in a serializer.
type MetaFactory interface {
	// Interpret should return the version and kind of the wire-format of
	// the object.
	Interpret(data []byte) (*schema.GroupVersionKind, error)
}

var DefaultMetaFactory = SimpleMetaFactory{}

// SimpleMetaFactory provides default methods for retrieving the type and version of objects
// that are identified with an "apiVersion" and "kind" fields in their JSON
// serialization. It may be parameterized with the names of the fields in memory, or an
// optional list of base structs to search for those fields in memory.
type SimpleMetaFactory struct {
}
