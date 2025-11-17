package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-test01/gormlearn/globa"
	"gorm.io/gorm"
)

func main() {

	db := globa.DB
	initData(db)

	db.Transaction(func(tx *gorm.DB) error {

		fromAccount := &Account{}
		tx.First(fromAccount, 1)

		toAccount := &Account{}
		tx.First(toAccount, 2)

		if fromAccount.Balance > 100 {
			fromAccount.Balance -= 100

			err1 := tx.Save(fromAccount).Error
			if err1 != nil {
				return err1
			}

			toAccount.Balance += 100
			err2 := tx.Save(toAccount).Error
			if err2 != nil {
				return err2
			}

			err3 := tx.Save(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100}).Error
			if err3 != nil {
				return err3
			}
		} else {
			err3 := tx.Save(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 0}).Error
			if err3 != nil {
				return err3
			}
		}

		return nil
	})
}

func initData(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.Debug().Create(&Account{1, 101})
	db.Debug().Create(&Account{2, 11})

	db.AutoMigrate(&Transaction{})
}

type Account struct {
	ID      int `gorm:"primarykey"`
	Balance int
}

type Transaction struct {
	ID            int `gorm:"primarykey;autoIncrement"`
	FromAccountId int
	ToAccountId   int
	Amount        int
}
