package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/usagifm/golang-jwt/dto"
	"github.com/usagifm/golang-jwt/entity"
	"github.com/usagifm/golang-jwt/repository"
)

type DiaryService interface {
	Insert(d dto.DiaryCreateDTO) entity.Diary
	Update(d dto.DiaryUpdateDTO) entity.Diary
	Delete(d entity.Diary)
	All() []entity.Diary
	FindById(diaryID uint64) entity.Diary
	IsAllowedToEdit(userID string, diaryID uint64) bool
}

type diaryService struct {
	diaryRepository repository.DiaryRepository
}

func NewDiaryService(diaryRepo repository.DiaryRepository) DiaryService {
	return &diaryService{
		diaryRepository: diaryRepo,
	}
}

func (service *diaryService) Insert(d dto.DiaryCreateDTO) entity.Diary {
	diary := entity.Diary{}
	err := smapping.FillStruct(&diary, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed to map %v: ", err)
	}
	res := service.diaryRepository.InsertDiary(diary)
	return res

}

func (service *diaryService) Update(d dto.DiaryUpdateDTO) entity.Diary {
	diary := entity.Diary{}
	err := smapping.FillStruct(&diary, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed to map %v: ", err)
	}
	res := service.diaryRepository.UpdateDiary(diary)
	return res
}

func (service *diaryService) Delete(d entity.Diary) {
	service.diaryRepository.DeleteDiary(d)
}

func (service *diaryService) All() []entity.Diary {
	return service.diaryRepository.AllDiary()
}

func (service *diaryService) FindById(diaryID uint64) entity.Diary {
	return service.diaryRepository.FindDiaryById(diaryID)
}

func (service *diaryService) IsAllowedToEdit(userID string, diaryID uint64) bool {
	d := service.diaryRepository.FindDiaryById(diaryID)
	id := fmt.Sprintf("%v", d.UserID)
	return userID == id
}
