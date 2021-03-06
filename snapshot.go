package raft

import "io"

type Snapshot struct {
	LastTerm  uint64 `json:"lastIndex"`
	LastIndex uint64 `josn:"lastTerm"`

	config configuration `json:"configuration"`
	State  []byte        `json:"state"`
	Path   string        `json:"path"`
}

// 服务器从 Snapshot 状态启动
type RequestSnapshotRecovery struct {
	LeaderName string
	LastIndex  uint64
	LastTerm   uint64
	Config     configuration
	State      []byte
}

//Snapshot response
type SnapshotRecoveryResponse struct {
	Term        uint64
	Success     bool
	CommitIndex int
}

type RequestSnapshot struct {
	LeaderName string
	LastIndex  uint64
	LastTerm   uint64
}

type SnapshotResponse struct {
}

func (ss *Snapshot) save() error {
	//todo : 1、打开文件 2、序列化protuuf  3、写入文件

	var err error

	return err

}

func (ss *Snapshot) remove() error {
	// todo : 1、将此文件移除
	var err error

	return err
}

func (req *RequestSnapshot) Encode(w io.Writer) (int, error) {
	// todo

	return 0, nil
}

func (req *RequestSnapshot) Decode(w io.Writer) (int, error) {
	//todo

	return 0, nil
}

func (req *SnapshotResponse) Encode(w io.Writer) (int, error) {
	// todo

	return 0, nil
}

func (req *SnapshotResponse) Decode(w io.Writer) (int, error) {
	//todo

	return 0, nil
}

func (req *RequestSnapshotRecovery) Encode(w io.Writer) (int, error) {
	// todo

	return 0, nil
}

func (req *RequestSnapshotRecovery) Decode(w io.Writer) (int, error) {
	//todo

	return 0, nil
}

func (req *SnapshotRecoveryResponse) Encode(w io.Writer) (int, error) {
	// todo

	return 0, nil
}

func (req *SnapshotRecoveryResponse) Decode(w io.Writer) (int, error) {
	//todo

	return 0, nil
}

func newSnapshotRecoryRequest(success bool) *RequestSnapshotRecovery {
	return &RequestSnapshotRecovery{}
}

func newSnapshotRecoryResponse(term uint64, success bool, commitIndex uint64) SnapshotRecoveryResponse {
	return SnapshotRecoveryResponse{Success: success}
}

func newRequestSnapshot(success bool) *RequestSnapshot {
	return &RequestSnapshot{}
}

func newSnapshotResponse(success bool) *SnapshotResponse {
	return &SnapshotResponse{}
}
