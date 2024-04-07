package domain

// MARK: - Objects

var logSeparator = "."

type LogScope struct {
	Parent *LogScope

	Identifier string
}

// MARK: - Functions

func (l LogScope) String() string {
	if l.Parent != nil {
		return l.Parent.String() + logSeparator + l.Identifier
	}
	return l.Identifier
}
