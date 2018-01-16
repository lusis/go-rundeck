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
	s.NoError(createErr)
	s.NotNil(createToken)
	s.CreatedTokens = append(s.CreatedTokens, *createToken)
}

func (s *TokenIntegrationTestSuite) TestGetToken() {
	createToken, createErr := s.TestClient.CreateToken("admin")
	if createErr != nil {
		s.T().Fatalf("Can't create a token. cannot continue: %s", createErr.Error())
	}
	s.CreatedTokens = append(s.CreatedTokens, *createToken)
	getToken, getErr := s.TestClient.GetToken(createToken.Token)
	s.NoError(getErr)
	s.Equal(createToken.ID, getToken.ID)
}

func (s *TokenIntegrationTestSuite) TestListTokens() {
	list, listErr := s.TestClient.ListTokens()
	s.NoError(listErr)
	s.Len(list, len(s.CreatedTokens))
}

func (s *TokenIntegrationTestSuite) TestListTokensForUser() {
	allTokens, allErr := s.TestClient.ListTokensForUser("admin")
	s.NoError(allErr)
	s.Len(allTokens, len(s.CreatedTokens))
}

func TestIntegrationTokenSuite(t *testing.T) {
	if testRundeckRunning() {
		suite.Run(t, &TokenIntegrationTestSuite{})
	} else {
		t.Skip("rundeck isn't running for integration testing")
	}
}
