package logs

import (
	"fmt"
	"strconv"
)

type Entry struct {
	command interface{}
	term    int64
}

type Log struct {
	logs       []Entry
	startIndex int64
}

func (l *Log) init(withDummy bool) {
	// init logs
	l.logs = make([]Entry, 0)
	if withDummy {
		l.logs = append(l.logs, Entry{
			command: nil,
			term:    0,
		})
	}
}

func (l *Log) append(entry Entry) {
	l.logs = append(l.logs, entry)
}

func (l *Log) get(index int64) Entry {
	return l.logs[index-l.startIndex]
}

func (l *Log) getLast() Entry {
	return l.logs[len(l.logs)-1]
}

// cut off entries after input index
func (l *Log) cutEnd(index int64) {
	l.logs = l.logs[:index-l.startIndex]
}

// cut off entries before input index
func (l *Log) cutStart(index int64) {
	l.logs = l.logs[index-l.startIndex:]
	l.startIndex = index
}

// get start index
func (l *Log) getStartIndex() int64 {
	return l.startIndex
}

// get last index
func (l *Log) getLastIndex() int64 {
	return l.startIndex + int64(len(l.logs))
}

// get length
func (l *Log) getLength() int64 {
	return int64(len(l.logs))
}

func (l *Log) String() string {
	reply := ""
	for i := int64(0); i < int64(len(l.logs)); i++ {
		reply += "("
		reply += strconv.Itoa(int(i + l.startIndex))
		reply += ")"
		reply += l.logs[i].String()
		reply += " "
	}
	return reply
}

func (e *Entry) String() string {
	return fmt.Sprintf("[%v:%v]", e.term, e.command)
}
