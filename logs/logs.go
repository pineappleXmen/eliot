package logs

type LogEntry struct {
	index   int64
	command interface{}
	term    int64
}

type Logs struct {
	logs       []LogEntry
	startIndex int64
}

func (l *Logs) init(withDummy bool) {
	// init logs
	l.logs = make([]LogEntry, 0)
	if withDummy {
		l.logs = append(l.logs, LogEntry{
			index:   0,
			command: nil,
			term:    0,
		})
	}
}

func (l *Logs) append(entry LogEntry) {
	l.logs = append(l.logs, entry)
}

func (l *Logs) get(index int64) LogEntry {
	return l.logs[index-l.startIndex]
}

func (l *Logs) getLast() LogEntry {
	return l.logs[len(l.logs)-1]
}

// cut off entries after input index
func (l *Logs) cutEnd(index int64) {
	l.logs = l.logs[:index-l.startIndex]
}

// cut off entries before input index
func (l *Logs) cutStart(index int64) {
	l.logs = l.logs[index-l.startIndex:]
	l.startIndex = index
}

// get start index
func (l *Logs) getStartIndex() int64 {
	return l.startIndex
}

// get last index
func (l *Logs) getLastIndex() int64 {
	return l.startIndex + int64(len(l.logs))
}

// get length
func (l *Logs) getLength() int64 {
	return int64(len(l.logs))
}
