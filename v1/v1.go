package v1

type V1 interface {
	GetManifest() Manifest
	Process(value interface{})
}
