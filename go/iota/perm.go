package main

import (
	"fmt"
	"strings"
)

type FilePerm uint8

const (
	ReadPerm FilePerm = 1 << iota
	WritePerm
	ExecPerm
	MaxPerm = ExecPerm | WritePerm | ReadPerm
)

func (p FilePerm) String() string {
	if p > MaxPerm {
		return fmt.Sprintf("unknown FilePerm - %d", p)
	}

	switch p {
	case ReadPerm:
		return "read"
	case WritePerm:
		return "write"
	case ExecPerm:
		return "exec"
	}

	// A combination
	var perms []string
	for _, bit := range []FilePerm{ReadPerm, WritePerm, ExecPerm} {
		if p&bit != 0 {
			perms = append(perms, bit.String())
		}
	}

	return strings.Join(perms, "|")
}

func main() {
	p := ReadPerm | WritePerm
	fmt.Printf("p = %s (%04b)\n", p, p)
}
