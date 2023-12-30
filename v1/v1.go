package v1

type V1 interface {
	GetManifest() (m Manifest, err error)
	Process(value interface{})
}
