package service

import (
	"context"
	"log"
	v "root/genprotos/election"
	pb "root/genprotos/public"
	"root/storage/postgres"
)

type PublicService struct {
	stg *postgres.PublicStorage
	pb.UnimplementedPublicServiceServer
}

func NewPublicService(stg *postgres.PublicStorage) *PublicService {
	return &PublicService{stg: stg}
}

func (c *PublicService) CreatePublic(ctx context.Context, public *pb.Public) (*v.Void, error) {
	v, err := c.stg.CreatePublic(public)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *PublicService) GetAllPublics(ctx context.Context, v *v.Void) (*pb.GetAllPublic, error) {
	publics, err := c.stg.GetAllPublic(v)
	if err != nil {
		log.Print(err)
	}

	return publics, err
}

func (c *PublicService) GetByIdPublic(ctx context.Context, id *v.ById) (*pb.Public, error) {
	prod, err := c.stg.GetByIdPublic(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PublicService) UpdatePublic(ctx context.Context, public *pb.Public) (*v.Void, error) {
	v, err := c.stg.UpdatePublic(public)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *PublicService) DeletePublic(ctx context.Context, id *v.ById) (*v.Void, error) {
	v, err := c.stg.DeletePublic(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
