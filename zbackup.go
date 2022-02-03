package zsoap

type ZBackup struct {
	Label      string           `json:"label,omitempty"`
	Type       string           `json:"type,omitempty"`
	Aborted    bool             `json:"aborted,omitempty"`
	Start      int              `json:"start,omitempty"`
	End        int              `json:"end,omitempty"`
	MinRedoSeq int              `json:"minRedoSeq,omitempty"`
	MaxRedoSeq int              `json:"maxRedoSeq,omitempty"`
	Live       bool             `json:"live,omitempty"`
	Accounts   []ZBackupAccount `json:"accounts,omitempty"`
}

type ZBackupAccount struct {
	Total           int `json:"total,omitempty"`
	CompletionCount int `json:"completionCount,omitempty"`
}

func (b *ZBackup) Account() *ZBackupAccount {
	return &b.Accounts[0]
}

func (b *ZBackup) Date() string {
	runes := []rune(b.Label)
	return string(runes[5:13])
}

func (a *ZBackupAccount) DiffTotalCompletionCount() int {
	return (a.Total - a.CompletionCount)
}
