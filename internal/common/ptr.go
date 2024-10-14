package common

// XXX: copied from osbuild/images:internal/common/pointers.go
func ToPtr[T any](x T) *T {
	return &x
}
