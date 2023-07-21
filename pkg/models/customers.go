package models

import (
	"github.com/matawis/matawis/pkg/common"
)

type Customer struct {
	common.BioData
	CustomerNumber string
}
