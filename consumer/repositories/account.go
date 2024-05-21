package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindById(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.Table("bank").AutoMigrate(&BankAccount{}) // don't need to use .Table
	return accountRepository{db: db}
}

func (r accountRepository) Save(bankAccount BankAccount) error {
	return r.db.Table("bank").Save(bankAccount).Error
}

func (r accountRepository) Delete(id string) error {
	return r.db.Table("bank").Where("id = ?", id).Delete(&BankAccount{}).Error
}

func (r accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = r.db.Table("bank").Find(&bankAccounts).Error
	return bankAccounts, err
}

func (r accountRepository) FindById(id string) (bankAccount BankAccount, err error) {
	err = r.db.Table("bank").Where("id = ?", id).First(&bankAccount).Error
	return bankAccount, err
}
