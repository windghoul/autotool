package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func operateFile(fileName, timecontent, head string) bool {
	fi, err := os.Open(fileName) // 打开文件
	defer fi.Close()             //关闭文件
	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "operate  error",
		}).Error(err.Error())
		return false
	}

	// timecontent := time.Now().Format("2006-01-02T15:04:05")
	// timetitle := time.Now().Format("20060102")
	// head := "--- \ntitle: \"技术日报(" + fileName[47:57] + ")\" \ndate: " + timecontent + "\ndraft: false\n---\n"
	contents, _ := ioutil.ReadAll(fi) // 读取所有内容
	contentString := string(contents[:])

	if contentString[:3] == "---" {
		log.WithFields(log.Fields{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}).Warn("already add the title")
		return false
	}

	newcontents := head + contentString // 组装新的内容

	newfi, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0600)
	defer newfi.Close()
	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "operate  error",
		}).Error(err.Error())
		return false
	}
	num, err := newfi.WriteString(newcontents) // 写入文件
	if err != nil || num < 1 {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "operate  error",
		}).Error(err.Error())
		return false
	}
	return true
}

func git(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info("success connect")
	pull := "cd /root/website/content/post/news.caas.one && git pull && git add . && git commit -m \"add title\" && git push"
	cmd := exec.Command("sh", "-c", pull)
	var e bytes.Buffer
	cmd.Stderr = &e

	out, err := cmd.Output()

	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "pull  error",
		}).Error(e.String())
		return
	}
	s := string(out[:])
	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info(s)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	gitJSON := GitJSON{}
	gitJSON.UnmarshalJSON([]byte(body))
	timecontent := gitJSON.Headcommit.Timestamp
	fileAddList := gitJSON.Headcommit.Added
	var flag bool
	for _, fileLocate := range fileAddList {
		IsImage := strings.Split(fileLocate, "/")[1]
		if IsImage == "images" {
			log.WithFields(log.Fields{
				"time": time.Now().Format("2006-01-02 15:04:05"),
			}).Warn("images folder")
			continue
		}
		categories := strings.Split(fileLocate, "/")[0]
		if categories[0] == '.' {
			log.WithFields(log.Fields{
				"time": time.Now().Format("2006-01-02 15:04:05"),
			}).Warn("Hidden folder")
			continue
		}
		fileName := strings.Split(fileLocate, "/")[1]
		if fileName[0] == '.' {
			log.WithFields(log.Fields{
				"time": time.Now().Format("2006-01-02 15:04:05"),
			}).Warn("Hidden folder")
			continue
		}
		switch categories {
		case "posts":
			head := "--- \ntitle: \"技术日报(" + fileName[:10] + ")\" \ndate: " + timecontent + "\ncategories: [ \"daily\"]\ndraft: false\n---\n"
			flag = operateFile("/root/website/content/post/news.caas.one/posts/"+fileName, timecontent, head)
		case "translation":
			head := "--- \ntitle: \"" + fileName[:len(fileName)-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
			flag = operateFile("/root/website/content/post/news.caas.one/translation/"+fileName, timecontent, head)
		default:
			return
		}
	}
	if flag == false {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "operate  error",
		}).Error("operate  error")
		return
	}

	hugoRun := "cd /root/website &&  hugo  "
	cmd = exec.Command("/bin/bash", "-c", hugoRun)
	cmd.Stderr = &e

	out, err = cmd.Output()

	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "hugo  error",
		}).Error(e.String())
		return
	}
	s = string(out[:])
	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info(s)

	push := "cd /root/website/public && git add . && git commit -m \" up load time @" + timecontent + "\"&& git push "

	cmd = exec.Command("/bin/bash", "-c", push)
	cmd.Stderr = &e

	out, err = cmd.Output()

	if err != nil {
		log.WithFields(log.Fields{
			"time":        time.Now().Format("2006-01-02 15:04:05"),
			"error style": "push  error",
		}).Error(e.String())
		return
	}
	s = string(out[:])
	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info(s)

	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info("end process")
}

func main() {
	router := httprouter.New()
	router.POST("/git", git)
	router.POST("/fileoperate", Fileoperate)
	log.WithFields(log.Fields{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}).Info("server running......... ")
	log.Fatal(http.ListenAndServe(":8080", router))
}
