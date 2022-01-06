package api

import (
	"fmt"
	"net/http"
	"policy-server/pkg/db/model"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type TagsPOST struct {
	Type     string `json:"type" validate:"required"`
	Key      string `json:"key" validate:"required"`
	Value    string `json:"value" validate:"required"`
	State    string `json:"state" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

func (s *Server) CreateTags(c echo.Context) error {
	var request TagsPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	tags, err := s.handler.QueryTagsByKeys(request.Type, request.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tags: %v", err))
	}
	if tags != nil {
		return c.JSON(http.StatusFound, "Tags already exists")
	}
	tags = &model.Tags{
		ID:         uuid.Must(uuid.NewV4()).String(),
		Type:       request.Type,
		Key:        request.Key,
		Value:      request.Value,
		State:      request.State,
		Provider:   request.Provider,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := s.handler.CreateTags(tags); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create tags: %v", err))
	}

	return nil
}

type QueryTagsPOST struct {
	Type string `json:"type" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

func (s *Server) QueryTags(c echo.Context) error {
	var request QueryTagsPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	tags, err := s.handler.QueryTagsByKeys(request.Type, request.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tags: %v", err))
	}
	if tags == nil {
		return c.JSON(http.StatusNotFound, "Tags not found")
	}

	return c.JSON(http.StatusOK, tags)
}
