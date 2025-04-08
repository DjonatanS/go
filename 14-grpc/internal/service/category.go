package service

import (
	"context"
	"database/sql"

	"grpc/internal/database"
	"grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{
		CategoryDB: database.NewCategory(db),
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.Categories, error) {
	categories, err := s.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.Categories{Categories: pbCategories}, nil
}
