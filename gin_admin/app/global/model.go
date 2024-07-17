package global


type GVA_MODEL struct {
	ID      uint  `gorm:"primarykey" json:"id"` // 主键ID
	Updated int64 `gorm:"autoUpdateTime:milli" json:"update_at"` // 使用时间戳毫秒数填充更新时间
	Created int64 `gorm:"autoCreateTime:milli" json:"create_at"`       // 使用时间戳秒数填充创建时间
}
