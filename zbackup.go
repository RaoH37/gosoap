package zsoap

import (
	"log"

	"github.com/RaoH37/gosoap/zimbraAdmin"
)

func (s *ZcsClient) GetAllBackups() ([]zimbraAdmin.ZBackup, error) {
	resp, err := s.ZAdminClient.BackupQueryRequest()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.Content.Backups, nil
}
