package zmailbox

import (
	"log"

	"github.com/RaoH37/gosoap/zimbraAccount"
	"github.com/RaoH37/gosoap/zimbraMail"
)

func (s *ZMailbox) AuthRequest() (*zimbraAccount.AuthResponse, error) {
	if s.password != "" {
		return s.AuthRequestByPassword()
	}

	return s.AuthRequestByPreauth(0)
}

func (s *ZMailbox) AuthRequestByPassword() (*zimbraAccount.AuthResponse, error) {
	req, resp := zimbraAccount.NewAuthRequestByPassword(s.authRequestByNode(), s.password)

	connector := s.BuildZimbraConnector()
	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) AuthRequestByPreauth(expires int) (*zimbraAccount.AuthResponse, error) {
	by := s.authRequestByNode()

	preauth := zimbraAccount.NewPreauth(by, expires, s.domainKey)

	req, resp := zimbraAccount.NewAuthRequestByPreauth(by, preauth)

	connector := s.BuildZimbraConnector()
	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) GetFolderRequest(view zimbraMail.FolderView) (*zimbraMail.GetFolderResponse, error) {
	req, resp := zimbraMail.NewGetFolderRequest(view, false, false, false)

	connector := s.BuildZimbraConnectorLogged()

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (s *ZMailbox) GetInfoRequest(rights []string, sections zimbraAccount.InfoSectionList) (*zimbraAccount.GetInfoResponse, error) {
	req, resp := zimbraAccount.NewGetInfoRequest(rights, sections)

	connector := s.BuildZimbraConnectorLogged()

	if err := connector.Invoke(req, &resp); err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}
