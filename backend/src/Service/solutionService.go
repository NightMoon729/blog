package Service

import "blog/backend/src/entity"

type SolutionService interface {
	GetListSolutions(page, size int) ([]entity.Solution, error)
}
