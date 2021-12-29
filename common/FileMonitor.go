package common

import (
	"context"
	"fmt"
	"gonet/actor"
	"gonet/server/rpc"
	"os"
	"time"
	"unsafe"
)

type (
	FileRead func() //reload
	FileInfo struct {
		Info os.FileInfo
		Call FileRead //call reload
	}

	FileMonitor struct {
		actor.Actor
		m_FilesMap map[string]*FileInfo
	}

	IFileMonitor interface {
		actor.IActor
		addFile(string, FileRead)
		delFile(string)
		update()
		AddFile(string, FileRead)
	}
)

func (this *FileMonitor) Init() {
	this.Actor.Init()
	this.m_FilesMap = map[string]*FileInfo{}
	this.RegisterTimer(3*time.Second, this.update)
	this.RegisterCall("addfile", func(ctx context.Context, fileName string, p *int64) {
		pFunc := (*FileRead)(unsafe.Pointer(p))
		this.addFile(fileName, *pFunc)
	})

	this.RegisterCall("delfile", func(ctx context.Context, fileName string) {
		this.delFile(fileName)
	})
	this.Actor.Start()
}

func (this *FileMonitor) AddFile(fileName string, pFunc FileRead) {
	ponit := unsafe.Pointer(&pFunc)
	this.SendMsg(rpc.RpcHead{}, "addfile", fileName, (*int64)(ponit))
}

func (this *FileMonitor) addFile(fileName string, pFunc FileRead) {
	file, err := os.Open(fileName)
	if err == nil {
		fileInfo, err := file.Stat()
		if err == nil {
			this.m_FilesMap[fileName] = &FileInfo{fileInfo, pFunc}
		}
	}
}

func (this *FileMonitor) delFile(fileName string) {
	delete(this.m_FilesMap, fileName)
}

func (this *FileMonitor) update() {
	for i, v := range this.m_FilesMap {
		file, err := os.Open(i)
		if err == nil {
			fileInfo, err := file.Stat()
			if err == nil && v.Info.ModTime() != fileInfo.ModTime() {
				v.Call()
				v.Info = fileInfo
				fmt.Println(fmt.Sprintf("file [%s] reload", v.Info.Name()))
			}
		}
	}
}
