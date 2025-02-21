package taillog

import (
	"fmt"
	"testing"
	"time"

	"github.com/hpcloud/tail"
	"github.com/stretchr/testify/assert"
)

func TestTailLog(t *testing.T) {
	fileName := "my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true}
	tails, err := tail.TailFile(fileName, config)
	assert.NoError(t, err, "tail file failed")

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
