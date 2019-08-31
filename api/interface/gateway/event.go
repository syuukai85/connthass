package gateway

import (
	"fmt"
	"net/http"

	"github.com/connthass/connthass/api/entity"
	"github.com/connthass/connthass/api/infrastructure/orm"
	"github.com/connthass/connthass/api/infrastructure/orm/model"
	"github.com/jinzhu/gorm"
)

const (
	eventNotFoundError = "イベントが見つかりませんでした"
)

// Event DBモデルとやり取りをする
type Event struct {
	db *gorm.DB
}

// NewEvent コンストラクタ
func NewEvent() *Event {
	return &Event{
		db: orm.GetDB(),
	}
}

// SearchEvents 検索条件からイベントを検索する
func (e *Event) SearchEvents(fields entity.Fields, query entity.Query, page entity.Page, perPage entity.PerPage) ([]entity.Event, *entity.Error) {
	// fmt.Println(query)
	// events := []model.Event{}
	// e.db.Offset(page).Limit(perPage).Select(fields).Find(&events)

	// TODO: モデルからエンティティに変換する

	return make([]entity.Event, 0), nil
}

// FindByID IDからイベントを検索する
func (e *Event) FindByID(eventID entity.EventID) (*entity.Event, *entity.Error) {
	var event model.Event
	var group model.Group
	var venue model.Venue
	var categories []model.Category

	stringEventID := fmt.Sprint(eventID)
	firstEvent := e.db.First(&event, stringEventID)

	if firstEvent.RecordNotFound() {
		return nil, &entity.Error{
			Code:   http.StatusNotFound,
			Errors: []string{eventNotFoundError}}
	}

	firstEvent.Related(&group)
	firstEvent.Related(&venue)
	firstEvent.Related(&categories, "Categories")

	gatewayUser := NewUser()
	entityEvent := &entity.Event{
		ID:               entity.EventID(stringEventID),
		ColorCode:        event.ColorCode,
		Title:            event.Title,
		Description:      event.Description,
		Capacity:         event.Capacity,
		ImageURL:         event.ImageURL,
		QRCodeURL:        event.QRCodeURL,
		HoldStartDate:    event.HoldStartDate,
		HoldEndDate:      event.HoldEndDate,
		RecruitStartDate: event.RecruitStartDate,
		RecruitEndDate:   event.RecruitEndDate,
		Group:            groupToEntity(group),
		Venue:            venueToEntity(venue),
		Entries:          gatewayUser.findGeneralByEventID(stringEventID),
		Organizer:        gatewayUser.findOrganizerByEventID(stringEventID),
		Categories:       categoriesToEntities(categories),
	}

	return entityEvent, nil
}
