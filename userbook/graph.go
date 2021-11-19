package userbook

import (
	"bookstop/graph/model"
	"strconv"

	"github.com/jackc/pgtype"
)

func ToGraph(userBook *UserBook) *model.UserBook {
	if userBook == nil {
		return nil
	}

	val := model.UserBook{
		ID:     strconv.Itoa(int(userBook.ID.Int)),
		BookID: userBook.BookId.String,
		UserID: strconv.Itoa(int(userBook.UserId.Int)),
	}

	if userBook.StartedAt.Status == pgtype.Present {
		startedAt := userBook.StartedAt.Time.Format("2006-01-02")
		val.StartedAt = &startedAt
	}

	if userBook.EndedAt.Status == pgtype.Present {
		endedAt := userBook.EndedAt.Time.Format("2006-01-02")
		val.EndedAt = &endedAt
	}

	if userBook.IDOriginal.Status == pgtype.Present {
		originalId := strconv.Itoa(int(userBook.IDOriginal.Int))
		val.OriginalUserBookID = &originalId
	}

	return &val
}