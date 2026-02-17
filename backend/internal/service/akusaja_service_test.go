package service

import (
	"testing"

	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	mock_repository "github.com/hadi-projects/go-react-starter/internal/mock/repository"
	mock_pkg_cache "github.com/hadi-projects/go-react-starter/internal/mock/pkg/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type AkusajaServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockAkusajaRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   AkusajaService
}

func (s *AkusajaServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockAkusajaRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewAkusajaService(s.mockRepo, s.mockCache)
}

func (s *AkusajaServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *AkusajaServiceTestSuite) TestCreate_Success() {
	req := dto.CreateAkusajaRequest{
		Name: "test",
	}

	entity := &entity.Akusaja{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("akusaja:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestAkusajaServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AkusajaServiceTestSuite))
}
