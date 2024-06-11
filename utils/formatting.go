package f

type Colors_ struct {
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Purple    string
	Cyan      string
	White     string
	Magenta   string
	Reset     string
	Brown     string
	CadetBlue string
	Plum      string
	LightPink string
}

type Font_ struct {
	Bold      string
	Underline string
	Italic    string
	Reset     string
}

type Icons_ struct {
	Success string
	Warning string
	Error   string
	Info    string
	Pkg     string
	Subtest string
	Output  string
}

var Font = Font_{
	Bold:      "\033[1m",
	Underline: "\033[4m",
	Italic:    "\033[3m",
	Reset:     "\033[0m",
}

var Icons = Icons_{
	Success: "\uf058 \u00A0",
	Pkg:     "\ueb29 \u00A0",
	Warning: "\u26A0 \u00A0",
	Error:   "\u2718 \u00A0",
	Info:    "\u2139 \u00A0",
	Subtest: "\uf061 \u00A0",
	Output:  "\ueb9d \u00A0",
}

var Colors = Colors_{
	Red:       "\033[38;5;196m",
	Green:     "\033[38;5;34m",
	Yellow:    "\033[38;5;100m",
	Blue:      "\033[38;5;21m",
	Purple:    "\033[38;5;141m",
	Cyan:      "\033[38;5;51m",
	White:     "\033[38;5;15m",
	Magenta:   "\033[38;5;201m",
	Brown:     "\033[38;5;94m",
	Reset:     "\033[0m",
	CadetBlue: "\033[38;5;74m",
	Plum:      "\033[38;5;96m",
	LightPink: "\033[38;5;204m",
}
