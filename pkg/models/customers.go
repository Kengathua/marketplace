package models

import (
	"github.com/Kengathua/marketplace/pkg/common"
)

type Customer struct {
	common.BioData
	CustomerNumber string
}
