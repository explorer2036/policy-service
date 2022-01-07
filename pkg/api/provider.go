package api

import (
	"fmt"
	"net/http"
	"policy-service/pkg/db/model"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type ProviderPOST struct {
	ID           string `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Url          string `json:"url" validate:"required"`
	ProviderType string `json:"provider_type" validate:"required"`
	State        string `json:"state" validate:"required"`
	Description  string `json:"description"`
}

func (s *Server) validateProviderRequest(c echo.Context, request *ProviderPOST) error {
	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	// validate the provider type
	providerType, err := s.handler.QueryProviderType(request.ProviderType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider type: %v", err))
	}
	if providerType == nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid provider type %s", request.ProviderType))
	}

	return nil
}

func (s *Server) CreateProvider(c echo.Context) error {
	var request ProviderPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validateProviderRequest(c, &request); err != nil {
		return err
	}

	provider := &model.Provider{
		ID:           uuid.Must(uuid.NewV4()).String(),
		Name:         request.Name,
		Url:          request.Url,
		ProviderType: request.ProviderType,
		State:        request.State,
		Description:  request.Description,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	if err := s.handler.CreateProvider(provider); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create provider: %v", err))
	}

	return nil
}

type QueryProviderPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) QueryProvider(c echo.Context) error {
	var request QueryProviderPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	provider, err := s.handler.QueryProvider(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider: %v", err))
	}
	if provider == nil {
		return c.JSON(http.StatusNotFound, "Provider not found")
	}

	return c.JSON(http.StatusOK, provider)
}

func (s *Server) QueryProviders(c echo.Context) error {
	providers, err := s.handler.QueryProviders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query providers: %v", err))
	}
	if providers == nil {
		return c.JSON(http.StatusNotFound, "Providers not found")
	}

	return c.JSON(http.StatusOK, providers)
}

func (s *Server) UpdateProvider(c echo.Context) error {
	var request ProviderPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validateProviderRequest(c, &request); err != nil {
		return err
	}

	provider, err := s.handler.QueryProvider(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query provider: %v", err))
	}
	if provider == nil {
		return c.JSON(http.StatusNotFound, "Provider not found")
	}

	provider.State = request.State
	provider.Name = request.Name
	provider.Url = request.Url
	provider.ProviderType = request.ProviderType
	provider.Description = request.Description
	provider.UpdateTime = time.Now()
	if err := s.handler.UpdateProvider(provider); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Update provider: %v", err))
	}

	return nil
}
