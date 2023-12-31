package core

import (
	"gopkg.in/yaml.v3"
)

type ManifestType string

type Manifest struct {
	Package     string
	Name        string
	Description string
	Type        ManifestType
	Publisher   string
}

const (
	StringManifestType  = "string"
	NumberManifestType  = "number"
	BooleanManifestType = "boolean"
)

func (receiver *Instance) LoadManifest(in []byte) (m Manifest, err error) {
	err = yaml.Unmarshal(in, &m)

	return
}

func (receiver *Instance) GetManifest() Manifest {
	return receiver.manifest
}
