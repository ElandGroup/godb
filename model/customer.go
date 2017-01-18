package model

import "time"
import "errors"

type Customer struct {
	Id        int64
	Mobile    string    `xorm:"unique"`
	CreatedAt time.Time `xorm:"created 'in_date_time'"`
	UpdatedAt time.Time `xorm:"updated 'modi_date_time'"`
}

func (Customer) TableName() string {
	return "user"
}

func (c *Customer) Create() error {
	if affected, err := db.InsertOne(c); err != nil {
		return err
	} else if affected == 0 {
		return errors.New("Affected rows: 0")
	}

	return nil
}

func (Customer) Retrive() (*Customer, error) {
	var c Customer
	_, err := db.Get(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (Customer) Get(id int64) (*Customer, error) {
	var c Customer
	_, err := db.Id(id).Get(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
