package storage

import (
	pbp "root/genprotos/public"
	pbp_v "root/genprotos/public_vote"
	pbv "root/genprotos/vote"
	pbe "root/genprotos/election"
	pbpt "root/genprotos/party"
	pbc "root/genprotos/candidate"
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
