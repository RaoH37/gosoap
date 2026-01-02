package zimbraAccount

import "github.com/RaoH37/gosoap/zimbraCommon"

type AuthRequest struct {
	Content AuthRequestContent `json:"AuthRequest"`
}

type AuthRequestContent struct {
	Account  zimbraCommon.ByNode `json:"account,attr"`
	Password string              `json:"password,attr,omitempty"`
	Preauth  Preauth             `json:"preauth,omitzero"`
	Urn      string              `json:"_jsns,attr"`
}

type Preauth struct {
	Timestamp int64  `json:"timestamp,attr"`
	Expires   int    `json:"expires,attr"`
	Value     string `json:"_content,attr"`
}

type AuthResponse struct {
	Content AuthResponseContent `json:"AuthResponse"`
}

type AuthResponseContent struct {
	TOKEN    []AuthResponseToken `json:"authToken"`
	Lifetime int                 `json:"lifetime"`
}

type AuthResponseToken struct {
	Content string `json:"_content"`
}

type GetInfoRequest struct {
	Content GetInfoRequestContent `json:"GetInfoRequest"`
}

type GetInfoRequestContent struct {
	Urn      string `json:"_jsns,attr"`
	Rights   string `json:"rights,attr,omitempty"`
	Sections string `json:"sections,attr,omitempty"`
}

type GetInfoResponse struct {
	Content GetInfoResponseContent `json:"GetInfoResponse"`
}

type GetInfoResponseContent struct {
	Version   string                  `json:"version"`
	ID        string                  `json:"id"`
	Name      string                  `json:"name"`
	Rest      string                  `json:"rest"`
	PublicURL string                  `json:"publicUrl"`
	Used      int                     `json:"used"`
	Attributs []zimbraCommon.AttrNode `json:"attr"`
	Prefs     []zimbraCommon.AttrNode `json:"pref"`
	Cos       zimbraCommon.IdNameNode `json:"cos"`
}

type CreateIdentityRequest struct {
	Content CreateIdentityRequestContent `json:"CreateIdentityRequest"`
}

type CreateIdentityRequestContent struct {
	Urn      string         `json:"_jsns,attr"`
	Identity ZimbraIdentity `json:"identity"`
}

type ZimbraIdentity struct {
	ID        string                  `json:"id,attr,omitempty"`
	Name      string                  `json:"name,attr"`
	Attributs []zimbraCommon.NameNode `json:"a,omitempty"`
}

type CreateIdentityResponse struct {
	Content struct {
		Identity ZimbraIdentity `json:"identity"`
	} `json:"CreateIdentityResponse"`
}

type DeleteIdentityRequest struct {
	Content DeleteIdentityRequestContent `json:"DeleteIdentityRequest"`
}

type DeleteIdentityRequestContent struct {
	Urn      string `json:"_jsns,attr"`
	Identity struct {
		ID string `json:"id,attr,omitempty"`
	} `json:"identity"`
}

type DeleteIdentityResponse struct {
	Content struct{} `json:"DeleteIdentityResponse"`
}

type CreateSignatureRequest struct {
	Content CreateSignatureRequestContent `json:"CreateSignatureRequest"`
}

type CreateSignatureRequestContent struct {
	Urn       string          `json:"_jsns,attr"`
	Signature ZimbraSignature `json:"signature"`
}

type ZimbraSignature struct {
	ID      string                   `json:"id,attr,omitempty"`
	Name    string                   `json:"name,attr"`
	Content []ZimbraSignatureContent `json:"content,omitempty"`
}

type ZimbraSignatureContent struct {
	Value       string `json:"_content"`
	ContentType string `json:"type,attr"`
}

type CreateSignatureResponse struct {
	Content struct {
		Signature ZimbraSignature `json:"signature"`
	} `json:"CreateSignatureResponse"`
}

type DeleteSignatureRequest struct {
	Content DeleteSignatureContent `json:"DeleteSignatureRequest"`
}

type DeleteSignatureContent struct {
	Urn       string `json:"_jsns,attr"`
	Signature struct {
		ID string `json:"id,attr,omitempty"`
	} `json:"signature"`
}

type DeleteSignatureResponse struct {
	Content struct{} `json:"DeleteSignatureResponse"`
}

type ModifyIdentityRequest struct {
	Content ModifyIdentityRequestContent `json:"ModifyIdentityRequest"`
}

type ModifyIdentityRequestContent struct {
	Urn      string         `json:"_jsns,attr"`
	Identity ZimbraIdentity `json:"identity"` // Réutilise la struct du Create
}

type ModifyIdentityResponse struct {
	Content struct{} `json:"ModifyIdentityResponse"`
}

type ModifySignatureRequest struct {
	Content ModifySignatureRequestContent `json:"ModifySignatureRequest"`
}

type ModifySignatureRequestContent struct {
	Urn       string          `json:"_jsns,attr"`
	Signature ZimbraSignature `json:"signature"` // Réutilise la struct du Create
}

type ModifySignatureResponse struct {
	Content struct{} `json:"ModifySignatureResponse"`
}

type ModifyPrefsRequest struct {
	Content ModifyPrefsRequestContent `json:"ModifyPrefsRequest"`
}

type ModifyPrefsRequestContent struct {
	Urn   string                  `json:"_jsns,attr"`
	Prefs []zimbraCommon.AttrNode `json:"pref"`
}

type ModifyPrefsResponse struct {
	Content struct{} `json:"ModifyPrefsResponse"`
}

type GrantRightsRequest struct {
	Content GrantRightsRequestContent `json:"GrantRightsRequest"`
}

type GrantRightsRequestContent struct {
	Urn    string                   `json:"_jsns,attr"`
	Rights []zimbraCommon.GrantNode `json:"ace"`
}

type GrantRightsResponse struct {
	Content struct {
		Rights []zimbraCommon.GrantNode `json:"ace"`
	} `json:"GrantRightsResponse"`
}

type RevokeRightsRequest struct {
	Content RevokeRightsContent `json:"RevokeRightsRequest"`
}

type RevokeRightsContent struct {
	Urn    string                   `json:"_jsns,attr"`
	Rights []zimbraCommon.GrantNode `json:"ace"`
}

type RevokeRightsResponse struct {
	Content struct{} `json:"RevokeRightsResponse"`
}
