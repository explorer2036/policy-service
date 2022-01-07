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

type BenchmarkPOST struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name" validate:"required"`
	State              string   `json:"state" validate:"required"`
	Provider           string   `json:"provider" validate:"required"`
	ResourceType       string   `json:"resource_type" validate:"required"`
	ResourcesEvaluated string   `json:"resources_evaluated" validate:"required"`
	Tags               []string `json:"tags" validate:"required"`
	Policies           []string `json:"policies" validate:"required"`
	Description        string   `json:"description"`
}

func (s *Server) validateBenchmarkRequest(c echo.Context, request *BenchmarkPOST) error {
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

	// validate the tags
	for _, item := range request.Tags {
		tag, err := s.handler.QueryTag(item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query tag: %v", err))
		}
		if tag == nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid tag %s", item))
		}
	}

	// validate the policies
	for _, item := range request.Policies {
		policy, err := s.handler.QueryPolicy(item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query policy: %v", err))
		}
		if policy == nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid policy %s", item))
		}
	}

	return nil
}

func (s *Server) CreateBenchmark(c echo.Context) error {
	var request BenchmarkPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validateBenchmarkRequest(c, &request); err != nil {
		return err
	}

	benchmark := &model.Benchmark{
		ID:                 uuid.Must(uuid.NewV4()).String(),
		Name:               request.Name,
		State:              request.State,
		Provider:           request.Provider,
		ResourceType:       request.ResourceType,
		ResourcesEvaluated: request.ResourcesEvaluated,
		Tags:               request.Tags,
		Policies:           request.Policies,
		Description:        request.Description,
		CreateTime:         time.Now(),
		UpdateTime:         time.Now(),
	}
	if err := s.handler.CreateBenchmark(benchmark); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return c.JSON(http.StatusFound, "Benchmark already exists")
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create benchmark: %v", err))
	}

	response := map[string]string{
		"id": benchmark.ID,
	}

	return c.JSON(http.StatusOK, response)
}

type DeleteBenchmarkPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) DeleteBenchmark(c echo.Context) error {
	var request DeleteBenchmarkPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	benchmark, err := s.handler.QueryBenchmark(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query benchmark: %v", err))
	}
	if benchmark == nil {
		return c.JSON(http.StatusNotFound, "Benchmark not found")
	}
	if err := s.handler.DeleteBenchmark(request.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Delete benchmark %s: %v", request.ID, err))
	}

	return c.JSON(http.StatusOK, "success")
}

func (s *Server) UpdateBenchmark(c echo.Context) error {
	var request BenchmarkPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if request.ID == "" {
		return c.JSON(http.StatusBadRequest, "id is empty")
	}
	if err := s.validateBenchmarkRequest(c, &request); err != nil {
		return err
	}

	benchmark, err := s.handler.QueryBenchmark(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query benchmark: %v", err))
	}
	if benchmark == nil {
		return c.JSON(http.StatusNotFound, "Benchmark not found")
	}

	benchmark.Name = request.Name
	benchmark.State = request.State
	benchmark.Provider = request.Provider
	benchmark.ResourceType = request.ResourceType
	benchmark.ResourcesEvaluated = request.ResourcesEvaluated
	benchmark.Tags = request.Tags
	benchmark.Policies = request.Policies
	benchmark.Description = request.Description
	benchmark.UpdateTime = time.Now()
	if err := s.handler.UpdateBenchmark(benchmark); err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Update benchmark: %v", err))
	}

	return c.JSON(http.StatusOK, "success")
}

type QueryBenchmarkPOST struct {
	ID string `json:"id" validate:"required"`
}

func (s *Server) QueryBenchmark(c echo.Context) error {
	var request QueryBenchmarkPOST
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind request: %v", err))
	}

	if err := s.validate.Struct(&request); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Validate: %v", err))
	}

	benchmark, err := s.handler.QueryBenchmark(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query benchmark: %v", err))
	}
	if benchmark == nil {
		return c.JSON(http.StatusNotFound, "Benchmark not found")
	}

	return c.JSON(http.StatusOK, benchmark)
}

func (s *Server) QueryBenchmarks(c echo.Context) error {
	benchmarks, err := s.handler.QueryBenchmarks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Query benchmarks: %v", err))
	}
	if benchmarks == nil {
		return c.JSON(http.StatusNotFound, "Benchmarks not found")
	}

	return c.JSON(http.StatusOK, benchmarks)
}
