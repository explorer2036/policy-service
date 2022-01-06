package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"policy-server/pkg/mock"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite

	server *Server
}

func (ts *testSuite) SetupSuite() {
	ts.server = &Server{
		validate: validator.New(),
		handler:  mock.NewStore(),
	}
}

func (ts *testSuite) TestCreatePolicy() {
	createPolicyTestCases := []struct {
		Name               string `json:"name"`
		State              string `json:"state"`
		ProviderName       string `json:"provider_name"`
		ResourceType       string `json:"resource_type"`
		ResourcesEvaluated string `json:"resources_evaluated"`
		Tags               string `json:"tags"`
		Steampipe          string `json:"steampipe"`
		Error              string `json:"-"`
	}{
		{
			Name:               "policy-test-01",
			State:              "active",
			ProviderName:       "provider-name-test-01",
			ResourceType:       "resource-type-test-01",
			ResourcesEvaluated: "resources-evaluated-test-01",
			Tags:               "tags-test-01",
			Steampipe:          "steampipe-test-01",
			Error:              "",
		},
	}

	for i, tc := range createPolicyTestCases {
		tc := tc
		ts.T().Run(fmt.Sprintf("CreatePolicyTestCase-%d", i), func(t *testing.T) {
			e := echo.New()
			data, _ := json.Marshal(tc)
			r := httptest.NewRequest(http.MethodPost, "/policy", bytes.NewBuffer(data))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			ec := e.NewContext(r, w)
			if err := ts.server.CreatePolicy(ec); err != nil {
				ts.T().Fatalf("create policy: %v", err)
			}

			d, err := ioutil.ReadAll(w.Body)
			if err != nil {
				ts.T().Fatalf("read all: %v", err)
			}
			ts.Equal(tc.Error, string(d))
		})
	}
}
