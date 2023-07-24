package models

import (
	"github.com/Kengathua/marketplace/pkg/common"
)

type BusinessPartner struct {
	common.Base
	Name           string
	BP_Code        string
	MainBranchCode string
	Description    string
}
