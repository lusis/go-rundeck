package rundeck_test

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck"
	"github.com/stretchr/testify/suite"
)

type TokenIntegrationTestSuite struct {
	suite.Suite
	CreatedTokens []rundeck.Token
	TestClient    *rundeck.Client
}

func (s *TokenIntegrationTestSuite) TearDownSuite() {
	for _, token := range s.CreatedTokens {
		e := s.TestClient.DeleteToken(token.Token)
		if e != nil {
			s.T().Logf("unable to clean up test token %s (err: %s)", token.Token, e.Error())
		}
	}
}

func (s *TokenIntegrationTestSuite) SetupSuite() {
	client := testNewTokenAuthClient()
	s.TestClient = client
	s.CreatedTokens = []rundeck.Token{}
}

func (s *TokenIntegrationTestSuite) TestCreateToken() {
	createToken, createErr := s.TestClient.CreateToken("admin")
	s.Require().NoError(createErr)
	s.Require().NotNil(createToken)
	s.CreatedTokens = append(s.CreatedTokens, *createToken)
}

func (s *TokenIntegrationTestSuite) TestGetToken() {
	createToken, err := s.TestClient.CreateToken("admin")
	s.Require().NoError(err)

	s.CreatedTokens = append(s.CreatedTokens, *createToken)
	getToken, err := s.TestClient.GetToken(createToken.Token)
	s.Require().NoError(err)
	s.Require().Equal(createToken.ID, getToken.ID)
}

func (s *TokenIntegrationTestSuite) TestListTokens() {
	list, err := s.TestClient.ListTokens()
	s.Require().NoError(err)
	s.Require().Len(list, len(s.CreatedTokens))
}

func (s *TokenIntegrationTestSuite) TestListTokensForUser() {
	allTokens, err := s.TestClient.ListTokensForUser("admin")
	s.Require().NoError(err)
	s.Require().Len(allTokens, len(s.CreatedTokens))
}

func TestIntegrationTokenSuite(t *testing.T) {
	if testing.Short() || testRundeckRunning() == false {
		t.Skip("skipping integration testing")
	}
	
	suite.Run(t, &TokenIntegrationTestSuite{})
}
