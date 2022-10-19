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
	GreaterThan map[string]int
	LessThan    map[string]int
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

	return params
}

func (this *GetOptions) Validate() error {
	if this.Offset < 0 {
		err := fmt.Errorf("Offset cannot be < 0. Given: %d", this.Offset)
		return err
	}
	return nil
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
