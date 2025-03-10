package stat

import (
	"gorm.io/datatypes"
	"http/test/pkg/db"
	"time"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())

	repo.Db.First(&stat, "link_id = ? and date = ?", linkId, currentDate)

	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.Db.Save(&stat)
	}

}
