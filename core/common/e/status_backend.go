package e

import "encoding/hex"

//多维度状态值创建、获取某一维度、设置某一维度
type STATUS []byte

const (
	runIndex = iota
	initIndex
	kindIndex
	anchorIndex

	statusLen = 4
)

func (s STATUS) String() string {
	return hex.EncodeToString(s)
}

func (s STATUS) SetRunStatus(run byte) STATUS {
	s[runIndex] = run
	return s
}

func (s STATUS) SetInitStatus(init byte) STATUS {
	s[initIndex] = init
	return s
}

func (s STATUS) SetKindStatus(kind byte) STATUS {
	s[kindIndex] = kind
	return s
}

func (s STATUS) SetAnchorStatus(anchor byte) STATUS {
	s[anchorIndex] = anchor
	return s
}

func NewStatus(run, init, kind, anchor byte) STATUS {
	var s STATUS = make([]byte, statusLen)
	s.SetRunStatus(run)
	s.SetInitStatus(init)
	s.SetKindStatus(kind)
	s.SetAnchorStatus(anchor)
	return s
}

func StatusDecode(status string) STATUS {
	var s STATUS
	var err error
	s, err = hex.DecodeString(status)
	if err != nil {
		panic(err)
	}
	return s
}
