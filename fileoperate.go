package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

var categories string

func rewriteFile(file, timecontent, head string) bool {
	fi, err := os.Open(file) // 打开文件
	defer fi.Close()         //关闭文件
	if err != nil {
		fmt.Println(err.Error())
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

func showtime() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
}

func gitPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "123")
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)

	fmt.Println(raw["commits"].([]interface{})[0].(map[string]interface{})["timestamp"])
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
					timecontent := file.Name()[:10] + "T15:03:04+08:00"
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

func main1() {
	folder := "/root/website/content/post/news.caas.one"
	listFile(folder)
}

// func main() {

// 	ls := " ls -t | head -1"
// 	cmd := exec.Command("/bin/bash", "-c", ls)

// 	var e bytes.Buffer
// 	cmd.Stderr = &e

// 	out, err := cmd.Output()

// 	if err != nil {
// 		fmt.Println(e.String())
// 		fmt.Println("ls error")
// 		return
// 	}
// 	n := len(out)
// 	s := string(out[:n-1])
// 	rewriteFile(s)
// }

// func main() {
// 	rewriteFile("2018-09-24-new.md")
// }

// func main() {
// 	fmt.Println(len("/Users/yp-tc-m-5063/website/content/post/news.caas.one/posts/"))
// }
// func main() {
// 	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
// }
// func main() {

// 	s := "123123/123123"
// 	fmt.Printf("%q\n", strings.Split(s, "/"))
// }
// ls := "cd /root/website/content/post/news.caas.one/posts && ls -t | head -1"
// cmd = exec.Command("/bin/bash", "-c", ls)

// cmd.Stderr = &e

// out, err = cmd.Output()

// if err != nil {
// 	log.WithFields(log.Fields{
// 		"time":        time.Now().Format("2006-01-02 15:04:05"),
// 		"error style": "ls  error",
// 	}).Error(e.String())
// 	return
// }

// n := len(out)
// s = string(out[:n-1])
// body, err := ioutil.ReadAll(r.Body)
// if err != nil {
// 	return
// }
// var raw map[string]interface{}
// json.Unmarshal(body, &raw)
// timecontent := raw["commits"].([]interface{})[0].(map[string]interface{})["timestamp"].(string)
// fileLocate := raw["commits"].([]interface{})[0].(map[string]interface{})["added"].([]interface{})[0].(string)
// categories := strings.Split(fileLocate, "/")[0]
// fileName := strings.Split(fileLocate, "/")[1]
// var flag bool
// switch categories {
// case "posts":
// 	head := "--- \ntitle: \"技术日报(" + fileName[:10] + ")\" \ndate: " + timecontent + "\ncategories: [ \"daily\"]\ndraft: false\n---\n"
// 	flag = operateFile("/root/website/content/post/news.caas.one/posts/"+fileName, timecontent, head)
// 	// fmt.Println(folder + "/" + file.Name())
// 	// timecontent := "2018-10-15T15:03:04+08:00"
// 	// rewriteFile(folder+"/"+file.Name(), timecontent, head)
// case "translation":
// 	head := "--- \ntitle: \"" + fileName[:len(s)-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
// 	flag = operateFile("/root/website/content/post/news.caas.one/translation/"+fileName, timecontent, head)
// 	// fmt.Println(folder + "/" + file.Name())
// 	// timecontent := "2018-10-15T15:03:04+08:00"
// 	// head := "--- \ntitle: \"" + file.Name()[47:len(file.Name())-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
// 	// rewriteFile(folder+"/"+file.Name(), timecontent, head)
// default:
// 	return
// }

