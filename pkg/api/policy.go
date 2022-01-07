package api

import (
	"fmt"
	"net/http"
	"policy-service/pkg/db/model"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type PolicyPOST struct {
	ID                 string   `json:"id" validate:"id"`
	Name               string   `json:"name" validate:"required"`
	State              string   `json:"state" validate:"required"`
	Provider           string   `json:"provider" validate:"required"`
	ResourceType       string   `json:"resource_type" validate:"required"`
	ResourcesEvaluated string   `json:"resources_evaluated" validate:"required"`
	Tags               []string `json:"tags" validate:"required"`
	Steampipe          string   `json:"steampipe" validate:"required"`
}

func (s *Server) validatePolicyRequest(c echo.Context, request *PolicyPOST) error {
	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	// validate the provider
	provider, err := s.handler.QueryProvider(request.Provider)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider: %v", err))
	}
	if provider == nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid provider %s", request.Provider))
	}

	// validate the tags
	for _, tag := range request.Tags {
		tag, err := s.handler.QueryTag(tag)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tag: %v", err))
		}
		if tag == nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid tag %s", tag))
		}
	}
	return nil
}

func (s *Server) CreatePolicy(c echo.Context) error {
	var request PolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validatePolicyRequest(c, &request); err != nil {
		return nil
	}

	policy := &model.Policy{
		ID:                 uuid.Must(uuid.NewV4()).String(),
		Name:               request.Name,
		State:              request.State,
		Provider:           request.Provider,
		ResourceType:       request.ResourceType,
		ResourcesEvaluated: request.ResourcesEvaluated,
		Tags:               request.Tags,
		Steampipe:          request.Steampipe,
		CreateTime:         time.Now(),
		UpdateTime:         time.Now(),
	}
	if err := s.handler.CreatePolicy(policy); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create policy: %v", err))
	}

	return nil
}

type DeletePolicyPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) DeletePolicy(c echo.Context) error {
	var request DeletePolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	policy, err := s.handler.QueryPolicy(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query policy: %v", err))
	}
	if policy == nil {
		return c.JSON(http.StatusNotFound, "Policy not found")
	}
	if err := s.handler.DeletePolicy(request.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Delete policy %s: %v", request.ID, err))
	}

	return nil
}

func (s *Server) UpdatePolicy(c echo.Context) error {
	var request PolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validatePolicyRequest(c, &request); err != nil {
		return nil
	}

	policy, err := s.handler.QueryPolicy(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query policy: %v", err))
	}
	if policy == nil {
		return c.JSON(http.StatusNotFound, "Policy not found")
	}

	policy.State = request.State
	policy.Provider = request.Provider
	policy.ResourceType = request.ResourceType
	policy.ResourcesEvaluated = request.ResourcesEvaluated
	policy.Tags = request.Tags
	policy.Steampipe = request.Steampipe
	policy.UpdateTime = time.Now()
	if err := s.handler.UpdatePolicy(policy); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Update policy: %v", err))
	}

	return nil
}

type QueryPolicyPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) QueryPolicy(c echo.Context) error {
	var request QueryPolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	policy, err := s.handler.QueryPolicy(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query policy: %v", err))
	}
	if policy == nil {
		return c.JSON(http.StatusNotFound, "Policy not found")
	}

	return c.JSON(http.StatusOK, policy)
}

func (s *Server) QueryPolicies(c echo.Context) error {
	policies, err := s.handler.QueryPolicies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query policies: %v", err))
	}
	if policies == nil {
		return c.JSON(http.StatusNotFound, "Policies not found")
	}

	return c.JSON(http.StatusOK, policies)
}
