package conversion

import "reflect"

type typePair struct {
	source reflect.Type
	dest   reflect.Type
}
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}
type scope struct {
	converter *Converter
	meta      *Meta
}
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredUntypedConversions map[typePair]struct{}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}
type ConversionFunc func(a, b interface{}, scope Scope) error
