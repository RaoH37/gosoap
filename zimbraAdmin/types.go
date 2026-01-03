package zimbraAdmin

import (
	"encoding/json"
	"strings"
)

type SearchType string

const (
	SearchTypeAccounts          SearchType = "accounts"
	SearchTypeDistributionLists SearchType = "distributionlists"
	SearchTypeAliases           SearchType = "aliases"
	SearchTypeResources         SearchType = "resources"
	SearchTypeDomains           SearchType = "domains"
	SearchTypeCoses             SearchType = "coses"
)

type SearchTypeList []SearchType

func (l SearchTypeList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

func (l SearchTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
