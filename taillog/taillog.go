package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var tailObj *tail.Tail

func Init(filename string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 文件滚动时重新打开文件
		Follow:    true,                                 // 跟踪文件的新行
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
		MustExist: false,                                // 文件不存在时不报错
		Poll:      true,                                 // 使用轮询来监视文件更改
	}

	// 打开并追踪指定的日志文件
	tailObj, err = tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err) // 打印错误信息
		return
	}

	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
