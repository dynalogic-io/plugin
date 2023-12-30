package plugin

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
