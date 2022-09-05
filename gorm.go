package gk

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type FieldId struct {
	Id string `gorm:"type:char(32);not null;primaryKey" json:"id"`
}

func (i *FieldId) BeforeCreate(tx *gorm.DB) (err error) {
	i.Id = strings.ToUpper(strings.Replace(uuid.New().String(), "-", "", -1))
	return
}

type FieldTime struct {
	CreatedAt TimeNormal `gorm:"column:created_at;default:null" json:"created_at"`
	UpdatedAt TimeNormal `gorm:"column:updated_at;default:null" json:"updated_at"`
}

func (f *FieldTime) BeforeSave(tx *gorm.DB) (err error) {
	f.CreatedAt = TimeNormal{time.Now()}
	f.UpdatedAt = TimeNormal{time.Now()}
	return
}

type TimeNormal struct { // 内嵌方式（推荐）
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// gorm 条件查询
func WhereBuilder(where map[string]map[string]interface{}) (whereSql string, values []interface{}, err error) {
	var whereSlice []string
	for field, queryBuilder := range where {
		for op, value := range queryBuilder {
			switch op {
			case "=", ">", ">=", "<", "<=", "<>", "!=", "in", "not in":
				whereSlice = append(whereSlice, fmt.Sprintf("%s %s ?", field, op))
				values = append(values, value)
				break
			case "like", "not like":
				whereSlice = append(whereSlice, fmt.Sprintf("%s %s ?", field, op))
				values = append(values, value)
				break
			case "null":
				whereSlice = append(whereSlice, fmt.Sprintf("%s is null and ", field))
				break
			case "not null":
				whereSlice = append(whereSlice, fmt.Sprintf("%s is not null and ", field))
				break
			default:
				return "", nil, errors.New("error in query condition")
			}
		}
	}

	return strings.Join(whereSlice, " and "), values, nil
}
