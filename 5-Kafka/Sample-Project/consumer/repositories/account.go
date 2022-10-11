package repositories

import (
	"time"

	"gorm.io/gorm"
)

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountTransaction struct {
	ID              string
	AccountID       string
	TransactionType string
	Amount          float64
	CreateAt        time.Time
}

type AccountRepository interface {
	SaveAccount(bankAccount BankAccount) error
	DeleteAccount(id string) error
	FindAccountByID(id string) (bankAccount BankAccount, err error)
	CreateTransaction(accountTransaction AccountTransaction) error
	FindTransactionsByID(id string) (accountTransaction []AccountTransaction, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&BankAccount{})
	db.AutoMigrate(&AccountTransaction{})
	return accountRepository{db}
}

func (obj accountRepository) SaveAccount(bankAccount BankAccount) error {
	return obj.db.Table("bank_accounts").Save(bankAccount).Error
}

func (obj accountRepository) DeleteAccount(id string) error {
	return obj.db.Table("bank_accounts").Where("id=?", id).Delete(&BankAccount{}).Error
}

func (obj accountRepository) FindAccountByID(id string) (bankAccount BankAccount, err error) {
	err = obj.db.Table("bank_accounts").Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}

func (obj accountRepository) CreateTransaction(accountTransaction AccountTransaction) error {
	return obj.db.Table("account_transactions").Create(accountTransaction).Error
}

func (obj accountRepository) FindTransactionsByID(id string) (accountTransaction []AccountTransaction, err error) {
	err = obj.db.Table("account_transactions").Where("account_id=?", id).Order("create_at desc").Find(&accountTransaction).Error
	return accountTransaction, err
}
