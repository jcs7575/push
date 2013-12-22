package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

var dburl, pushurl, batchpushurl, filepath string

func init() {
	dburl = beego.AppConfig.String("mysqlurl")
	pushurl = beego.AppConfig.String("pushurl")
	batchpushurl = beego.AppConfig.String("batchpushurl")
	filepath = beego.AppConfig.String("filepath")
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	var r *http.Request = this.Ctx.Request
	r.ParseForm()
	this.TplNames = "home.tpl"
}

func (this *IndexController) Post() {

	var r *http.Request = this.Ctx.Request
	this.TplNames = "success.tpl"
	var title string = r.FormValue("title")
	var pushType string = r.FormValue("pushType")
	var subTitle string = r.FormValue("subTitle")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("fileInput")
	fmt.Println(handler.Filename, handler.Header)
	if err != nil {
		fmt.Println(err)
		this.TplNames = "home.tpl"
		this.Data["fileInputErr"] = "请选择文件"
		return
	}
	newFileName := fmt.Sprint(time.Now().Unix())
	f, err := os.OpenFile(filepath+newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		this.Data["errors"] = "出错鸟"
		return
	}

	defer f.Close()
	io.Copy(f, file)

	var record PushRecord
	record.Title = title
	record.SubTitle = subTitle
	record.FilePath = filepath + newFileName
	record.ContentUrl = r.FormValue("url")
	record.PushType = pushType
	record.DisType = r.FormValue("disType")
	record.PushDateStr = r.FormValue("pushDate")
	record.Channel = r.FormValue("channel")

	pushTimeValue, err := time.ParseInLocation("2006-01-02 15", r.FormValue("pushDate"), time.UTC)
	if err != nil {
		fmt.Println(r.FormValue("pushDate"))
		this.Data["errors"] = "时间格式不对"
		return
	}
	record.PushDate = pushTimeValue
	data := ""
	pushResult := sendFileToPush(record.FilePath, record.Channel, data, pushTimeValue, pushTimeValue.Add(3*24*time.Hour))
	record.Comments = pushResult

	err = saveRecord(record)
	if err != nil {
		fmt.Println(err)
		this.Data["errors"] = "出错鸟"
		return
	}
	defer file.Close()
	this.Data["result"] = "成功鸟"
}

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	this.TplNames = "create.tpl"
}

func (this *TestController) Post() {
	var r *http.Request = this.Ctx.Request

	var pushType string = r.FormValue("pushType")
	var id string = r.FormValue("id")
	var title string = r.FormValue("title")
	var subTitle string = r.FormValue("subTitle")
	var disType string = r.FormValue("disType")
	var url string = r.FormValue("url")

	rs := sendPush(id, "md", url, disType, pushType, title, subTitle)
	if rs != "{\"data\":\"\",\"error\":0}" {
		this.Data["result"] = rs
	} else {
		this.Data["result"] = "测试成功"
	}

	this.TplNames = "success.tpl"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func saveRecord(record PushRecord) error {
	db, err := sql.Open("mysql", dburl)

	if err != nil {
		return err
	}
	defer db.Close()
	//插入数据
	stmt, errs := db.Prepare("INSERT INTO PushRecord(creation,modification,pushDate,filePath,title,subTitle,contentUrl,pushType,disType,channel,comments) values(?,?,?,?,?,?,?,?,?,?,?)")
	if errs != nil {
		return errs
	}

	res, errres := stmt.Exec(time.Now(), time.Now(), record.PushDate, record.FilePath, record.Title, record.SubTitle, record.ContentUrl, record.PushType, record.DisType, record.Channel, record.Comments)
	if errres != nil {
		return errres
	}

	id, errid := res.LastInsertId()
	if errid != nil {
		return errid
	}

	fmt.Println(id)
	return nil
}

func sendPush(id string, channel string, content string, disType string, pushType string, title string, subTitle string) string {

	var udid, uid string = "", ""
	if pushType == "udid" {
		udid = id
	} else {
		uid = id
	}
	data := applyParams(id, channel, content, disType, pushType, title, subTitle)
	fmt.Println(data, udid, uid, channel)
	response, err := http.PostForm(pushurl, url.Values{"data": {data}, "did": {udid}, "uid": {uid}, "duration": {"21600"}, "destType": {"PHOENIX"}, "type": {"COMMON"}, "channel": {channel}})
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body)
}

func applyParams(id string, channel string, content string, disType string, pushType string, title string, subTitle string) string {

	var udid string = ""
	var uid string = ""
	if pushType == "udid" {
		udid = id
	} else {
		uid = id
	}
	var intent string = ""
	var page string = ""
	if disType == "url" {
		page = content
	} else {
		intent = content
	}
	data := "[{\"srcType\":0,\"srcChannel\":\"" + channel + "\",\"destType\":\"2\",\"destDid\":\"" + udid + "\",\"destUid\":\"" + uid + "\",\"msgType\":1,\"notification\":{\"type\":\"notify\",\"title\":\"" + title + "\",\"desc\":\"" + subTitle + "\",\"icon\":\"\",\"redirect\":{\"url\":\"" + page + "\",\"intent\":{\"component\":\"" + intent + "\",\"extras\":\"\"}}},\"redirect\":\"\",\"callback\":\"\",\"extra\":\"\"}]"
	return data
}

type QueryController struct {
	beego.Controller
}

func (this *QueryController) Get() {
	this.TplNames = "query.tpl"
	list := queryPushRecord()
	this.Data["list"] = list
}

type PushRecord struct {
	Id          int32
	Title       string
	SubTitle    string
	ContentUrl  string
	PushType    string
	DisType     string
	Creation    string
	PushDateStr string
	PushDate    time.Time
	Status      int8
	Comments    string
	FilePath    string
	Channel     string
}

func queryPushRecord() []PushRecord {

	db, err := sql.Open("mysql", dburl)
	var list []PushRecord
	if err != nil {
		return list
	}
	defer db.Close()

	//查询数据
	rows, errrow := db.Query("SELECT id,ifnull(title,''),ifnull(subTitle,''),ifnull(contentUrl,''),ifnull(pushType,''),ifnull(disType,''),ifnull(creation,''),ifnull(status,''),ifnull(pushDate,''),ifnull(comments,''),ifnull(filePath,'') FROM PushRecord")
	if errrow != nil {
		return list
	}

	for rows.Next() {
		var record PushRecord
		err = rows.Scan(&record.Id, &record.Title, &record.SubTitle, &record.ContentUrl, &record.PushType, &record.DisType, &record.Creation, &record.Status, &record.PushDateStr, &record.Comments, &record.FilePath)

		checkErr(err)
		list = append(list, record)
	}
	return list
}

func sendFileToPush(filePath string, channel string, data string, startDate time.Time, endDate time.Time) string {
	fmt.Println(batchpushurl)
	v := url.Values{}
	v.Set("data", data)
	v.Encode()
	encode_data := strings.Replace(v.Get("data"), "data=", "", -1)
	fmt.Println(encode_data)

	cmd := exec.Command("curl", "-F", "dids=@{"+filePath+"}", batchpushurl+"?channel=${"+channel+"}&data=${"+encode_data+"}&destType=PHOENIX&startDate="+convertDate(startDate)+"&endDate="+convertDate(endDate))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "error"
	}
	return strings.Replace(out.String(), "\n", "", -1)
}
func convertDate(oriDate time.Time) string {
	s := oriDate.String()
	return s[0:4] + s[5:7] + s[8:10] + s[11:13] + "0000"
}
