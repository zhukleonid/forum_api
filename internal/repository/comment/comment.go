package comment

import (
	"context"
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

const (
	updateCommQuery  = `UPDATE comments SET description = $1, updated_at = $2 WHERE id = $3`
	deleteCommQuery  = `DELETE FROM comments WHERE id = $1 AND user_id = $2`
	commentByIDQuery = `SELECT * FROM comments WHERE id = $1`
	commentsQuery    = `SELECT * FROM comments`
	createCommQuery  = `INSERT INTO comments (post_id, user_id, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepo(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

func (repo *CommentsRepository) CreateComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, createCommQuery, comm.Post, comm.User, comm.Description, comm.CreatedDate, comm.UpdatedDate); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) UpdateComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, updateCommQuery, comm.Description, comm.UpdatedDate, comm.ID); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) DeleteComment(ctx context.Context, comm *model.Comment) error {
	if _, err := repo.db.ExecContext(ctx, deleteCommQuery, comm.ID, comm.User); err != nil {
		return err
	}
	return nil
}

func (repo *CommentsRepository) CommentByID(ctx context.Context, id int) (*model.Comment, error) {
	var comm model.Comment
	if err := repo.db.QueryRowContext(ctx, commentByIDQuery, id).Scan(&comm.ID, &comm.Post, &comm.User, &comm.Description, &comm.CreatedDate, &comm.UpdatedDate); err != nil {
		return nil, err
	}
	return &comm, nil
} 