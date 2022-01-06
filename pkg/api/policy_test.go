package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"policy-service/pkg/mock"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
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
		Name               string   `json:"name"`
		State              string   `json:"state"`
		Provider           string   `json:"provider"`
		ResourceType       string   `json:"resource_type"`
		ResourcesEvaluated string   `json:"resources_evaluated"`
		Tags               []string `json:"tags"`
		Steampipe          string   `json:"steampipe"`
		Error              string   `json:"-"`
	}{
		{
			Name:               "policy-test-01",
			State:              "active",
			Provider:           "provider-name-test-01",
			ResourceType:       "resource-type-test-01",
			ResourcesEvaluated: "resources-evaluated-test-01",
			Tags:               []string{uuid.Must(uuid.NewV4()).String()},
			Steampipe:          "steampipe-test-01",
			Error:              "",
		},
		{
			Name:               "policy-test-01",
			State:              "active",
			Provider:           "provider-name-test-01",
			ResourceType:       "resource-type-test-01",
			ResourcesEvaluated: "resources-evaluated-test-01",
			Tags:               []string{uuid.Must(uuid.NewV4()).String()},
			Steampipe:          "steampipe-test-01",
			Error:              "Policy already exists",
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

			var payload string
			if err := json.NewDecoder(w.Body).Decode(&payload); err != nil {
				if err != io.EOF {
					ts.T().Fatalf("json decode: %v", err)
				}
			}
			ts.Equal(tc.Error, payload)
		})
	}
}
