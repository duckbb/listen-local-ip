package listen_local_ip

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

var fileName string = "./ip.txt"
var Log *log.Logger
var debug = true

func Init() {
	fileName := "info.log"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open file error !")
	}
	if debug {
		Log = log.New(io.MultiWriter(os.Stderr, logFile), "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		Log = log.New(io.MultiWriter(logFile), "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func isIp(str string) (bool, error) {
	matched, err := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", str)
	return matched, err
}

//get ip and wirite in log
func Write(content string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		Log.Println("get file err:", err)
	}
	if matched, err := isIp(content); err != nil || !matched {
		return
	}
	if !prevIpEqual(content) {
		now := time.Now().Format("2006-01-02 15:04:05")
		_, err = io.WriteString(file, now+" ="+content)
		io.WriteString(file, "\n")
		prevIpEqual(content)
		if err != nil {
			Log.Println("write file err:", err)
		}
	}
	return
}

//previous ip equal current ip
func prevIpEqual(ip string) bool {
	f, err := os.Open(fileName)
	if err != nil {
		return false
	}
	buf := bufio.NewReader(f)
	linesSli := []string{}
	for {
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			return false
		}
		if err == io.EOF {
			break
		}
		linesSli = append(linesSli, line)
	}
	ipRegexp := regexp.MustCompile("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}")
	if len(linesSli) > 0 {
		prevIp := ipRegexp.FindString(linesSli[len(linesSli)-1])
		if prevIp == ip {
			return true
		}
	}
	return false
}
