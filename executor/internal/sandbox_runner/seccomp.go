package sandbox_runner

import (
	"github.com/criyle/go-sandbox/pkg/seccomp"
	"os"
	"syscall"
	"unsafe"
)

const AllowedFilePerm = os.FileMode(0644)

// compiled seccomp policy from following kafel (https://github.com/google/kafel) expression.
// allows read, write, exit
// open must only create files with 644 permission
// execve path must be given pointer (placeholder 0xFFFFFFFFEEEEEEEE)
/*
#define O_CREAT 0x40

DEFAULT KILL

ALLOW{
  read, write, exit,
  open(path, flags, mode) {
    (flags & O_CREAT) == 0 || mode == 0644
  },

  execve(path, argv, env) {
    path == 0xFFFFFFFFEEEEEEEE
  }
}
*/
// To disassemble use:
// disassemble, _ := bpf.Disassemble([]bpf.RawInstruction{{0x20, 0, 0, 0x4}, ....});
// for i, instruction := range disassemble {
// 	   fmt.Printf("%d: %v\n", i, instruction)
// }

func GetSeccompPolicy(argv0 *byte) *syscall.SockFprog {
	pointerHigh := uint32(uintptr(unsafe.Pointer(argv0)) >> 32)
	pointerLow := uint32(uintptr(unsafe.Pointer(argv0)))

	var seccompInstrs seccomp.Filter = []syscall.SockFilter{
		{0x20, 0, 0, 0x4},
		{0x15, 0, 0x10, 0xc000003e},
		{0x20, 0, 0, 0},
		{0x35, 0x1, 0, 0x3},
		{0x35, 0x2, 0xe, 0x2},
		{0x35, 0x7, 0, 0x3c},
		{0x35, 0x7, 0xb, 0x3b},
		{0x20, 0, 0, 0x18},
		{0x45, 0, 0xa, 0x40},
		{0x20, 0, 0, 0x24},
		{0x15, 0, 0x7, 0},
		{0x20, 0, 0, 0x20},
		{0x15, 0x6, 0x5, 0x1a4},
		{0x35, 0x4, 0x5, 0x3d},
		{0x20, 0, 0, 0x14},
		{0x15, 0, 0x2, pointerHigh},
		{0x20, 0, 0, 0x10},
		{0x15, 0x1, 0, pointerLow},
		{0x6, 0, 0, 0},
		{0x6, 0, 0, 0x7fff0000},
	}

	return seccompInstrs.SockFprog()
}
