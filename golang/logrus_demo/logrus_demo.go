package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

// refer to https://blog.csdn.net/sinat_38068807/article/details/106941878

func main() {
	LOG := log.New()
	LOG.SetOutput(os.Stdout)
	LOG.SetFormatter(&nested.Formatter{HideKeys: true, FieldsOrder: []string{"component", "category"}, NoFieldsColors: true})
	LOG.SetLevel(log.ErrorLevel)
	LOG.AddHook(lineHook{Field: "src", Skip: 0})
	LOG.AddHook(newLfsHook(2))
	i := 0
	for {
		LOG.Panicf("[%d] This is %s", i, "Debug")
		LOG.Infof("infoma\ntion\n")
		i++
		time.Sleep(time.Millisecond * 500)
		if i > 200 {
			break
		}
	}
}

// line number hook for log the call context,
type lineHook struct {
	Field string
	// skip为遍历调用栈开始的索引位置
	Skip   int
	levels []log.Level
}

// Levels implement levels
func (hook lineHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire implement fire
func (hook lineHook) Fire(entry *log.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip)
	return nil
}

func findCaller(skip int) string {
	file := ""
	line := 0
	var pc uintptr
	// 遍历调用栈的最大索引为第11层.
	for i := 0; i < 11; i++ {
		file, line, pc = getCaller(skip + i)
		// 过滤掉所有logrus包，即可得到生成代码信息
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}

	fullFnName := runtime.FuncForPC(pc)

	fnName := ""
	if fullFnName != nil {
		fnNameStr := fullFnName.Name()
		// 取得函数名
		parts := strings.Split(fnNameStr, ".")
		fnName = parts[len(parts)-1]
	}

	return fmt.Sprintf("%s:%d:%s()", file, line, fnName)
}

func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}
	n := 0

	// 获取包名
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line, pc
}

func newLfsHook(maxRemainCnt uint) log.Hook {
	writer, err := rotatelogs.New(
		"logName.%Y%m%d",
		rotatelogs.WithLinkName("logName"),
		rotatelogs.WithRotationSize(1024),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &nested.Formatter{HideKeys: false, NoColors: true})

	return lfsHook
}
