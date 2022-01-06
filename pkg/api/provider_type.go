package api

import (
	"fmt"
	"net/http"
	"policy-service/pkg/db/model"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type ProviderTypePOST struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	State       string `json:"state" validate:"required"`
	Description string `json:"description"`
}

func (s *Server) CreateProviderType(c echo.Context) error {
	var request ProviderTypePOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	providerType := &model.ProviderType{
		ID:          uuid.Must(uuid.NewV4()).String(),
		Name:        request.Name,
		State:       request.State,
		Description: request.Description,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := s.handler.CreateProviderType(providerType); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create provider type: %v", err))
	}

	return nil
}

type QueryProviderTypePOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) QueryProviderType(c echo.Context) error {
	var request QueryProviderTypePOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	providerType, err := s.handler.QueryProviderType(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider type: %v", err))
	}
	if providerType == nil {
		return c.JSON(http.StatusNotFound, "Provider type not found")
	}

	return c.JSON(http.StatusOK, providerType)
}

func (s *Server) QueryProviderTypes(c echo.Context) error {
	providerTypes, err := s.handler.QueryProviderTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider types: %v", err))
	}
	if providerTypes == nil {
		return c.JSON(http.StatusNotFound, "Provider types not found")
	}

	return c.JSON(http.StatusOK, providerTypes)
}

func (s *Server) UpdateProviderType(c echo.Context) error {
	var request ProviderTypePOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	providerType, err := s.handler.QueryProviderType(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider type: %v", err))
	}
	if providerType == nil {
		return c.JSON(http.StatusNotFound, "Provider type not found")
	}

	providerType.State = request.State
	providerType.Name = request.Name
	providerType.Description = request.Description
	providerType.UpdateTime = time.Now()
	if err := s.handler.UpdateProviderType(providerType); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Update provider type: %v", err))
	}

	return nil
}
