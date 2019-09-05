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
func (r *respositoryProduct) FindAll(pagination map[string]interface{}) (interface{}, error) {
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
		columns, err := rows.Columns()
		if err != nil {
			ch <- false
			return
		}
		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for i, _ := range values {
			var val interface{}
			values[i] = val
		}

		for rows.Next() {
			for i := range columns {
				valuePtrs[i] = &values[i]
			}

			rows.Scan(valuePtrs...)
			var arr = make(map[string]interface{})
			for i, col := range columns {
				val := values[i]
				b, ok := val.([]byte)
				var v interface{}
				if ok {
					v = string(b)
				} else {
					v = val
				}
				arr[col] = v
			}
			fmt.Println(arr)

			prod = append(prod, arr)
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
