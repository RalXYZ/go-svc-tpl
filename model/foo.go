package model

type Foo struct {
	Bar uint64 `gorm:"not null;autoIncrement;primaryKey"`
	Buz string `gorm:"not null"`
}
