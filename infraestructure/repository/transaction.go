package repository

import (
	"fmt"
	"github.com/Augustojvs/imersao-fullstack-fullcycle/domain/model"
	"github.com/jinzhu/gorm"
)


type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err:= t.Db.Create(transaction).Error
	if err != nil{
		return err
	}
	return nil
}

func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err:= t.Db.Save(transaction).Error
	if err != nil{
		return err
	}
	return nil
}

func (r TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}