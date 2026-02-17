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

type MakanServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockMakanRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   MakanService
}

func (s *MakanServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockMakanRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewMakanService(s.mockRepo, s.mockCache)
}

func (s *MakanServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *MakanServiceTestSuite) TestCreate_Success() {
	req := dto.CreateMakanRequest{
		Name: "test",
	}

	entity := &entity.Makan{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("makan:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestMakanServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MakanServiceTestSuite))
}
