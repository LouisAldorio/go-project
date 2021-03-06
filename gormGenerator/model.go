package model

//User User Struct
type FileUpload struct {
	ID        int     `gorm:"type:int;AUTO_INCREMENT;not null"`
	FileID    string  `gorm:"type:text;not null"`
	CreatedAt string  `gorm:"type:timestamp;not null"`
	UpdatedAt *string `gorm:"type:timestamp;null;default:null"`
	DeletedAt *string `gorm:"type:timestamp;null;default:null"`
	UserID    int     `gorm:"type:int;not null"`
}
