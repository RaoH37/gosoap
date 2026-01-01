package zimbraCommon

const ID = "id"
const NAME = "name"

type ByNode struct {
	By    string `json:"by"`
	Value string `json:"_content"`
}

func NewByNode(by string, value string) ByNode {
	return ByNode{
		By:    by,
		Value: value,
	}
}
