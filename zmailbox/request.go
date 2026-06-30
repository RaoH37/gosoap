package zmailbox

import (
	"strings"

	"github.com/RaoH37/gosoap/zimbraAccount"
	"github.com/RaoH37/gosoap/zimbraCommon"
	"github.com/RaoH37/gosoap/zimbraMail"
)

func (s *ZMailbox) AuthRequest() (*zimbraAccount.AuthResponse, error) {
	if s.password != "" {
		return s.AuthRequestByPassword()
	}

	return s.AuthRequestByPreauth(0)
}

func (s *ZMailbox) AuthRequestByPassword() (*zimbraAccount.AuthResponse, error) {
	req, resp := zimbraAccount.NewAuthRequestByPassword(zimbraCommon.NewByIdOrNameNode(s.id, s.name), s.password)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) AuthRequestByPreauth(expires int) (*zimbraAccount.AuthResponse, error) {
	by := zimbraCommon.NewByIdOrNameNode(s.id, s.name)

	preauth := zimbraAccount.NewPreauth(by, expires, s.domainKey)

	req, resp := zimbraAccount.NewAuthRequestByPreauth(by, preauth)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) CreateFolderRequest(name string, view zimbraMail.FolderView, parentId string, options func(*zimbraMail.CreateFolderRequest)) (*zimbraMail.CreateFolderResponse, error) {
	req, resp := zimbraMail.NewCreateFolderRequest(name, view, parentId)

	if options != nil {
		options(req)
	}

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) GetFolderRequest(view zimbraMail.FolderView) (*zimbraMail.GetFolderResponse, error) {
	req, resp := zimbraMail.NewGetFolderRequest(view, false, false, false)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) GetInfoRequest(rights zimbraCommon.StringList, sections zimbraAccount.InfoSectionList) (*zimbraAccount.GetInfoResponse, error) {
	req, resp := zimbraAccount.NewGetInfoRequest(rights, sections)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) SearchRequest(types []string, offset int, limit int, query string) (*zimbraMail.SearchResponse, error) {
	req, resp := zimbraMail.NewSearchRequest(strings.Join(types, ","), offset, limit, query)

	if err := s.Connector.Invoke(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
