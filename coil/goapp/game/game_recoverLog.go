package game

type recoverLog struct {
	pos
	org int
}

type recoverLogArray struct {
	logs   []recoverLog
	length int
}

func makeLogAraay(size int) recoverLogArray {
	var r recoverLogArray
	r.logs = make([]recoverLog, size)
	return r
}

func (r *recoverLogArray) append(data recoverLog) {
	r.logs[r.length] = data
	r.length++
}

func (r *recoverLogArray) getValue(index int) recoverLog {
	return r.logs[index]
}
