/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2022/4/9
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func fetch (url string) string {
	//fmt.Println("Fetch Url", url)
	/*client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := client.Do(req)*/
	resp,err:=http.Get(url)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}

func parse(url string)string {
	body := fetch(url)
	//body="如果女人一天走了两万步…</h2> <span class=\"subtitleview\">发布时间：2022-04-08 21:10&nbsp;&nbsp;441阅</span><div id=\"KL_margin\" style=\"margin:8px;\"></div><p align=\"center\"></p><!--listS-->1、如果女人一天走了两万步🚶，那么她是在逛街买东西；一天只走了18步，那么她是在网上买东西。<br/><br/>2、女：老公，我吃相丑不？<br/>男：不丑。<br/>女：我睡相丑不？<br/>男：不丑。<br/>女：那我啥丑？<br/>男：面相。<br/><br/>3、吃自助一开始是本能，吃到最后拼的是意志。<br/><br/>4、“天天看你减肥，也没见你瘦到哪儿去!”<br/>“你们只看我减肥没有瘦下来，你们有没有想过如果我不减肥会胖成什么样子！！！”<br/>嗯。。。这说的毫无破绽。<br/><br/>5、一直以为芝麻是从草莓身上来的，但是黑芝麻怎么来的，这件事困惑了我许多年，如今我看见了火龙果。<br/><br/>6、什么叫缺乏独立思考？<br/>就是别人问你晚上吃什么，你说随便。<br/><br/>7、饭桌上，我：“吃火锅的时候你最无法忍受那种行为？”<br/>好友：“咬过的东西没熟吐出来再煮。”<br/>我：“滚！！！”<!--listE-->"
	rm:=regexp.MustCompile(`(?i:<!--listS-->).*<!--listE-->`)
	re,_:=regexp.Compile(`<br/>`)
	text:=rm.FindStringSubmatch(body)
	var str string
	for _, s := range text {
		str+=re.ReplaceAllString(s,"\n")
	}
	return "\nFetch Url\t"+url+"\n"+str
}

func main() {
	var str string
	for i := 0; i < 10; i++ {
		str+=parse("http://xiaodiaodaya.cn/article/view.aspx?id="+strconv.Itoa(i+160))
	}
	fmt.Println(str)

	f,err:=os.Create("lv1/joke.txt")
	if err!=nil{
		fmt.Println(err)
	}
	n,err:=f.WriteString(str)
	if err!=nil{
		panic(err)
	}
	fmt.Println("wrote",n,"bytes")

}
