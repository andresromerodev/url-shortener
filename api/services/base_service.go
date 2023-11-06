package services

import (
	"context"

	"github.com/andresromeroh/url-shortener/api/db"
)

type BaseService struct {
	PrismaClient *db.PrismaClient
	ctx          context.Context
}

func NewBaseService() *BaseService {
	client := db.NewClient()
	ctx := context.Background()

	if err := client.Prisma.Connect(); err != nil {
		// Handle the error or panic
		panic(err)
	}

	return &BaseService{
		PrismaClient: client,
		ctx:          ctx,
	}
}

func (s *BaseService) Close() {
	if err := s.PrismaClient.Prisma.Disconnect(); err != nil {
		// Handle the error or log it
	}
}
