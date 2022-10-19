package lotr

import (
	"fmt"
	"strings"
)

type GetOptions struct {
	Limit  int
	Offset int
	Page   int

	Match       map[string][]string // param -> values. Ex - race -> Elf,Human
	NotMatch    map[string][]string
	GreaterThan map[string]int // key > value
	LessThan    map[string]int // key < value

	SortKey string // key to sort on
	SortAsc bool   // order for the sort, defaults false
}

func (this *GetOptions) GetParams() map[string]string {
	params := make(map[string]string)

	if this.Limit > 0 {
		params["limit"] = fmt.Sprintf("%d", this.Limit)
	}

	if this.Offset > 0 {
		params["offset"] = fmt.Sprintf("%d", this.Offset)
	}

	if this.Page > 0 {
		params["page"] = fmt.Sprintf("%d", this.Page)
	}

	if this.Match != nil {
		for key, values := range this.Match {
			params[key] = strings.Join(values, ",")
		}
	}

	if this.NotMatch != nil {
		for key, values := range this.NotMatch {
			params[key+"!"] = strings.Join(values, ",")
		}
	}

	if this.GreaterThan != nil {
		for key, val := range this.GreaterThan {
			params[key+">"+fmt.Sprintf("%d", val)] = ""
		}
	}

	if this.LessThan != nil {
		for key, val := range this.LessThan {
			params[key+"<"+fmt.Sprintf("%d", val)] = ""
		}
	}

	if this.SortKey != "" {
		sortType := "desc"
		if this.SortAsc {
			sortType = "asc"
		}
		params["sort"] = fmt.Sprintf("%s:%s", this.SortKey, sortType)
	}

	return params
}

func (this *GetOptions) Validate() error {
	if this.Offset < 0 {
		err := fmt.Errorf("Offset cannot be < 0. Given: %d", this.Offset)
		return err
	}
	return nil
}

func (this *GetOptions) SortOnKey(key string, asc bool) {
	this.SortKey = key
	this.SortAsc = asc
}

func NewGetOptions() *GetOptions {
	options := new(GetOptions)
	return options
}

func NewGetOptionsOffset(offset int) *GetOptions {
	options := NewGetOptions()
	options.Offset = offset
	return options
}

func NewGetOptionsOffsetLimited(offset int, limit int) *GetOptions {
	options := NewGetOptions()
	options.Offset = offset
	options.Limit = limit
	return options
}

func NewGetOptionsPage(page int) *GetOptions {
	options := NewGetOptions()
	options.Page = page
	return options
}

func NewGetOptionsPageLimited(page int, limit int) *GetOptions {
	options := NewGetOptions()
	options.Page = page
	options.Limit = limit
	return options
}
