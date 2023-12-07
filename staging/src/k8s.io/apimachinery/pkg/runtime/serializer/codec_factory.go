package serializer

import (
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/runtime/serializer/json"
)

type CodecFactoryOptions struct {
	// Strict configures all serializers in strict mode
	Strict bool
	// Pretty includes a pretty serializer along with the non-pretty one
	Pretty bool
}
type CodecFactoryOptionsMutator func(*CodecFactoryOptions)
type CodecFactory struct {
	scheme    *runtime.Scheme
	universal runtime.Decoder
	accepts   []runtime.SerializerInfo

	legacySerializer runtime.Serializer
}

func NewCodecFactory(scheme *runtime.Scheme, mutators ...CodecFactoryOptionsMutator) CodecFactory {
	options := CodecFactoryOptions{Pretty: true}
	for _, fn := range mutators {
		fn(&options)
	}
	serializers := newSerializersForScheme(scheme, json.DefaultMetaFactory, options)
	return newCodecFactory(scheme, serializers)
}
func newSerializersForScheme(scheme *runtime.Scheme, mf json.MetaFactory, options CodecFactoryOptions) []serializerType {

}
