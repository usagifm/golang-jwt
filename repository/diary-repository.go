package repository

import (
	"github.com/usagifm/golang-jwt/entity"
	"gorm.io/gorm"
)

type DiaryRepository interface {
	InsertDiary(d entity.Diary) entity.Diary
	UpdateDiary(d entity.Diary) entity.Diary
	DeleteDiary(d entity.Diary)
	AllDiary() []entity.Diary
	FindDiaryById(diaryID uint64) entity.Diary
}

type diaryConnection struct {
	connection *gorm.DB
}

func NewDiaryRepository(dbConn *gorm.DB) DiaryRepository {
	return &diaryConnection{
		connection: dbConn,
	}
}

func (db *diaryConnection) InsertDiary(d entity.Diary) entity.Diary {

	db.connection.Save(&d)
	db.connection.Preload("User").Find(&d)
	return d
}

func (db *diaryConnection) UpdateDiary(d entity.Diary) entity.Diary {

	db.connection.Save(&d)
	db.connection.Preload("User").Find(&d)
	return d

}
func (db *diaryConnection) DeleteDiary(d entity.Diary) {
	db.connection.Delete(&d)
}

func (db *diaryConnection) FindDiaryById(diaryID uint64) entity.Diary {
	var diary entity.Diary
	db.connection.Preload("User").Find(&diary, diaryID)
	return diary
}

func (db *diaryConnection) AllDiary() []entity.Diary {
	var diaries []entity.Diary
	db.connection.Preload("User").Find(&diaries)
	return diaries
}
