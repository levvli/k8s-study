package runtime

import (
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/conversion"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/schema"
	"reflect"
)

type Scheme struct {
	gvkToType                 map[schema.GroupVersionKind]reflect.Type
	typeToGVK                 map[reflect.Type][]schema.GroupVersionKind
	unversionedTypes          map[reflect.Type]schema.GroupVersionKind
	unversionedKinds          map[string]reflect.Type
	fieldLabelConversionFuncs map[schema.GroupVersionKind]FieldLabelConversionFunc
	defaulterFuncs            map[reflect.Type]func(interface{})
	converter                 *conversion.Converter
	versionPriority           map[string][]string
	observedVersions          []schema.GroupVersion
	schemeName                string
}
type FieldLabelConversionFunc func(label, value string) (internalLabel, internalValue string, err error)
