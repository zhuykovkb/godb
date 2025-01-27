package model

type CmdType string

const (
	CmdTypeSet CmdType = "SET"
	CmdTypeGet CmdType = "GET"
	CmdTypeDel CmdType = "DEL"
)
