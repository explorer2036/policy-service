package db

import (
	"policy-service/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableBenchmark - database table 'benchmark'
	TableBenchmark = "benchmark"
)

// QueryBenchmarks returns the benchmarks
func (s *Handler) QueryBenchmarks() ([]*model.Benchmark, error) {
	benchmarks := []*model.Benchmark{}

	res := s.db.Table(TableBenchmark).Find(&benchmarks)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return benchmarks, nil
}

// QueryBenchmark returns the benchmark by 'id'
func (s *Handler) QueryBenchmark(id string) (*model.Benchmark, error) {
	benchmark := model.Benchmark{}

	res := s.db.Table(TableBenchmark).Where("id = ?", id).First(&benchmark)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &benchmark, nil
}

// CreateBenchmark inserts a new benchmark
func (s *Handler) CreateBenchmark(benchmark *model.Benchmark) error {
	return s.db.Table(TableBenchmark).Create(benchmark).Error
}

// UpdateBenchmark updates a special benchmark
func (s *Handler) UpdateBenchmark(benchmark *model.Benchmark) error {
	return s.db.Table(TableBenchmark).Update(benchmark).Error
}

// DeleteBenchmark deletes the benchmark by id
func (s *Handler) DeleteBenchmark(id string) error {
	return s.db.Table(TableBenchmark).Exec("Delete from benchmark where id = ?", id).Error
}
