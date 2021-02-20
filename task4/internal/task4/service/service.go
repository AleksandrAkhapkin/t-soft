package service

import "github.com/AleksandrAkhapkin/testTNS/task4/internal/client/postgres"

type Service struct {
	p *postgres.Postgres
}

func NewService(pq *postgres.Postgres) (*Service, error) {

	return &Service{
		p: pq,
	}, nil
}
