package task21

// POSIX THREAD

type PosixThread struct {
}

func (thread *PosixThread) pThreadCreate() {
	// Create thread
}

func (thread *PosixThread) pThreadJoin() {
	// Join thread
}

func (thread *PosixThread) pThreadDetach() {
	// Detach thread
}

// WINDOWS THREAD

type WindowsThread struct {
}

func (thread *WindowsThread) CreateThread() {
	// Detach thread
}

func (thread *WindowsThread) WaitForSingleObject() {
	// Wait for single object
}

func (thread *WindowsThread) CloseHandle() {
	// Close handle
}

// Shared thread interface

type Thread interface {
	Create()
	Join()
	Detach()
}

// Implementation of the Thread interface for POSIX thread

type ThreadPosixImpl struct {
	thread PosixThread
}

func (t *ThreadPosixImpl) Create() {
	t.thread.pThreadCreate()
}

func (t *ThreadPosixImpl) Join() {
	t.thread.pThreadJoin()
}

func (t *ThreadPosixImpl) Detach() {
	t.thread.pThreadDetach()
}

// Implementation of the Thread interface for Windows thread

type ThreadWindowsImpl struct {
	thread WindowsThread
}

func (t *ThreadWindowsImpl) Create() {
	t.thread.CreateThread()
}

func (t *ThreadWindowsImpl) Join() {
	t.thread.WaitForSingleObject()
	t.thread.CloseHandle()
}

func (t *ThreadWindowsImpl) Detach() {
	t.thread.CloseHandle()
}

// Some function

func NewThread() Thread {
	const THREAD = "WINDOWS" // Can be changed according to current OS
	if THREAD == "WINDOWS" {
		return &ThreadWindowsImpl{WindowsThread{}}
	} else if THREAD == "POSIX" {
		return &ThreadPosixImpl{PosixThread{}}
	}

	return nil
}

func Task21() {
	thread := NewThread()
	thread.Join()
}
