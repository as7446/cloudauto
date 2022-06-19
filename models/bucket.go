package models

const (
	User_Delete_True  = 1
	User_Delete_False = 0
)

type OSSProjBucket struct {
}

func (OSSProjBucket) TableName() string {
	return "bucket"
}

func (o *OSSProjBucket) FindAll() {

}