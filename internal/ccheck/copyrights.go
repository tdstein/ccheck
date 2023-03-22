package ccheck

type CCheckCopyright struct {
	text *[]string
}

func NewCCheckCopyright(text *[]string) *CCheckCopyright {
	return &CCheckCopyright{
		text: text,
	}
}