// func testJSON(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return
// 	}
// 	var raw map[string]interface{}
// 	json.Unmarshal(body, &raw)
// 	timecontent := raw["commits"].([]interface{})[0].(map[string]interface{})["timestamp"].(string)
// 	fileAddList := raw["commits"].([]interface{})[0].(map[string]interface{})["added"].([]interface{})
// 	// var flag bool
// 	for _, fileLocate := range fileAddList {
// 		IsImage := strings.Split(fileLocate.(string), "/")[1]
// 		if IsImage == "images" {
// 			return
// 		}
// 		categories := strings.Split(fileLocate.(string), "/")[0]
// 		if categories[0] == '.' {
// 			continue
// 		}
// 		fileName := strings.Split(fileLocate.(string), "/")[1]
// 		if fileName[0] == '.' {
// 			continue
// 		}
// 		switch categories {
// 		case "posts":
// 			head := "--- \ntitle: \"技术日报(" + fileName[:10] + ")\" \ndate: " + timecontent + "\ncategories: [ \"daily\"]\ndraft: false\n---\n"
// 			fmt.Println(head)
// 			// flag = operateFile("/root/website/content/post/news.caas.one/posts/"+fileName, timecontent, head)
// 		case "translation":
// 			head := "--- \ntitle: \"" + fileName[:len(fileName)-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
// 			fmt.Println(head)
// 			// flag = operateFile("/root/website/content/post/news.caas.one/translation/"+fileName, timecontent, head)
// 		default:
// 			return
// 		}
// 	}
// 	// if flag == false {
// 	// 	log.WithFields(log.Fields{
// 	// 		"time":        time.Now().Format("2006-01-02 15:04:05"),
// 	// 		"error style": "operate  error",
// 	// 	}).Error("operate  error")
// 	// 	return
// 	// }
// }

// func testJSON(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return
// 	}
// 	var raw map[string]interface{}
// 	json.Unmarshal(body, &raw)
// 	timecontent := raw["commits"].([]interface{})[0].(map[string]interface{})["timestamp"].(string)
// 	fileAddList := raw["commits"].([]interface{})[0].(map[string]interface{})["added"].([]interface{})
// 	// var flag bool
// 	for _, fileLocate := range fileAddList {
// 		IsImage := strings.Split(fileLocate.(string), "/")[1]
// 		if IsImage == "images" {
// 			log.WithFields(log.Fields{
// 				"time": time.Now().Format("2006-01-02 15:04:05"),
// 			}).Warn("images folder")
// 			continue
// 		}
// 		categories := strings.Split(fileLocate.(string), "/")[0]
// 		if categories[0] == '.' {
// 			log.WithFields(log.Fields{
// 				"time": time.Now().Format("2006-01-02 15:04:05"),
// 			}).Warn("Hidden folder")
// 			continue
// 		}
// 		fileName := strings.Split(fileLocate.(string), "/")[1]
// 		if fileName[0] == '.' {
// 			log.WithFields(log.Fields{
// 				"time": time.Now().Format("2006-01-02 15:04:05"),
// 			}).Warn("Hidden folder")
// 			continue
// 		}
// 		switch categories {
// 		case "posts":
// 			head := "--- \ntitle: \"技术日报(" + fileName[:10] + ")\" \ndate: " + timecontent + "\ncategories: [ \"daily\"]\ndraft: false\n---\n"
// 			fmt.Println(head)
// 			// flag = operateFile("/root/website/content/post/news.caas.one/posts/"+fileName, timecontent, head)
// 		case "translation":
// 			head := "--- \ntitle: \"" + fileName[:len(fileName)-3] + "\" \ndate: " + timecontent + "\ncategories: [ \"translation\"]\ndraft: false\n---\n"
// 			fmt.Println(head)
// 			// flag = operateFile("/root/website/content/post/news.caas.one/translation/"+fileName, timecontent, head)
// 		default:
// 			return
// 		}
// 	}
// 	// if flag == false {
// 	// 	log.WithFields(log.Fields{
// 	// 		"time":        time.Now().Format("2006-01-02 15:04:05"),
// 	// 		"error style": "operate  error",
// 	// 	}).Error("operate  error")
// 	// 	return
// 	// }
// }
