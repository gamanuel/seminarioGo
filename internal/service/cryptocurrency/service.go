package cryptocurrency

import (
	"github.com/jmoiron/sqlx"
	"github.com/seminarioGo/internal/config"
	"fmt"
)

//Cryptocurrency structure
type Cryptocurrency struct {
	ID int64
	Type string
	Quantity int64
}

type Service interface {
	AddCryptocurrency(Cryptocurrency) (Cryptocurrency,error)
	FindAll() ([]*Cryptocurrency, error)
	updateCryptocurrency(Cryptocurrency) (bool,error)
	RemoveByID(int) (bool,error)
	FindByID(int) (*Cryptocurrency, error)
} 

type service struct {
	db *sqlx.DB
	conf *config.Config
}

func New(db *sqlx.DB,c *config.Config) (Service, error){
	return service{db,c}, nil
}

func (s service) AddCryptocurrency(c Cryptocurrency) (Cryptocurrency,error) {
	
	sqlStatement := "INSERT INTO cryptocurrency (type, quantity) VALUES (?,?)"
	
	res, err := s.db.Exec(sqlStatement, c.Type, c.Quantity)
	if err != nil {
		return c,err
	}

	c.ID,_ = res.LastInsertId()
	fmt.Println(res.LastInsertId())

	return c,nil
}

func (s service) FindByID(ID int) (*Cryptocurrency, error) {
	var cryptocurrency Cryptocurrency

	sqlStatement := "SELECT * FROM cryptocurrency WHERE ID=?"
	if err := s.db.Get(&cryptocurrency,sqlStatement, ID); 
	err != nil {
		return nil, err
	}
	
	return &cryptocurrency, nil

}

func (s service) FindAll() ([]*Cryptocurrency, error) {

	var list []*Cryptocurrency

	sqlStatement := "SELECT * FROM cryptocurrency"
	if err := s.db.Select(&list,sqlStatement); 
	err != nil {
		
		return nil, err
	}

	return list, nil

}


func (s service) RemoveByID(ID int) (bool,error) {

	sqlStatement := "DELETE FROM cryptocurrency WHERE ID = ?"
	_,err := s.db.Exec(sqlStatement, ID) 
	if err != nil {
		return false, err
	}

	return true, nil

}

func (s service) updateCryptocurrency(c Cryptocurrency) (bool,error) {
	
	sqlStatement := "UPDATE cryptocurrency SET Type = ?, Quantity = ? WHERE ID = ?"
	
	_, err := s.db.Exec(sqlStatement, c.Type, c.Quantity,c.ID)
	if err != nil {
		return false,err
	}

	return true,nil
}
