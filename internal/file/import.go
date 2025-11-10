package file

import (
	"fmt"
	"io"
	"os"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type Importer interface {
	unmarshalFileData(data []byte) (FileData, error)
	Import(filePath string) (
		[]domain.BankAccountDTO,
		[]domain.CategoryDTO,
		[]domain.OperationDTO,
		error,
	)
}

type ImporterImpl struct {
	importer Importer
}

func (i *ImporterImpl) Import(filePath string) (
	[]domain.BankAccountDTO,
	[]domain.CategoryDTO,
	[]domain.OperationDTO,
	error,
) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	fileData, err := i.importer.unmarshalFileData(bytes)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return fileData.Accounts, fileData.Categories, fileData.Operations, nil
}
