// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameZxAttachfile = "zx_attachfile"

// ZxAttachfile mapped from table <zx_attachfile>
type ZxAttachfile struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InvID    string `gorm:"column:inv_id" json:"inv_id"`
	FileName string `gorm:"column:file_name" json:"file_name"`
	FilePath string `gorm:"column:file_path" json:"file_path"`
}

// TableName ZxAttachfile's table name
func (*ZxAttachfile) TableName() string {
	return TableNameZxAttachfile
}
