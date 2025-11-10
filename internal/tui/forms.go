package tui

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Xausdorf/hse-bank/internal/command"
	"github.com/Xausdorf/hse-bank/internal/domain"
	"github.com/Xausdorf/hse-bank/internal/file"
)

func (m *Model) prepareFileForm(action string) {
	switch action {
	case "Export JSON":
		m.prompts = []string{"Enter file path to export JSON to:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			path := strings.TrimSpace(ans[0])
			_, err := m.svc.File.ExportAll.Handle(context.Background(), command.ExportAll{
				FilePath: path,
				Exporter: file.NewJSONExporter(),
			})
			if err != nil {
				return "", err
			}
			return "Exported JSON to: " + path, nil
		}
	case "Export YAML":
		m.prompts = []string{"Enter file path to export YAML to:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			path := strings.TrimSpace(ans[0])
			_, err := m.svc.File.ExportAll.Handle(context.Background(), command.ExportAll{
				FilePath: path,
				Exporter: file.NewYAMLExporter(),
			})
			if err != nil {
				return "", err
			}
			return "Exported YAML to: " + path, nil
		}
	case "Import JSON":
		m.prompts = []string{"Enter file path to import JSON from:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			path := strings.TrimSpace(ans[0])
			_, err := m.svc.File.Import.Handle(context.Background(), command.Import{
				FilePath: path,
				Importer: file.NewJSONImporter(),
			})
			if err != nil {
				return "", err
			}
			return "Imported JSON from: " + path, nil
		}
	case "Import YAML":
		m.prompts = []string{"Enter file path to import YAML from:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			path := strings.TrimSpace(ans[0])
			_, err := m.svc.File.Import.Handle(context.Background(), command.Import{
				FilePath: path,
				Importer: file.NewYAMLImporter(),
			})
			if err != nil {
				return "", err
			}
			return "Imported YAML from: " + path, nil
		}
	}
}

func (m *Model) prepareFormFor(action string) {
	m.prompts = nil
	m.answers = nil
	m.curField = 0
	m.onSubmit = nil

	switch m.entity {
	case "BankAccount":
		m.prepareBankAccountForm(action)
	case "Category":
		m.prepareCategoryForm(action)
	case "Operation":
		m.prepareOperationForm(action)
	case "File":
		m.prepareFileForm(action)
	}

	if len(m.prompts) > 0 {
		m.ti.Placeholder = m.prompts[0]
		m.ti.SetValue("")
	}
}

func (m *Model) prepareBankAccountForm(action string) {
	switch action {
	case "View by ID":
		m.prompts = []string{"Enter account ID:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			res, err := m.svc.Account.GetByID.Handle(context.Background(), command.GetBankAccountByID{ID: id})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("ID: %s\nName: %s\nBalance: %d", res.ID(), res.Name(), res.Balance()), nil
		}
	case "Create":
		m.prompts = []string{"Enter name:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			name := strings.TrimSpace(ans[0])
			res, err := m.svc.Account.Create.Handle(context.Background(), command.CreateBankAccount{Name: name})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Created account with ID: %s", res.ID()), nil
		}
	case "Delete by ID":
		m.prompts = []string{"Enter account ID to delete:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			_, err := m.svc.Account.Delete.Handle(context.Background(), command.DeleteBankAccount{ID: id})
			if err != nil {
				return "", err
			}
			return "Deleted", nil
		}
	case "Update":
		m.prompts = []string{"Enter ID:", "Enter name:", "Enter balance (int):"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			name := strings.TrimSpace(ans[1])
			balStr := strings.TrimSpace(ans[2])
			bal, err := strconv.ParseInt(balStr, 10, 64)
			if err != nil {
				return "", err
			}
			_, err = m.svc.Account.Update.Handle(context.Background(), command.UpdateBankAccount{ID: id, Name: name, Balance: bal})
			if err != nil {
				return "", err
			}
			return "Updated", nil
		}
	case "View all":
		m.prompts = []string{"(press Enter to view all)"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			res, err := m.svc.Account.GetAll.Handle(context.Background(), command.GetAllBankAccounts{})
			if err != nil {
				return "", err
			}
			b := &strings.Builder{}
			for _, a := range res {
				fmt.Fprintf(b, "- ID: %s | Name: %s | Balance: %d\n", a.ID(), a.Name(), a.Balance())
			}
			return b.String(), nil
		}
	}
}

