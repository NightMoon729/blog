package repository

import (
	"blog/backend/src/entity"
	"github.com/jmoiron/sqlx"
)

type SolutionRepo struct {
	DB *sqlx.DB
}

func NewSolutionRepo(db *sqlx.DB) *SolutionRepo {
	return &SolutionRepo{DB: db}
}
func (r *SolutionRepo) GetSolutionsByPage(page, size int) ([]entity.Solution, error) {
	var solutions []entity.Solution //创建一个solution列表
	query := `
        SELECT id, title, tags, web 
        FROM db_Solution 
        ORDER BY id DESC 
        LIMIT ? OFFSET ?
        ` //降序排序查询
	err := r.DB.Select(&solutions, query, size, page*size) //用sqlx的语法
	if err != nil {
		return nil, err
	}
	return solutions, nil
}
