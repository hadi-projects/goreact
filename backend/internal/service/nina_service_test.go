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

type NinaServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockNinaRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   NinaService
}

func (s *NinaServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockNinaRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewNinaService(s.mockRepo, s.mockCache)
}

func (s *NinaServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *NinaServiceTestSuite) TestCreate_Success() {
	req := dto.CreateNinaRequest{
		Names: "test",
	}

	entity := &entity.Nina{
		Names: req.Names,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("nina:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestNinaServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NinaServiceTestSuite))
}
