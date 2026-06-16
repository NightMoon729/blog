package Service

import (
	"blog/backend/src/entity"
	"blog/backend/src/repository"
)

type SolutionService interface { //对外使用的接口
	GetSolutions(page, N int) ([]entity.Solution, error)
}
type SolutionServiceImpl struct { //实现类
	repo *repository.SolutionRepo
}

func NewSolutionService(repo *repository.SolutionRepo) SolutionService { //构造函数
	return &SolutionServiceImpl{repo: repo}
}
func (s *SolutionServiceImpl) GetSolutions(page, N int) ([]entity.Solution, error) { //实现类方法
	return s.repo.GetSolutionsByPage(page, N)
}
