package server

import (
	"context"
	"testing"

	"github.com/kosalnik/keeper/internal/service"
	"github.com/kosalnik/keeper/pkg/gophkeeper"
	"github.com/stretchr/testify/require"
)

func TestKeeperServer_Ping(t *testing.T) {
	s := NewKeeperServer(&service.UserService{})
	got, err := s.Ping(context.Background(), &gophkeeper.Empty{})
	require.NoError(t, err)
	require.Equal(t, &gophkeeper.Empty{}, got)
}
