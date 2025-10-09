package zimbraAdmin

const urnAdmin = "urn:zimbraAdmin"

//func urn(name string) string {
//	return urnAdmin + "/" + name
//}

func convertToContentStrings(arr []string) []ContentString {
	contentStrings := make([]ContentString, len(arr))

	for i, a := range arr {
		contentStrings[i] = ContentString{Content: a}
	}

	return contentStrings
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
