package postgres_test

import (
	"context"

	"github.com/google/uuid"
	"github.com/kosalnik/keeper/internal/entity"
	"github.com/kosalnik/keeper/internal/postgres"
)

func (s *PostgresSuite) TestVaultRepository_CreateCredentials() {
	r := postgres.NewVaultRepository(s.conn)
	m := entity.NewCredentials("lilu creds", "lilu", "tratata")
	s.Require().NoError(r.CreateCredentials(context.Background(), TestUserID, m))
	s.Require().NotEqual(uuid.Nil, m.ID)
}
