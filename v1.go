package plugin

type V1 interface {
	GetManifest() Manifest
	Process(value interface{})
}
