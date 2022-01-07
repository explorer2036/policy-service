package api

import (
	"fmt"
	"net/http"
	"policy-service/pkg/db/model"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type TagPOST struct {
	Type     string `json:"type" validate:"required"`
	Key      string `json:"key" validate:"required"`
	Value    string `json:"value" validate:"required"`
	State    string `json:"state" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

func (s *Server) validateTagRequest(c echo.Context, request *TagPOST) error {
	if err := s.validate.Struct(request); err != nil {
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

	return c.JSON(http.StatusOK, "success")
}

func (s *Server) CreateTag(c echo.Context) error {
	var request TagPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validateTagRequest(c, &request); err != nil {
		return err
	}

	tag := &model.Tag{
		ID:         uuid.Must(uuid.NewV4()).String(),
		Type:       request.Type,
		Key:        request.Key,
		Value:      request.Value,
		State:      request.State,
		Provider:   request.Provider,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := s.handler.CreateTag(tag); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return c.JSON(http.StatusFound, "Tag already exists")
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create tag: %v", err))
	}
	response := map[string]string{
		"id": tag.ID,
	}

	return c.JSON(http.StatusOK, response)
}

type QueryTagPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) QueryTag(c echo.Context) error {
	var request QueryTagPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	tag, err := s.handler.QueryTag(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tag: %v", err))
	}
	if tag == nil {
		return c.JSON(http.StatusNotFound, "Tag not found")
	}

	return c.JSON(http.StatusOK, tag)
}

func (s *Server) QueryTags(c echo.Context) error {
	tags, err := s.handler.QueryTags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tags: %v", err))
	}
	if tags == nil {
		return c.JSON(http.StatusNotFound, "Tag not found")
	}

	return c.JSON(http.StatusOK, tags)
}
