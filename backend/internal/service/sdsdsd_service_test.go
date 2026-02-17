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

type SdsdsdServiceTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockRepo  *mock_repository.MockSdsdsdRepository
	mockCache *mock_pkg_cache.MockCacheService
	service   SdsdsdService
}

func (s *SdsdsdServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockSdsdsdRepository(s.ctrl)
	s.mockCache = mock_pkg_cache.NewMockCacheService(s.ctrl)
	s.service = NewSdsdsdService(s.mockRepo, s.mockCache)
}

func (s *SdsdsdServiceTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *SdsdsdServiceTestSuite) TestCreate_Success() {
	req := dto.CreateSdsdsdRequest{
		Name: "test",
	}

	entity := &entity.Sdsdsd{
		Name: req.Name,
	}

	s.mockRepo.EXPECT().Create(entity).Return(nil)
	s.mockCache.EXPECT().DeletePattern("sdsdsdsdd:*").Return(nil)

	res, err := s.service.Create(req)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), res)
}

func TestSdsdsdServiceTestSuite(t *testing.T) {
	suite.Run(t, new(SdsdsdServiceTestSuite))
}
