package schema

type GroupKind struct {
	Group string
	Kind  string
}
type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

func (gvk GroupVersionKind) Empty() bool {
	return len(gvk.Group) == 0 && len(gvk.Version) == 0 && len(gvk.Kind) == 0
}
func (gvk GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}
func (gvk GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvk.Group, Version: gvk.Version}
}
func (gvk GroupVersionKind) String() string {
	return gvk.Group + "/" + gvk.Version + ", Kind=" + gvk.Kind
}

type GroupVersion struct {
	Group   string
	Version string
}
