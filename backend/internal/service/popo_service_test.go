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

type PopoServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockPopoRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   PopoService
}

func (s *PopoServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockPopoRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewPopoService(s.mockRepo, s.mockCache)
}

func (s *PopoServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *PopoServiceTestSuite) TestCreate_Success() {
	req := dto.CreatePopoRequest{
		Name: "test",
	}

	entity := &entity.Popo{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("popo:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestPopoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PopoServiceTestSuite))
}
