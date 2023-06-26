package push

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/imroc/req/v3"
)

func Wxpush(corpid, corpsecret, msg, msgtype string) {
	wxconfig := struct{ Msgtype, Corpid, Corpsecret, Msg string }{msgtype,
		corpid, corpsecret, msg,
	}
	client := req.C()
	resp, err := client.R().SetQueryParams(map[string]string{
		"corpid":     wxconfig.Corpid,
		"corpsecret": wxconfig.Corpsecret}).Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken")
	if err != nil {
		fmt.Println(err)
	}
	j, _ := gjson.LoadContent(resp)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + j.Get("access_token").String()
	d, _ := gjson.LoadContent(`{
        "touser": "@all",
        "toparty": "",
        "totag": "",
        "msgtype": "",
        "agentid": "1000002",
        "enable_id_trans": 0,
        "enable_duplicate_check": 0,
        "duplicate_check_interval": 1800
    }`)
	if wxconfig.Msgtype == "textcard" {
		dd, _ := gjson.LoadContent(`{
            "title": "",
            "description": "",
            "url": "URL",
            "btntxt": "详情"
        }`)
		err := dd.Set("description", wxconfig.Msg)
		if err != nil {
			fmt.Println(err)
		}
		err = d.Set("textcard", dd.String())
		if err != nil {
			fmt.Println(err)
		}
	} else if wxconfig.Msgtype == "markdown" {
		dd, _ := gjson.LoadContent(`{"content": ""}`)
		err := dd.Set("content", wxconfig.Msg)
		if err != nil {
			fmt.Println(err)
		}
		err = d.Set("markdown", dd.String())
		if err != nil {
			fmt.Println(err)
		}
	} else {
		dd, _ := gjson.LoadContent(`{"content": ""}`)
		err := dd.Set("content", wxconfig.Msg)
		if err != nil {
			fmt.Println(err)
		}
		err = d.Set("text", dd.Map())
		if err != nil {
			fmt.Println(err)
		}
		err = d.Set("msgtype", "text")
		if err != nil {
			fmt.Println(err)
		}
	}
	_, err3 := client.R().SetBody(d).Post(url)
	if err3 != nil {
		fmt.Println(err3)
	}

}
