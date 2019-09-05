package repo

import (
	"fmt"
	"golang-boilerplates/api/helper"
	"golang-boilerplates/api/models"

	"github.com/jinzhu/gorm"
)

type respositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProducts(db *gorm.DB) *respositoryProduct {
	return &respositoryProduct{db}
}

func (r *respositoryProduct) Save(msg string) (string, error) {
	return msg, nil
}
func (r *respositoryProduct) FindAll(pagination map[string]string) (interface{}, error) {
	var err error
	// products := []*models.Product{}
	fmt.Println(pagination)
	done := make(chan bool)
	var prod []interface{}

	go func(ch chan<- bool) {
		defer close(ch)
		rows, err := r.db.Debug().Model(&models.Product{}).Select(pagination["field"]).Limit(pagination["limit"]).Rows()
		if err != nil {
			ch <- false
			return
		}
		cols, _ := rows.Columns()

		for rows.Next() {
			data := make(map[string]string)
			columns := make([]string, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}
			rows.Scan(columnPointers...)
			for i, colName := range cols {
				data[colName] = columns[i]
			}
			fmt.Println(data)
			prod = append(prod, data)
		}
		ch <- true
	}(done)
	if helper.OK(done) {
		return prod, nil
	}
	return nil, err
}
func (r *respositoryProduct) FindById(msg string) (string, error) {
	return msg, nil
}
func (r *respositoryProduct) Update(msg string) (string, error) {
	return msg, nil
}
func (r *respositoryProduct) Delete(msg string) (string, error) {
	return msg, nil
}
