package api

import (
	"fmt"
	"net/http"
	"policy-server/pkg/db/model"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type CreatePolicyPOST struct {
	Name               string `json:"name" validate:"required"`
	State              string `json:"state" validate:"required"`
	ProviderName       string `json:"provider_name" validate:"required"`
	ResourceType       string `json:"resource_type" validate:"required"`
	ResourcesEvaluated string `json:"resources_evaluated" validate:"required"`
	Tags               string `json:"tags" validate:"required"`
	Steampipe          string `json:"steampipe" validate:"required"`
}

func (s *Server) CreatePolicy(c echo.Context) error {
	var request CreatePolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	policy, err := s.handler.FindPolicyByName(request.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Find policy: %v", err))
	}
	if policy != nil {
		return c.JSON(http.StatusFound, "Policy already exists")
	}
	policy = &model.Policy{
		ID:                 uuid.Must(uuid.NewV4()).String(),
		Name:               request.Name,
		State:              request.State,
		ProviderName:       request.ProviderName,
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
	Name string `json:"name" validate:"required"`
}

func (s *Server) DeletePolicy(c echo.Context) error {
	var request DeletePolicyPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	policy, err := s.handler.FindPolicyByName(request.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Find policy: %v", err))
	}
	if policy == nil {
		return c.JSON(http.StatusNotFound, "Policy not found")
	}
	if err := s.handler.DeletePolicy(request.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Delete policy %s: %v", request.Name, err))
	}

	return nil
}

func (s *Server) UpdatePolicy(c echo.Context) error {
	return nil
}

func (s *Server) QueryPolicy(c echo.Context) error {
	return nil
}
