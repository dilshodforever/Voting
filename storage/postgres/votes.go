package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	pr "root/genprotos/product"
	

)

type ProductStorage struct {
	Db *sql.DB
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{Db: db}
}    


func (bc *ProductStorage) CreateProduct(pr *pr.ProductCreateReq) (*pr.Void,error){
	id:=uuid.NewString()
	_, err:=bc.Db.Exec(`insert into products(id, name, price, category_id,expired_at) 
						values($1, $2, $3, $4, $5)`,
						id, pr.Name, pr.Price, pr.CategoryId, pr.ExpiredAt)
	return nil,err
}



func (bc *ProductStorage) GetAllProduct(p *pr.Filter) (*pr.ProductGetAllResp,error){
	product:=pr.ProductGetAllResp{}
	row, err:=bc.Db.Query(`select name, price, expered_at from products`)
	if err!=nil{
		return nil, err
	}
	for row.Next(){
		all:=pr.ProductGetAllResp{}
		row.Scan(&all.Products[0].Name,
				 &all.Products[0].Price,
				 &all.Products[0].ExpiredAt)
				 product.Products=append(product.Products, all.Products[0])
	}

	return &product,err
}


