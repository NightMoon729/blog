package repository

import (
	"blog/backend/src/entity"
	"github.com/jmoiron/sqlx"
)

type solutionRepo struct {
	db *sqlx.DB
}

func (r *solutionRepo) GetSolutionsByPage(page, size int) ([]entity.Solution, error) {
	var solutions []entity.Solution //创建一个solution列表
	query := `
        SELECT id, title, tags, web 
        FROM db_Solution 
        ORDER BY id DESC 
        LIMIT ? OFFSET ?
        ` //降序排序查询
	err := r.db.Select(&solutions, query, size, page*10) //用sqlx的语法
	if err != nil {
		return nil, err
	}
	return solutions, nil
}
