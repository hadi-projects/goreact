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

type MinaServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockMinaRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   MinaService
}

func (s *MinaServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockMinaRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewMinaService(s.mockRepo, s.mockCache)
}

func (s *MinaServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *MinaServiceTestSuite) TestCreate_Success() {
	req := dto.CreateMinaRequest{
		Name: "test",
	}

	entity := &entity.Mina{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("mina:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestMinaServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MinaServiceTestSuite))
}
