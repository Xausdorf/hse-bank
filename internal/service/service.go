package service

import (
	"log/slog"

	"github.com/Xausdorf/hse-bank/internal/command"
)

type AccountCommands struct {
	Create  command.CreateBankAccountHandler
	Update  command.UpdateBankAccountHandler
	Delete  command.DeleteBankAccountHandler
	GetByID command.GetBankAccountByIDHandler
	GetAll  command.GetAllBankAccountsHandler
}

type CategoryCommands struct {
	Create  command.CreateCategoryHandler
	Update  command.UpdateCategoryHandler
	Delete  command.DeleteCategoryHandler
	GetByID command.GetCategoryByIDHandler
	GetAll  command.GetAllCategoriesHandler
}

type OperationCommands struct {
	Create        command.CreateOperationHandler
	Update        command.UpdateOperationHandler
	Delete        command.DeleteOperationHandler
	GetByID       command.GetOperationByIDHandler
	GetWithFilter command.GetOperationsWithFilterHandler
}

type FileCommands struct {
	ExportAll command.ExportAllHandler
	Import    command.ImportHandler
}

type Service struct {
	Account   AccountCommands
	Category  CategoryCommands
	Operation OperationCommands
	File      FileCommands
}

func NewService(
	accountFacade command.BankAccountFacade,
	categoryFacade command.CategoryFacade,
	operationFacade command.OperationFacade,
	logger *slog.Logger,
) *Service {
	acct := AccountCommands{
		Create:  command.NewCreateBankAccountHandler(accountFacade, logger),
		Update:  command.NewUpdateBankAccountHandler(accountFacade, logger),
		Delete:  command.NewDeleteBankAccountHandler(accountFacade, logger),
		GetByID: command.NewGetBankAccountByIDHandler(accountFacade, logger),
		GetAll:  command.NewGetAllBankAccountsHandler(accountFacade, logger),
	}

	cat := CategoryCommands{
		Create:  command.NewCreateCategoryHandler(categoryFacade, logger),
		Update:  command.NewUpdateCategoryHandler(categoryFacade, logger),
		Delete:  command.NewDeleteCategoryHandler(categoryFacade, logger),
		GetByID: command.NewGetCategoryByIDHandler(categoryFacade, logger),
		GetAll:  command.NewGetAllCategoriesHandler(categoryFacade, logger),
	}

	op := OperationCommands{
		Create:        command.NewCreateOperationHandler(operationFacade, logger),
		Update:        command.NewUpdateOperationHandler(operationFacade, logger),
		Delete:        command.NewDeleteOperationHandler(operationFacade, logger),
		GetByID:       command.NewGetOperationByIDHandler(operationFacade, logger),
		GetWithFilter: command.NewGetOperationsWithFilterHandler(operationFacade, logger),
	}

	fileCmds := FileCommands{
		ExportAll: command.NewExportAllHandler(accountFacade, categoryFacade, operationFacade, logger),
		Import:    command.NewImportHandler(accountFacade, categoryFacade, operationFacade, logger),
	}

	return &Service{Account: acct, Category: cat, Operation: op, File: fileCmds}
}
