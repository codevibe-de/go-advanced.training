package service_struct

import (
	"fmt"
)

type DataService struct {
	prefix string
}

func (s *DataService) GetData(id int) (string, error) {
	return fmt.Sprintf("%s :: Real Data #%d", s.prefix, id), nil
}
