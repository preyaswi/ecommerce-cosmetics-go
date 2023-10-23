package domain

import "time"

type image struct {
	Img_id           uint
	OriginalFilename string
	CroppedFilename  string
	UploadTime       time.Time
	
}
