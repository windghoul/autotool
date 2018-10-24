package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

var categories string

func rewriteFile(file, timecontent, head string) bool {
	fi, err := os.Open(file) // 打开文件
	defer fi.Close()         //关闭文件
	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "operate  error",
		}).Error(err.Error())
		return false
	}
	// br := bufio.NewReader(fi)
	// a, _, _ := br.ReadLine()

	// fmt.Println(a)

	// timecontent := time.Now().Format("2006-01-02T15:04:05")
	// timecontent := file[61:71] + "T15:04:05"
	// timetitle := time.Now().Format("20060102")

	contents, _ := ioutil.ReadAll(fi) // 读取所有内容
	contentString := string(contents[:])
	if contentString[:3] == "---" {
		fmt.Println("already add the title")
		return false
	}
	newcontents := head + contentString // 组装新的内容

	newfi, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0600)
	// newfi, err := os.OpenFile(file, os.O_RDWR, 0666) // 打开文件
	defer newfi.Close()
	if err != nil {
		return false
	}
	// newfi.Seek(0, os.SEEK_SET)
	// num, err := newfi.WriteAt([]byte(newcontents), 0) // 在开头覆盖插入内容
	num, err := newfi.WriteString(newcontents) // 写入文件
	if err != nil || num < 1 {
		return false
	}
	return true
}

func listFile(folder string) {
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	// head := "--- \ntitle: \"技术日报(" + file[61:71] + ")\" \ndate: " + timecontent + "+08:00 \ncategories: [ \"" + categories + "\"]\ndraft: false\n---\n"

	for _, file := range files {
		if file.IsDir() {
			if file.Name()[0] != '.' {
				categories = file.Name()
				// fmt.Println(file.Name())
				listFile(folder + "/" + file.Name())
				// rewriteFile(folder + "/" + file.Name())
			} else {
				continue
			}

		} else {
			if file.Name()[0] != '.' {
				// fmt.Println(categories)
				// fmt.Println(folder + "/" + file.Name())
				switch categories {
				case "posts":
					// fmt.Println(folder + "/" + file.Name())
					timecontent := file.Name()[:10] + "T00:00:00+08:00"
					head := "--- \ntitle: \"技术日报(" + file.Name()[:10] + ")\" \ndate: " + timecontent + "\ncategories: [ \"daily\"]\ndraft: false\n---\n"
					rewriteFile(folder+"/"+file.Name(), timecontent, head)
				case "translation":
					// fmt.Println(folder + "/" + file.Name())
					timecontent := "2018-10-15T15:03:04+08:00"
					head := "--- \ntitle: \"" + file.Name()[:len(file.Name())-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
					rewriteFile(folder+"/"+file.Name(), timecontent, head)
				default:
					continue
				}

			} else {
				continue
			}
		}
	}

}

func Fileoperate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	folder := "/root/website/content/post/news.caas.one"
	// folder := "/Users/yp-tc-m-5063/website/content/post/news.caas.one"
	listFile(folder)
	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info("finish jobs ")
}
