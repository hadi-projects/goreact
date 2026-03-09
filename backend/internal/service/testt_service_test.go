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

type TesttServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockTesttRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   TesttService
}

func (s *TesttServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockTesttRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewTesttService(s.mockRepo, s.mockCache)
}

func (s *TesttServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *TesttServiceTestSuite) TestCreate_Success() {
	req := dto.CreateTesttRequest{
		Name: "test",
	}

	entity := &entity.Testt{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("testt:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestTesttServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TesttServiceTestSuite))
}