func (m *Model) prepareCategoryForm(action string) {
	switch action {
	case "View by ID":
		m.prompts = []string{"Enter category ID:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			res, err := m.svc.Category.GetByID.Handle(context.Background(), command.GetCategoryByID{ID: id})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("ID: %s\nName: %s\nType: %s", res.ID(), res.Name(), res.OperationType().String()), nil
		}
	case "Create":
		m.prompts = []string{"Enter name:", "Enter type (income|expense):"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			name := strings.TrimSpace(ans[0])
			typ := strings.ToLower(strings.TrimSpace(ans[1]))
			var opType domain.OperationType
			switch strings.ToLower(typ) {
			case "income":
				opType = domain.Income
			case "expense":
				opType = domain.Expense
			default:
				return "", errors.New("unkown operation type")
			}
			res, err := m.svc.Category.Create.Handle(context.Background(), command.CreateCategory{Name: name, OpType: opType})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Created category with ID: %s", res.ID()), nil
		}
	case "Delete by ID":
		m.prompts = []string{"Enter category ID to delete:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			_, err := m.svc.Category.Delete.Handle(context.Background(), command.DeleteCategory{ID: id})
			if err != nil {
				return "", err
			}
			return "Deleted", nil
		}
	case "Update":
		m.prompts = []string{"Enter ID:", "Enter name:", "Enter type (income|expense):"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			name := strings.TrimSpace(ans[1])
			typ := strings.ToLower(strings.TrimSpace(ans[2]))
			var opType domain.OperationType
			switch strings.ToLower(typ) {
			case "income":
				opType = domain.Income
			case "expense":
				opType = domain.Expense
			default:
				return "", errors.New("unkown operation type")
			}
			_, err := m.svc.Category.Update.Handle(context.Background(), command.UpdateCategory{ID: id, Name: name, OpType: opType})
			if err != nil {
				return "", err
			}
			return "Updated", nil
		}
	case "View all":
		m.prompts = []string{"(press Enter to view all)"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			res, err := m.svc.Category.GetAll.Handle(context.Background(), command.GetAllCategories{})
			if err != nil {
				return "", err
			}
			b := &strings.Builder{}
			for _, c := range res {
				fmt.Fprintf(b, "- ID: %s | Name: %s | Type: %s\n", c.ID(), c.Name(), c.OperationType().String())
			}
			return b.String(), nil
		}
	}
}

func (m *Model) prepareOperationForm(action string) {
	switch action {
	case "View by ID":
		m.prompts = []string{"Enter operation ID:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			res, err := m.svc.Operation.GetByID.Handle(context.Background(), command.GetOperationByID{ID: id})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("ID: %s\nAccountID: %s\nCategoryID: %s\nAmount: %d\nDate: %s\nDesc: %s", res.ID(), res.AccountID(), res.CategoryID(), res.Amount(), res.Date().Format(time.RFC3339), res.Description()), nil
		}
	case "Create":
		m.prompts = []string{"Enter account ID:", "Enter category ID:", "Enter amount (int):", "Enter date (YYYY-MM-DD):", "Enter description:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			accountID := strings.TrimSpace(ans[0])
			categoryID := strings.TrimSpace(ans[1])
			amtStr := strings.TrimSpace(ans[2])
			amt, err := strconv.ParseInt(amtStr, 10, 64)
			if err != nil {
				return "", err
			}
			dateStr := strings.TrimSpace(ans[3])
			parsed, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return "", err
			}
			desc := strings.TrimSpace(ans[4])
			res, err := m.svc.Operation.Create.Handle(context.Background(), command.CreateOperation{AccountID: accountID, CategoryID: categoryID, Amount: amt, Date: parsed, Description: desc})
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("Created operation ID: %s", res.ID()), nil
		}
	case "Delete by ID":
		m.prompts = []string{"Enter operation ID to delete:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			_, err := m.svc.Operation.Delete.Handle(context.Background(), command.DeleteOperation{ID: id})
			if err != nil {
				return "", err
			}
			return "Deleted", nil
		}
	case "Update":
		m.prompts = []string{"Enter ID:", "Enter account ID:", "Enter category ID:", "Enter amount (int):", "Enter date (YYYY-MM-DD):", "Enter description:"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			id := strings.TrimSpace(ans[0])
			accountID := strings.TrimSpace(ans[1])
			categoryID := strings.TrimSpace(ans[2])
			amt, err := strconv.ParseInt(strings.TrimSpace(ans[3]), 10, 64)
			if err != nil {
				return "", err
			}
			parsed, err := time.Parse("2006-01-02", strings.TrimSpace(ans[4]))
			if err != nil {
				return "", err
			}
			desc := strings.TrimSpace(ans[5])
			_, err = m.svc.Operation.Update.Handle(context.Background(), command.UpdateOperation{ID: id, AccountID: accountID, CategoryID: categoryID, Amount: amt, Date: parsed, Description: desc})
			if err != nil {
				return "", err
			}
			return "Updated", nil
		}
	case "View all":
		m.prompts = []string{"(press Enter to view all)"}
		m.answers = make([]string, len(m.prompts))
		m.onSubmit = func(ans []string) (string, error) {
			res, err := m.svc.Operation.GetWithFilter.Handle(context.Background(), command.GetOperationsWithFilter{Predicate: func(_ *domain.Operation) bool { return true }})
			if err != nil {
				return "", err
			}
			b := &strings.Builder{}
			for _, o := range res {
				fmt.Fprintf(b, "- ID: %s | Acc: %s | Cat: %s | Amt: %d | Date: %s | Desc: %s\n", o.ID(), o.AccountID(), o.CategoryID(), o.Amount(), o.Date().Format(time.RFC3339), o.Description())
			}
			return b.String(), nil
		}
	}
}
