// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import ()

type Project struct {
	ID      int32
	Name    string
	Cpu     int64
	Memory  int64
	Storage int64
	Cluster string
}
