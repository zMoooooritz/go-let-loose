package config

const (
	LIST_DELIMITER  = "\t"
	NEWLINE         = "\n"
	ESCAPED_NEWLINE = "\\n"
)

type ResponseFormat int

const (
	RF_DIRECT ResponseFormat = iota
	RF_INDEXEDLIST
	RF_UNINDEXEDLIST
)
