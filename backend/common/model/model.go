package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Page[T any] struct {
	PageNum   int64 `json:"pageNum" form:"pageNum"`   // 页数
	PageSize  int64 `json:"pageSize" form:"pageSize"` // 页码
	Total     int64 `json:"total"`                    // 总数  由服务器返回回去
	Condition T     `json:"condition"`                // 条件
}

func NewPage[T any](condition T) *Page[T] {
	return &Page[T]{Condition: condition}
}

type Base struct {
	ID        string    `gorm:"column:id;size:100;comment:编码;primaryKey;autoIncrement:false;" json:"id"` // 编码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间;" json:"createdAt"`          // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:修改时间;" json:"updatedAt"`          // 修改时间
}

// 数组数据类型 数据中以逗号分隔的字符串
type Arr []string

// 实现数据库驱动接口
func (a *Arr) Scan(value interface{}) error {
	// 数据库驱动将数据转换为字符串
	if v, ok := value.(string); ok {

		*a = Arr(strings.Split(v, ","))
		return nil
	}

	return fmt.Errorf("类型错误")
}

func (a *Arr) Value() (driver.Value, error) {
	return strings.Join(*a, ","), nil
}

func (a *Arr) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`["%s"]`, strings.Join(*a, `","`))), nil
}

func (a *Arr) UnmarshalJSON(data []byte) error {
	*a = Arr(strings.Split(string(data[2:len(data)-2]), `","`))
	return nil
}
