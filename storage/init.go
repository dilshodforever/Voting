package storage

import (
	basket "root/genprotos/basket"
	pr "root/genprotos/product"
	user "root/genprotos/user"
)

type InitRoor interface {
	Product() Product
	User() User
	Basket() Basket
}
type Product interface {
	CreateProduct(pr *pr.ProductCreateReq) (*pr.Void, error)
	GetAllProduct(p *pr.Filter) (*pr.ProductGetAllResp,error)
}

type User interface {
	CreateUser(user *user.CreateUser) (*user.UserREsponse, error)
	GetAllUser(fl *user.Filter) (*user.GetAllUsers,error)
}

type Election interface {
	CreateElection(basket *basket.CreateBasket) (*basket.Void, error)
}
