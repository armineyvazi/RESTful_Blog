package di

import (
	"gorm.io/gorm"

	cc "restful_blog/pkg/category/controller"
	cr "restful_blog/pkg/category/repository"
	cs "restful_blog/pkg/category/service"

	pc "restful_blog/pkg/post/controller"
	pr "restful_blog/pkg/post/repository"
	ps "restful_blog/pkg/post/service"
)

type Container struct {
	CategoryController *cc.CategoryController
	PostController     *pc.PostController
}

func NewContainer(db *gorm.DB) *Container {

	// Create new instance category repository
	categoryRepo := cr.NewCategoryRepository(db)

	// Create new instance category service
	categoryService := cs.NewCategoryService(categoryRepo)

	// Create new instance category controller
	categoryController := cc.NewCategoryController(categoryService)

	// Create new instance post repository
	postRepo := pr.NewPostRepository(db)

	// Create new instance post service
	postService := ps.NewPostService(postRepo)

	// Create new instance post controller
	postController := pc.NewPostController(postService)

	return &Container{
		CategoryController: categoryController,
		PostController:     postController,
	}
}
