package domain

import "time"

type OperationType int

const (
	Income OperationType = iota
	Expense
)

func (t OperationType) String() string {
	switch t {
	case Income:
		return "Income"
	case Expense:
		return "Expense"
	default:
		return "UnkownOperationType"
	}
}

type BankAccount struct {
	id   string
	name string
	// Balance - balance in kopecks.
	balance int64
}

type Category struct {
	id            string
	name          string
	operationType OperationType
}

type Operation struct {
	id         string
	accountID  string
	categoryID string
	// Amount - amount in kopecks.
	amount      int64
	date        time.Time
	description string
}

func ReverseOperationType(opType OperationType) OperationType {
	if opType == Income {
		return Expense
	}
	return Income
}

func (b *BankAccount) ID() string {
	return b.id
}

func (b *BankAccount) Name() string {
	return b.name
}

func (b *BankAccount) Balance() int64 {
	return b.balance
}

func (b *BankAccount) ApplyOperation(amount int64, operationType OperationType) {
	if operationType == Expense {
		amount = -amount
	}
	b.balance += amount
}

func (c *Category) ID() string {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) OperationType() OperationType {
	return c.operationType
}

func (o *Operation) ID() string {
	return o.id
}

func (o *Operation) AccountID() string {
	return o.accountID
}

func (o *Operation) CategoryID() string {
	return o.categoryID
}

func (o *Operation) Amount() int64 {
	return o.amount
}

func (o *Operation) Date() time.Time {
	return o.date
}

func (o *Operation) Description() string {
	return o.description
}

func NewBankAccount(id, name string, balance int64) *BankAccount {
	return &BankAccount{
		id:      id,
		name:    name,
		balance: balance,
	}
}

func NewCategory(id, name string, operationType OperationType) *Category {
	return &Category{
		id:            id,
		name:          name,
		operationType: operationType,
	}
}
func NewOperation(id, accountID, categoryID string, amount int64, date time.Time, description string) *Operation {
	return &Operation{
		id:          id,
		accountID:   accountID,
		categoryID:  categoryID,
		amount:      amount,
		date:        date,
		description: description,
	}
}
