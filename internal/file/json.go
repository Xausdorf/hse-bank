package file

import (
	"encoding/json"
)

type JSONMarshalVisitor struct {
	marshalVisitor
}

func (v *JSONMarshalVisitor) BuildData() ([]byte, error) {
	return json.MarshalIndent(v.fileData, "", "  ")
}

func NewJSONExporter() *Exporter {
	return NewExporter(&JSONMarshalVisitor{})
}

type JSONImporter struct {
	Importer
}

func (i *JSONImporter) unmarshalFileData(data []byte) (FileData, error) {
	var fileData FileData
	err := json.Unmarshal(data, &fileData)
	return fileData, err
}

func NewJSONImporter() *ImporterImpl {
	return &ImporterImpl{
		importer: &JSONImporter{},
	}
}
