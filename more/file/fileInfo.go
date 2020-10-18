package main

import (
	"fmt"
	"os"
)

/*
type FileInfo interface {
	Name() string       //   文件名.扩展名
	Size() int64        //   文件大小 字节数
	Mode() FileMode     // 	文件权限 -(type -:file d:createDir |:link )rwx(owner)---(group)rw-(others)
	ModTime() time.Time // 修改时间
	IsDir() bool        // 是否文件夹
	Sys() interface{}   // underlying data source (can return nil) 基础数据源接口
}
*/

/*
文件权限 mode
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky
	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777 // Unix permission bits
)
*/

//type File struct {
//	*file // os specific
//}

//type file struct {
//	pfd         poll.FD
//	name        string
//	dirinfo     *dirInfo // nil unless directory being read
//	nonblock    bool     // whether we set nonblocking mode
//	stdoutOrErr bool     // whether this is stdout or stderr
//	appendMode  bool     // whether file is opened for appending
//}

//type FD struct {
//	// Lock sysfd and serialize access to Read and Write methods.
//	fdmu fdMutex
//
//	// System file descriptor. Immutable until Close.
//	Sysfd int
//
//	// I/O poller.
//	pd pollDesc
//
//	// Writev cache.
//	iovecs *[]syscall.Iovec
//
//	// Semaphore signaled when file is closed.
//	csema uint32
//
//	// Non-zero if this file has been set to blocking mode.
//	isBlocking uint32
//
//	// Whether this is a streaming descriptor, as opposed to a
//	// packet-based descriptor like a UDP socket. Immutable.
//	IsStream bool
//
//	// Whether a zero byte read indicates EOF. This is false for a
//	// message based socket connection.
//	ZeroReadIsEOF bool
//
//	// Whether this is a file rather than a network socket.
//	isFile bool
//}

func main() {
	/*
		FileInfo : 文件信息
			interface
				Name() string          文件名.扩展名
				Size() int64           文件大小 字节数
				Mode() FileMode        文件权限 -(type -:file d:createDir |:link )rwx(owner)---(group)rw-(others)
				ModTime() time.Time    修改时间
				IsDir() bool           是否文件夹
				Sys() interface{}      underlying data source (can return nil) 基础数据源接口
	*/
	fileInfo, err := os.Stat("fib.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("\n", fileInfo.Name(), "\n",
		fileInfo.Size(), "\n",
		fileInfo.Mode(), "\n",
		fileInfo.ModTime(), "\n",
		fileInfo.IsDir(), "\n",
		fileInfo.Sys())
}
