package file

import (
	"gopkg.in/yaml.v3"
)

type YAMLMarshalVisitor struct {
	marshalVisitor
}

func (v *YAMLMarshalVisitor) BuildData() ([]byte, error) {
	return yaml.Marshal(v.fileData)
}

func NewYAMLExporter() *Exporter {
	return NewExporter(&YAMLMarshalVisitor{})
}

type YAMLImporter struct {
	Importer
}

func (i *YAMLImporter) unmarshalFileData(data []byte) (FileData, error) {
	var fileData FileData
	err := yaml.Unmarshal(data, &fileData)
	return fileData, err
}

func NewYAMLImporter() *ImporterImpl {
	return &ImporterImpl{
		importer: &YAMLImporter{},
	}
}
