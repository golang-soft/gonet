package base

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type LG_TYPE int

const (
	LG_WARN  LG_TYPE = iota
	LG_DEBUG LG_TYPE = iota
	LG_INFO  LG_TYPE = iota
	LG_ERROR LG_TYPE = iota
	LG_FATAL LG_TYPE = iota
	LG_MAX   LG_TYPE = iota
)
const (
	PATH = "log"
)

type (
	CLog struct {
		m_Logger    [LG_MAX]log.Logger
		m_pFile     [LG_MAX]*os.File
		m_Time      time.Time
		m_FileName  string
		m_LogSuffix string
		m_ErrSuffix string
		m_Loceker   sync.Mutex
	}

	ILog interface {
		Init(string) bool
		Write(LG_TYPE)
		WriteFile(LG_TYPE)
		Println(...interface{})
		Print(...interface{})
		Printf(string, ...interface{})
		Fatalln(...interface{})
		Fatal(...interface{})
		Fatalf(string, ...interface{})
	}
)

var (
	GLOG *CLog
)

func (this *CLog) Init(fileName string) bool {
	//this.m_pFile = nil
	this.m_FileName = fileName
	this.m_LogSuffix = "log"
	this.m_ErrSuffix = "err"
	GLOG = this
	return true
}

func (this *CLog) GetSuffix(nType LG_TYPE) string {
	if nType == LG_WARN {
		return this.m_LogSuffix
	} else {
		return this.m_ErrSuffix
	}
}

func (this *CLog) Write(nType LG_TYPE) {
	this.WriteFile(nType)
	tTime := time.Now()
	this.m_Logger[nType].SetPrefix(fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d] [%s]", tTime.Year(), tTime.Month(), tTime.Day(),
		tTime.Hour(), tTime.Minute(), tTime.Second(), this.m_FileName))
}

func (this *CLog) Debug(v1 ...interface{}) {
	this.Write(LG_WARN)
	params := make([]interface{}, len(v1)+2)
	params[0] = "[" + this.m_FileName + "][Debug]"
	for i, v := range v1 {
		params[i+1] = v
	}
	params[len(v1)+1] = "\r"
	this.m_Logger[LG_WARN].Output(2, "["+this.m_FileName+"][Debug]"+fmt.Sprintln(params...))
	log.Println(params...)
}

func (this *CLog) Debugf(format string, params ...interface{}) {
	this.Write(LG_WARN)
	format += "\r\n"
	this.m_Logger[LG_WARN].Output(2, fmt.Sprintf("["+this.m_FileName+"][Debug]"+format, params...))
	log.Printf("["+this.m_FileName+"][Debug]"+format, params...)
}

func (this *CLog) Errorf(format string, params ...interface{}) {
	this.Write(LG_ERROR)
	format += "\r\n"
	this.m_Logger[LG_ERROR].Output(2, fmt.Sprintf("["+this.m_FileName+"][Debug]"+format, params...))
	log.Printf("["+this.m_FileName+"][Debug]"+format, params...)
}

func (this *CLog) Infof(format string, params ...interface{}) {
	this.Write(LG_INFO)
	format += "\r\n"
	this.m_Logger[LG_INFO].Output(2, fmt.Sprintf("["+this.m_FileName+"][Debug]"+format, params...))
	log.Printf("["+this.m_FileName+"][Debug]"+format, params...)
}

func (this *CLog) Warnf(format string, params ...interface{}) {
	this.Write(LG_WARN)
	format += "\r\n"
	this.m_Logger[LG_WARN].Output(2, fmt.Sprintf("["+this.m_FileName+"][Debug]"+format, params...))
	log.Printf("["+this.m_FileName+"][Debug]"+format, params...)
}

func (this *CLog) Println(v1 ...interface{}) {
	this.Write(LG_WARN)
	params := make([]interface{}, len(v1)+1)
	for i, v := range v1 {
		params[i] = v
	}
	params[len(v1)] = "\r"
	this.m_Logger[LG_WARN].Output(2, fmt.Sprintln(params...))
	log.Println(params...)
}

func (this *CLog) Print(v1 ...interface{}) {
	this.Write(LG_WARN)
	params := make([]interface{}, len(v1)+1)
	for i, v := range v1 {
		params[i] = v
	}
	params[len(v1)] = "\r\n"
	this.m_Logger[LG_WARN].Output(2, fmt.Sprint(params...))
	log.Print(params...)
}

func (this *CLog) Printf(format string, params ...interface{}) {
	this.Write(LG_WARN)
	format += "\r\n"
	this.m_Logger[LG_WARN].Output(2, fmt.Sprintf(format, params...))
	log.Printf(format, params...)
}

func (this *CLog) Fatalln(v1 ...interface{}) {
	this.Write(LG_ERROR)
	params := make([]interface{}, len(v1)+1)
	for i, v := range v1 {
		params[i] = v
	}
	params[len(v1)] = "\r"
	this.m_Logger[LG_ERROR].Output(2, fmt.Sprintln(params...))
	log.Println(params...)
}

func (this *CLog) Fatal(v1 ...interface{}) {
	this.Write(LG_ERROR)
	params := make([]interface{}, len(v1)+1)
	for i, v := range v1 {
		params[i] = v
	}
	params[len(v1)] = "\r\n"
	this.m_Logger[LG_ERROR].Output(2, fmt.Sprint(params...))
	log.Print(params...)
}

func (this *CLog) Fatalf(format string, params ...interface{}) {
	this.Write(LG_FATAL)
	format += "\r\n"
	this.m_Logger[LG_FATAL].Output(2, fmt.Sprintf(format, params...))
	log.Printf(format, params...)
}

func (this *CLog) WriteFile(nType LG_TYPE) {
	var err error
	tTime := time.Now()
	if this.m_pFile[nType] == nil || this.m_Time.Year() != tTime.Year() ||
		this.m_Time.Month() != tTime.Month() || this.m_Time.Day() != tTime.Day() {
		this.m_Loceker.Lock()
		if this.m_pFile[nType] != nil {
			defer this.m_pFile[nType].Close()
		}

		if PathExists(PATH) == false {
			os.Mkdir(PATH, os.ModeDir)
		}

		sFileName := fmt.Sprintf("%s/%s_%d%02d%02d.%s", PATH, this.m_FileName, tTime.Year(), tTime.Month(), tTime.Day(),
			this.GetSuffix(nType))

		if PathExists(sFileName) == false {
			os.Create(sFileName)
		}

		this.m_pFile[nType], err = os.OpenFile(sFileName, os.O_RDWR|os.O_APPEND, 0)
		if err != nil {
			log.Fatalf("open logfile[%s] error", sFileName)
		}

		this.m_Logger[nType].SetOutput(this.m_pFile[nType])
		this.m_Logger[nType].SetPrefix(fmt.Sprintf("[%04d-%02d-%02d %02d:%02d:%02d]", tTime.Year(), tTime.Month(), tTime.Day(),
			tTime.Hour(), tTime.Minute(), tTime.Second()))
		this.m_Logger[nType].SetFlags(log.Llongfile)
		Stat, _ := this.m_pFile[nType].Stat()
		if Stat != nil {
			this.m_Time = Stat.ModTime()
		} else {
			this.m_Time = time.Now()
		}
		this.m_Loceker.Unlock()
	}
}
