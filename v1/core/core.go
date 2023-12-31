package core

type Core interface {
	LoadManifest(manifest []byte) (m Manifest, err error)
}

type Instance struct {
	manifest Manifest
}

func New() Instance {
	return Instance{}
}
