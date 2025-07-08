package diary

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, diary *Diary) (int64, error)
	Update(ctx context.Context, diary *Diary) (int64, error)
	Delete(ctx context.Context, id int64) error
}
