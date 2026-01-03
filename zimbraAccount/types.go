package zimbraAccount

import (
	"encoding/json"
	"strings"
)

type InfoSection string

const (
	InfoSectionMbox     InfoSection = "mbox"
	InfoSectionPrefs    InfoSection = "prefs"
	InfoSectionAttrs    InfoSection = "attrs"
	InfoSectionZimlets  InfoSection = "zimlets"
	InfoSectionProps    InfoSection = "props"
	InfoSectionIdents   InfoSection = "idents"
	InfoSectionSigs     InfoSection = "sigs"
	InfoSectionDsrcs    InfoSection = "dsrcs"
	InfoSectionChildren InfoSection = "children"
)

type InfoSectionList []InfoSection

func (l InfoSectionList) String() string {
	if len(l) == 0 {
		return ""
	}
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = string(v)
	}
	return strings.Join(s, ",")
}

func (l InfoSectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
