package entity

// Bookmark おきにいりイベントの関連テーブル
type Bookmark struct {
	Base
	UserID  uint
	EventID uint
}
