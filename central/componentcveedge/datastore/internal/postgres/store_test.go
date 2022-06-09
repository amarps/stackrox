// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type ImageComponentCveEdgesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestImageComponentCveEdgesStore(t *testing.T) {
	suite.Run(t, new(ImageComponentCveEdgesStoreSuite))
}

func (s *ImageComponentCveEdgesStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.Require().NoError(err)

	Destroy(ctx, pool)

	s.pool = pool
	gormDB := pgtest.OpenGormDB(s.T(), source)
	defer pgtest.CloseGormDB(s.T(), gormDB)
	s.store = CreateTableAndNewStore(ctx, pool, gormDB)
}

func (s *ImageComponentCveEdgesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.pool.Exec(ctx, "TRUNCATE image_component_cve_edges CASCADE")
	s.T().Log("image_component_cve_edges", tag)
	s.NoError(err)
}

func (s *ImageComponentCveEdgesStoreSuite) TearDownSuite() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *ImageComponentCveEdgesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	componentCVEEdge := &storage.ComponentCVEEdge{}
	s.NoError(testutils.FullInit(componentCVEEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComponentCVEEdge, exists, err := store.Get(ctx, componentCVEEdge.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComponentCVEEdge)

}
