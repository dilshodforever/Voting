package service

import (
	"context"
	"log"
	v "root/genprotos/election"
	pb "root/genprotos/party"
	"root/storage/postgres"
)

type PartyService struct {
	stg *postgres.PartyStorage
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgres.PartyStorage) *PartyService {
	return &PartyService{stg: stg}
}

func (c *PartyService) CreateParty(ctx context.Context, Party *pb.Party) (*v.Void, error) {
	v, err := c.stg.CreateParty(Party)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *PartyService) GetAllPartys(ctx context.Context, v *v.Void) (*pb.GetAllParty, error) {
	Partys, err := c.stg.GetAllParty(v)
	if err != nil {
		log.Print(err)
	}

	return Partys, err
}

func (c *PartyService) GetByIdParty(ctx context.Context, id *v.ById) (*pb.Party, error) {
	prod, err := c.stg.GetByIdParty(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PartyService) UpdateParty(ctx context.Context, Party *pb.Party) (*v.Void, error) {
	v, err := c.stg.UpdateParty(Party)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *PartyService) DeleteParty(ctx context.Context, id *v.ById) (*v.Void, error) {
	v, err := c.stg.DeleteParty(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
