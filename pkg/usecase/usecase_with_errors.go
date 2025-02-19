package usecase

import (
	"github.com/alexdogonin/errors-handling/pkg/common"
	"github.com/pkg/errors"
)

type Repository interface {
	GetByID(int) error
}

type UsecaseWithErrors struct {
	repo Repository
}

func New(repo Repository) *UsecaseWithErrors {
	return &UsecaseWithErrors{repo}
}

func (u *UsecaseWithErrors) ProcessByID(id int) error {
	err := u.repo.GetByID(id)
	if err != nil {
		// демонстрация различного поведения в зависимости от ошибок, пришедших снизу по стеку
		var e common.ErrNotFound

		// никакой другой тип ошибки не пройдет проверку, пруф:
		errr := errors.New("not found some records")
		if errors.As(errr, &e) {
			panic("не сработало, проверка типа пропустила левый тип!")
		}

		if errors.As(err, &e) {
			// допустим, если это ошибка not found типа, то мы ее пробрасываем выше
			return errors.Wrapf(err, "process by id, id %d creates a conflict", id)
		}

		// здесь будет креш
		if errors.Is(err, common.VsePipetz) {
			panic("не могу обработать!!!")
		}

		return err
	}

	return nil
}
