package controller

import (
	"awesomeProject1/api"
	"awesomeProject1/database"
	"awesomeProject1/middleware"
	"awesomeProject1/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type User struct {
}

var Res string

func (*User) UserInfo(ctx *gin.Context) {
	name := ctx.Param("name")
	list := database.Search_browse_byname(name)
	ctx.JSON(200, gin.H{
		"message": "",
		"list":    list,
	})
}

func (*User) Userregister(ctx *gin.Context) {
	code := ctx.Query("code")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, _ := client.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wxdf131fbe9fc9887f&secret=659e5540cf1c36bdadd5f2e08fd877eb&js_code=" + code + "&grant_type=authorization_code")
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData = result.String()
	comma := strings.Index(resultData, ":")
	pos := strings.Index(resultData[comma:], resultData)
	d := strings.Index(resultData, ",")
	k := strings.Index(resultData, "}")
	_ = resultData[comma+pos+3 : d-1]
	res := resultData[d+11 : k-1]
	fmt.Println("请求结果", resultData)

	resp2,_ := client.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxdf131fbe9fc9887f&secret=659e5540cf1c36bdadd5f2e08fd877eb",)
	defer resp2.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp2.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)
	comma = strings.Index(resultData2, ":")
	d = strings.Index(resultData2,",")
	res2 := resultData2[comma+2 :d-1]
	log.Println(res2)
	err := database.Create_loginnotes(res)
	data := models.User{
		Account: res,
		Createdate : time.Now(),
		Name:  	"匿名用户"+strconv.Itoa(rand.Intn(99999)),
		Max_quantity: 5,
		Hour: 	time.Now().Hour(),
	}
	database.Createuser(data)
	ctx.JSON(200, gin.H{
		"openid": res,
		"access_token":res2,
		"err": err,
	})
}

func (*User) Userlogin(ctx *gin.Context) {
	openid := ctx.Query("openid")
	res := database.Search_user_byaccount(openid)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Login(ctx *gin.Context) {
	type Data struct {
		Name     string `json:"username"form:"username"`
		Password string `json:"password"form:"password"`
	}

	var logininfo Data
	_ = ctx.BindJSON(&logininfo)
	logininfo.Password = database.Md5Encode(logininfo.Password)
	log.Println(logininfo.Password)
	isSuccess := database.Login(logininfo.Name, logininfo.Password)

	response := make(gin.H)
	if isSuccess {
		token, err := middleware.NewToken(logininfo.Name)
		if err == nil {
			fmt.Println(err)
			response["toke"] = token
		}

		client := &http.Client{Timeout: 5 * time.Second}
		resp, _ := client.Get("https://api.weixin.qq.com/cgi-bin/token?appid=wxdf131fbe9fc9887f&secret=659e5540cf1c36bdadd5f2e08fd877eb&grant_type=client_credential")
		defer resp.Body.Close()
		var buffer [512]byte
		result := bytes.NewBuffer(nil)
		for {
			n, err := resp.Body.Read(buffer[0:])
			result.Write(buffer[0:n])
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}
		var resultData = result.String()
		log.Println(resultData)
		comma := strings.Index(resultData, ":")
		d := strings.Index(resultData,",")
		Res = resultData[comma+2 :d-1]
		log.Println(Res)
		ctx.JSON(200, gin.H{
			"code":    200,
			"id":      logininfo.Name,
			"message": "登录成功",
			"token":   token,
			"err":     err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "",
			"list":    "登录失败，密码错误或没有权限",
		})
	}

	//ctx.JSON(200, response)
}

func (*User) Createexp(ctx *gin.Context) {

	name := ctx.Query("name")
	type Form struct {
		Name string `form:"name"`
	}

	var form Form
	ctx.Bind(&form)
	ctx.JSON(200, gin.H{
		"message": form,
		"query":   name,
	})
}

func (*User) Createuser(ctx *gin.Context) {
	type Data struct {
		Name    string `json:"name"form:"name"`
		Account string `json:"account"form:"account"`
	}
	var user Data
	_ = ctx.BindJSON(&user)
	data := models.User{
		Name:    user.Name,
		Account: user.Account,
		Createdate:time.Now(),
		Max_quantity:5,
		Hour: time.Now().Hour(),
	}
	var err = database.Createuser(data)
	if err.Error != nil {
		ctx.JSON(200, gin.H{
			"message": "用户已存在",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "成功"})
	}

}

func (*User) Createcard(ctx *gin.Context) {
	/*type Data struct {
		Companyid       int      `json:"company_id" form:"company_id"`
		Useraccount     string   `json:"user_account" form:"user_account"`
		Privacy         bool     `json:"privacy" form:"privacy"`
		User_name       string   `json:"user_name" form:"user_name"`
		Company_name    string   `json:"company_name" form:"company_name"`
		Company_address string   `json:"company_address" form:"company_address"`
		Phone           string   `json:"phone" form:"phone"`
		Post            string   `json:"post" form:"post"`
		Email           string   `json:"email" form:"email"`
		Introduction    string   `json:"introduction" form:"introduction"`
		Picture         []string `json:"picture" form:"picture"`
		Video           []string `json:"video" form:"video"`
	}*/

	//type FormInfo struct {
	//	CompanyAddress string `json:"company_address"`
	//	Phone          string `json:"phone"`
	//	Email          string `json:"email"`
	//	Post           string `json:"post"`
	//	UserName       string `json:"user_name"`
	//}

	type Form struct {
		UserAccount string `json:"user_account"`
		FormInfo struct {
			CompanyAddress string `json:"company_address"`
			Phone          string `json:"phone"`
			Email          string `json:"email"`
			Post           string `json:"post"`
			UserName       string `json:"user_name"`
			Company_name	string `json:"company_name"`
			Introduction	string `json:"introduction"`
			Picture         []string `json:"picture" form:"picture"`
			Video           []string `json:"video" form:"video"`
			Privacy         bool     `json:"privacy" form:"privacy"`
			Avatar			string	`json:"avatar"`
			Mark			bool	`json:"mark"`
			Delete			[]string `json:"delete"`
		} `json:"fromInfo"`
	}

	/*loginip := strings.Split(ctx.Request.RemoteAddr, ":")[0]
	fmt.Println(loginip)*/


	var form Form
	//var form map[string]interface{}
	if err := ctx.Bind(&form); err != nil {
		log.Println(err)
	}
	log.Printf(form.UserAccount)
	log.Println("表单信息", form)
	log.Println(form.FormInfo.Picture)
	log.Println("默认",form.FormInfo.Mark)


	pic := strings.Join(form.FormInfo.Picture, ",")
	var card_id int
	video := strings.Join(form.FormInfo.Video, ",")
	database.Delete_uploadFile(form.FormInfo.Delete)
	data := models.Card_warehouse{
		Card_user_account: form.UserAccount,
		Card_user_name:    form.FormInfo.UserName,
		Card_company_name: form.FormInfo.Company_name,
		Card_address:      form.FormInfo.CompanyAddress,
		Card_phone:        form.FormInfo.Phone,
		Card_company_post: form.FormInfo.Post,
		Card_email:        form.FormInfo.Email,
		Card_introduction: form.FormInfo.Introduction,
		Card_picture:      pic,
		Card_video:        video,
		Card_privacy:      form.FormInfo.Privacy,
		Mark:				form.FormInfo.Mark,
		Card_browse:       0,
		Card_collect:      0,
		Website_address:   "",
		Card_mould_id:     0,
		Avatar:            form.FormInfo.Avatar,
	}
	user := database.Search_user_byaccount(form.UserAccount)
	log.Println(form.UserAccount)
	log.Println(user.Card_quantity)
	log.Println(user.Max_quantity)
	if user.Card_quantity < user.Max_quantity{
		if form.FormInfo.Mark == true{
			database.Updata_card_mark(form.UserAccount)
		}
		var err = database.Createcard(data)
		if err.Error != nil {
			ctx.JSON(200, gin.H{
				"message": "创建失败",
				"err":     err,
				"data":    data,
			})
		} else {
			log.Println("789")
			database.Updata_user(form.UserAccount,user.Card_quantity+1)
			log.Println("789")
			card_id = int(database.Search_card_last(form.UserAccount).Card_id)
			log.Println("头像",form.FormInfo.Avatar)
			avatar := Download(form.FormInfo.Avatar , card_id)
			database.Updata_card_avatar(card_id,avatar)
			log.Println(user.Card_quantity)
			if user.Card_quantity == 0{
				database.Updata_user_name(form.UserAccount,form.FormInfo.UserName,avatar)
			}
			card := strconv.Itoa(card_id)
			log.Println("fsfsfsfsf",card)
			ctx.JSON(200, gin.H{
				"message": "创建成功",
				"cardid":card,
			})
		}
	}else{
		ctx.JSON(200, gin.H{
			"message": "名片数已达上限",
			"qty":user.Card_quantity,
			"max":user.Max_quantity,
		})
	}

	/*for i := 0;i< len(form.FormInfo.Picture);i++{
		imgUrl := form.FormInfo.Picture[i]
		// Get the data
		log.Println(imgUrl)

		resp, err := http.Get(imgUrl)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// 创建一个文件用于保存
		n := rand.Intn(999)
		k := strconv.Itoa(n)
		os.Mkdir("./upload/"+strconv.Itoa(card_id)+"/",os.ModePerm)
		time := time.Now().Unix()
		name := strconv.Itoa(int(time))
		out, err := os.Create("./upload/"+strconv.Itoa(card_id)+"/"+k+name+".jpg")
		if err != nil {
			panic(err)
		}
		defer out.Close()

		// 然后将响应流和文件流对接起来
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
	}*/

}






func (*User) Search_user(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": database.Search_user(),
	})
}

func (*User) Search_user_byaccount(ctx *gin.Context) {
	useraccount := ctx.Query("user_account")
	ctx.JSON(200, gin.H{
		"message": database.Search_user_byaccount(useraccount),
	})
}

//func (*User) Createbrowse(ctx *gin.Context) {
//	type Data struct {
//		Useraccount string `json:"account"form:"account"`
//		Username    string `json:"username"form:"username"`
//		Cardid      int16  `json:"cardid"form:"cardid"`
//		Time        time.Time
//	}
//	var b Data
//	_ = ctx.BindJSON(&b)
//	log.Println(b)
//	data := models.BrowseRecord{
//		Useraccount: b.Useraccount,
//		Username:    b.Username,
//		Cardid:      b.Cardid,
//		Date:        time.Now(),
//	}
//	err := database.Createbrowse(b.Useraccount, int(b.Cardid))
//	if err != nil {
//		ctx.JSON(200, gin.H{
//			"message": "添加浏览记录失败",
//			"data":    data,
//		})
//	} else {
//		database.Updata_browse(b.Cardid)
//		ctx.JSON(200, gin.H{
//			"message": "添加浏览成功",
//		})
//	}
//
//}

func (*User) Updatainformation(ctx *gin.Context) {
	type Data struct {
		Picture []string `json:"picture"form:"picture"`
		Video   []string `json:"video"form:"video"`
		Cardid  int      `json:"card_id"form:"card_id"`
		Privacy bool     `json:"privacy"form:"privacy"`
	}
	var information Data
	_ = ctx.BindJSON(&information)
	pic := strings.Join(information.Picture, ",")
	video := strings.Join(information.Video, ",")
	data := models.Card_warehouse{
		Card_introduction: "",
		Card_picture:      pic,
		Card_video:        video,
		Card_privacy:      information.Privacy,
	}
	database.Updata_card(information.Cardid, data)

}

func (*User) Createinformation0(ctx *gin.Context) {
	type Data struct {
		Card_id                  int16  `json:"card_id"form:"card_id"`
		Card_information_type    int8   `json:"type"form:"type"`
		Card_information_cont    string `json:"cont"form:"cont"`
		Card_information_picture string `json:"picture"form:"picture"`
		Card_information_video   string `json:"video"form:"video"`
		Card_information_pdf     string `json:"pdf"form:"pdf"`
		Time                     time.Time
	}
	var information Data
	_ = ctx.BindJSON(&information)
	data := models.Card_information{
		Card_information_type:    information.Card_information_type,
		Card_information_card_id: information.Card_id,
		Card_information_cont:    information.Card_information_cont,
		Card_information_picture: information.Card_information_picture,
		Card_information_video:   information.Card_information_video,
		Card_information_pdf:     information.Card_information_pdf,
	}

	err := database.Createinformation(data)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "添加失败，名片不存在",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "名片信息添加成功",
		})
	}
	/*cardid := ctx.Query("card_id")
	card_id, _ := strconv.Atoi(cardid)
	var picture [] string
	picture = ctx.QueryArray("picture")
	var data models.Card_information
	k := 0
	for k=0;k< len(picture);k++  {
		data = models.Card_information{
			Card_information_type:    0,
			Card_information_card_id: int16(card_id),
			Card_information_cont:    "",
			Card_information_picture: picture[k],
			Card_information_video:   "",
			Card_information_pdf:     "",
		}
		err := database.Createinformation(data)
		if err != nil{
			ctx.JSON(200, gin.H{
				"message": "添加失败，名片不存在",
			})
		}else {
			ctx.JSON(200, gin.H{
				"message": "名片信息添加成功",
			})
		}

	}*/

}

func (*User) Createchat(ctx *gin.Context) {
	type Data struct {
		UserA_account string `json:"user_a_account"form:"user_a_account"`
		UserB_account string `json:"user_b_account"form:"user_b_account"`
		UserA_name    string `json:"user_a_name"form:"user_a_name"`
		UserB_name    string `json:"user_b_name"form:"user_b_name"`
		Cont          string `json:"cont"form:"cont"`
	}
	var chat Data
	_ = ctx.BindJSON(&chat)

	data := models.Chat{
		Chat_userA_account: chat.UserA_account,
		Chat_userB_account: chat.UserB_account,
		Chat_userA_name:    chat.UserA_name,
		Chat_userB_name:    chat.UserB_name,
		Chat_content:       chat.Cont,
		Chat_date:          time.Now(),
	}
	err := database.Createchat(data)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "发送失败，用户信息不匹配",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "发送成功",
		})
	}
}

func (*User) Createcollect(ctx *gin.Context) {
	type Collect struct {
		Card_id int `json:"card_id"`
		Account string `json:"account"`
	}
	var collect Collect
	ctx.BindJSON(&collect)
	res := database.Search_collect_byids(collect.Account, collect.Card_id)
	if res.Collect_id == 0 {
		data := models.Collect_record{
			User_account:      collect.Account,
			Card_id:      collect.Card_id,
			Collect_mark: true,
			Collect_date: time.Now(),
		}
		err := database.Createcollect(data).Error
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "收藏失败，用户或名片不存在",
			})
		} else {
			database.Collect_add(collect.Card_id)
			ctx.JSON(200, gin.H{
				"message": "收藏成功",
			})
		}
	} else if res.Collect_mark == false {
		err := database.Update_collectt(collect.Account, collect.Card_id)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "收藏失败，用户或名片不存在",
			})
		} else {
			database.Collect_add(collect.Card_id)
			ctx.JSON(200, gin.H{
				"message": "收藏成功",
			})
		}
	} else if res.Collect_mark == true {
		err := database.Update_collectf(collect.Account, collect.Card_id)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "取消收藏失败，用户或名片不存在",
			})
		} else {
			database.Collect_minus(collect.Card_id)
			ctx.JSON(200, gin.H{
				"message": "取消收藏成功",
			})
		}
	}

}

func (*User) Createkeyword(ctx *gin.Context) {
	name := ctx.Query("keyword_content")
	if name == "" {
		ctx.JSON(200, gin.H{
			"message": "标签不能为空",
		})
	} else {
		data := models.Keyword_warehouse{
			Keyword_content: name,
		}
		if database.Createkeyword(data) == true {
			ctx.JSON(200, gin.H{
				"message": "标签创建成功",
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": "标签已存在",
			})
		}
	}

}

func (*User) Createneed(ctx *gin.Context) {
	userid := ctx.Query("user_id")
	user_id, _ := strconv.Atoi(userid)

	y := ctx.Query("year")
	year, _ := strconv.Atoi(y)
	mo := ctx.Query("month")
	month, _ := strconv.Atoi(mo)
	d := ctx.Query("day")
	day, _ := strconv.Atoi(d)
	h := ctx.Query("hour")
	hour, _ := strconv.Atoi(h)
	m := ctx.Query("min")
	min, _ := strconv.Atoi(m)
	s := ctx.Query("sec")
	sec, _ := strconv.Atoi(s)

	data := models.Need_warehouse{
		Need_user_id:   int16(user_id),
		Need_user_name: ctx.Query("user_name"),
		Need_name:      ctx.Query("need_name"),
		Need_content:   ctx.Query("need_content"),
		Need_keyword:   ctx.Query("need_keyword"),
		Need_starttime: time.Now(),
		Need_endtime:   time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local),
	}
	err := database.Createneed(data)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "标签或用户名不存在",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "需求创建成功",
		})
	}
}

//查询
func (*User) Search_user_byname(ctx *gin.Context) {
	user_name := ctx.Query("user_name")
	res := database.Search_user_byname(user_name)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Search_card(ctx *gin.Context) {
	res := database.Search_card()
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Search_card_public(ctx *gin.Context) {
	type Paging struct {
		Account string `json:"account"`
		Page int `json:"page"`
		Size int `json:"size"`
		Keyword string `json:"keyword"`
		Sort string `json:"sort"`
		Desc bool `json:"desc"`
	}
	var pageing Paging
	ctx.BindJSON(&pageing)
	if pageing.Page == 0 {
		pageing.Page++
	}
	if pageing.Size == 0{
		pageing.Size = 5
	}
		res := database.Search_card_public_bykeyword(pageing.Keyword)
		message := database.Search_card_bypageing(pageing.Size, pageing.Page, pageing.Keyword ,pageing.Sort ,pageing.Desc)
		card := database.Card_stirng2array(message)
		type Collect_cards struct {
			Card models.Card
			Collect bool
		}
		var collect_card [100]Collect_cards
		for i := 0 ; i < len(card) ; i++{
			mark := database.Search_collect_byids(pageing.Account, int(card[i].Card_id)).Collect_mark
			collect_card[i] = Collect_cards{
				Card:    card[i],
				Collect: mark,
			}
		}
		page_all := len(res) / pageing.Size
		page_all_f := float32(len(res)) / float32(pageing.Size)
		if page_all_f > float32(page_all){
			page_all++
		}
		ctx.JSON(200, gin.H{
			"data":collect_card[:len(card)],
			"num": len(res),
			"len": pageing.Size,
			"page_now":pageing.Page,
			"page_all":page_all,
		})


}

func (*User) Search_card_public_bycomname(ctx *gin.Context) {
	comname := ctx.Query("comname")
	res := database.Search_card_public_bycomname(comname)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Search_card_bycomname(ctx *gin.Context) {
	comname := ctx.Query("comname")
	res := database.Search_card_bycomname(comname)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Search_chat_byuserid(ctx *gin.Context) {
	userid := ctx.Query("user_id")
	user_id, _ := strconv.Atoi(userid)
	res := database.Search_chat_byuserid(user_id)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Search_card_byaccount(ctx *gin.Context) {
	useraccount := ctx.Query("user_account")
	res := database.Search_card_byuseraccount(useraccount)
	user := database.Search_user_byaccount(useraccount)
	ctx.JSON(200, gin.H{
		"cards": res,
		"user": user,
	})
}

func (*User) Search_defcard_byaccount(ctx *gin.Context) {
	type Card struct {
		Useraccount string `json:"user_account"`
		Cardid string	`json:"card_id"`
	}
	var card Card
	ctx.BindJSON(&card)
	log.Println("sada"+card.Cardid)
	user := database.Search_user_byaccount(card.Useraccount)
	now := user.Card_quantity
	max := user.Max_quantity
	card_id,err := strconv.Atoi(card.Cardid)
	log.Println("int ",card_id,err)
	if card_id == 0 {
		log.Println("card_id为",card_id)
			result := database.Search_card_byuseraccount(card.Useraccount)
			res := database.Card_stirng2array(result)

			if len(result) > 0 {
				ctx.JSON(200, gin.H{
					"message": res[0:len(result)],
					"now": now,
					"max":max,
				})
			} else {
				ctx.JSON(200, gin.H{
					"message": "没有名片",
					"now": now,
					"max":max,
				})
			}
		} else{

			collect := database.Search_collect_byids(card.Useraccount , card_id).Collect_mark
			/*browse := database.Search_browse(card.Useraccount,card_id)
			if len(browse) == 0{*/
				//log.Println(len(browse))
				card_account := database.Search_card_byid(card_id).Card_user_account
				if card_account == card.Useraccount{
					err := database.Createbrowse(card.Useraccount, user.Name ,card_id)
					if err == nil{
						log.Println("21312")
						database.Browse_add(card_id)
					}
				}

			/*}
*/
			log.Println(card.Useraccount)
		r :=database.Search_card_byid(card_id)
		pic := database.S2a(r.Card_picture)
		video := database.S2a(r.Card_video)
		web_introduction_public := database.S2a(r.Web_introduction_public)
		web_picture_public := database.S2a(r.Web_picture_public)
		web_video_public := database.S2a(r.Web_video_public)
		//web_introduction := strings.Split(result.Web_introduction, ",")
		web_picture := database.S2a(r.Web_picture)
		web_video := database.S2a(r.Web_video)
		keyword := database.S2a(r.Keyword)
		var res = models.Card{
			Card_id:                 r.Card_id,
			Card_company_id:         r.Card_company_id,
			Card_user_account:       r.Card_user_account,
			Card_user_name:          r.Card_user_name,
			Card_company_name:       r.Card_company_name,
			Card_address:            r.Card_address,
			Card_phone:              r.Card_phone,
			Card_company_post:       r.Card_company_post,
			Card_email:              r.Card_email,
			Card_introduction:       r.Card_introduction,
			Card_picture:            pic,
			Card_video:              video,
			Card_privacy:            r.Card_privacy,
			Card_browse:             r.Card_browse,
			Card_collect:            r.Card_collect,
			Website_address:         r.Website_address,
			Card_mould_id:           r.Card_mould_id,
			Mark:                    r.Mark,
			Web_introduction:        r.Web_introduction,
			Web_picture:             web_picture,
			Web_video:               web_video,
			Web_introduction_public: web_introduction_public,
			Web_picture_public:      web_picture_public,
			Web_video_public:        web_video_public,
			Keyword:                 keyword,
			Avatar:                  r.Avatar,
		}
		ctx.JSON(200, gin.H{
			"message": res,
			"now": now,
			"max":max,
			"collect":collect,
			"err":err,
		})
	}

}
func (*User) Closecard(ctx *gin.Context) {
	cardid := ctx.Query("card_id")
	card_id, _ := strconv.Atoi(cardid)
	database.Updata_card_mark1(card_id)
}

func (*User) Search_card_byid(ctx *gin.Context) {
	id := ctx.Query("card_id")
	card_id, _ := strconv.Atoi(id)
	res := database.Search_card_byid(card_id)
	pic := strings.Split(res.Card_picture, ",")
	video := strings.Split(res.Card_video, ",")
	web_introduction_public := strings.Split(res.Web_introduction_public, ",")
	web_picture_public := strings.Split(res.Web_picture_public, ",")
	web_video_public := strings.Split(res.Web_video_public, ",")
	web_introduction := strings.Split(res.Web_introduction, ",")
	web_picture := strings.Split(res.Web_picture, ",")
	web_video := strings.Split(res.Web_video, ",")
	ctx.JSON(200, gin.H{
		"message": res,
		"pic":     pic,
		"video":   video,
		"web_introduction_public": web_introduction_public,
		"web_picture_public":web_picture_public,
		"web_video_public":web_video_public,
		"web_introduction":web_introduction,
		"web_picture":web_picture,
		"web_video":web_video,
	})
}

func (*User) Search_card_public_bykeyword(ctx *gin.Context) {
	type Search struct {
		Account string `json:"account"`
		Keyword string `json:"keyword"`
	}
	var search Search
	ctx.BindJSON(&search)
	res := database.Search_card_public_bykeyword(search.Keyword)
	ctx.JSON(200, gin.H{
		"message": res,
	})
	if len(database.Search_search_note_account(search.Account))>0{
		num := len(database.Search_search_note(search.Account, search.Keyword))
		if num > 0{
			database.Add_search_num(search.Keyword)
			database.Add_search_note(search.Account, search.Keyword)
		}else{

			if len(database.Search_keyword_num(search.Keyword)) > 0{
				database.Add_search_num(search.Keyword)

			}else{
				database.Create_search_num(search.Keyword)
			}
			data := models.Search_notes{
				Account: search.Account,
				Keyword: search.Keyword,
				Num:     1,
				Time:    time.Now(),
				Name:    database.Search_user_byaccount(search.Account).Name,
			}
			database.Create_search_notes(data)
		}
	}else{
		num := len(database.Search_search_note(search.Account, search.Keyword))
		if num > 0{
			database.Add_search_num(search.Keyword)
			database.Add_search_note(search.Account, search.Keyword)
		}else{

			if len(database.Search_keyword_num(search.Keyword)) > 0{
				database.Add_search_num(search.Keyword)

			}else{
				database.Create_search_num(search.Keyword)
			}
			data := models.Search_notes{
				Account: search.Account,
				Keyword: search.Keyword,
				Num:     1,
				Time:    time.Now(),
				Name:    database.Search_user_byaccount(search.Account).Name,
				First_search: true,
			}
			database.Add_search_user_num(search.Keyword)
			database.Create_search_notes(data)
		}
	}

}


//删除
func (*User) Deletecard_bycardid(ctx *gin.Context) {
	type Data struct {
		Cardid int    `json:"card_id"form:"card_id"`
		Userid string `json:"user_account"form:"user_account"`
	}
	var card Data
	_ = ctx.BindJSON(&card)

	err := database.Deletecard_bycardid(card.Cardid, card.Userid)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "失败",
			"nil":     err,
		})
	} else {
		qty := database.Search_user_byaccount(card.Userid).Card_quantity
		database.Updata_user(card.Userid,qty - 1)
		ctx.JSON(200, gin.H{
			"message": "删除成功",
		})
	}

}

func Download(url string, card_id int) string {
	imgPath := "/upload/avatar/"+strconv.Itoa(card_id)+".jpg"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create("."+imgPath)
	io.Copy(out, bytes.NewReader(body))
	return "http://192.168.16.126:8080"+imgPath
}

func (*User) Upload(ctx *gin.Context) {
	/*cardid := ctx.Query("card_id")
	card_id, _ := strconv.Atoi(cardid)
	form, _ := ctx.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		n := rand.Intn(999)
		k := strconv.Itoa(n)
		name := k+file.Filename
		// 上传文件到指定的路径
		os.Mkdir("./upload/"+cardid+"/",os.ModePerm)
		//os.Create(cardid)
		ctx.SaveUploadedFile(file, "./upload/"+cardid+"/"+name)

		v := database.Search_card_byid(card_id)
		v.Card_video = v.Card_video + "," + name
		data := models.Card_warehouse{
			Card_video:        v.Card_video,
		}
		database.Updata_card(card_id,data)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))*/
	file, err := ctx.FormFile("file")
	filename := file.Filename
	s := strings.LastIndex(filename,".")
	fname := filename[s:]
	n := rand.Intn(999)
	k := strconv.Itoa(n)
	time := time.Now().Unix()
	name := strconv.Itoa(int(time))
	url := "http://192.168.16.126:8080/upload/"+k+name+fname
	log.Println(url)
	error := ctx.SaveUploadedFile(file,"./upload/"+k+name+fname)
	if err != nil{
		ctx.JSON(200, gin.H{
			"err":error,
		})
	}else {
		ctx.JSON(200, gin.H{
			"url": url,
			"err":error,
		})
	}
}

func (*User) Deletevideo(ctx *gin.Context) {
	cardid := ctx.Query("card_id")
	card_id,_ := strconv.Atoi(cardid)
	type Data struct {
		Arr []string `json:"arr"`
	}
	var videos Data
	ctx.Bind(&videos)
	log.Println(videos)

	res := database.Search_card_byid(card_id)
	for k := 0;k < len(videos.Arr);k++{
		os.Remove("./upload/"+cardid+"/"+videos.Arr[k])
		res.Card_video = strings.Replace(res.Card_video, ","+videos.Arr[k], "", 1)
	}
		log.Println(res.Card_video)
}



func (*User) Data_analysis(ctx *gin.Context) {

	type Growth struct {
		Today_growth int `json:"todaygrowth"`
		Yestoday_growth int `json:"yestoday_growth"`
		Rate	string `json:"rate"`
	}

	var all_login models.User_visit


	all_login = database.Search_loginall()
	usersum := len(database.Search_user())
	log.Println("用户数", usersum)

	tgrowth := len(database.Search_user_bytime())
	yestodayuser := usersum - tgrowth
	log.Println("昨日用户数",yestodayuser)
	log.Println("今日增长数", tgrowth)
	ygrowth := len(database.Search_user_bytimey())
	log.Println("昨日增长数", ygrowth)

	type Data struct {
		Usersum int `json:"usersum"`
		Yestodayuser int `json:"yestodayuser"`
		Today_growth int `json:"today_growth"`
		Today_growthrate string `json:"today_growthrate"`
	}
	data := Data{
		Usersum:          usersum,
		Yestodayuser:     yestodayuser,
		Today_growth:     tgrowth,
		Today_growthrate: fmt.Sprintf("%.2f",float32(usersum-yestodayuser)/float32(yestodayuser)*100),
	}
	var growth Growth
	growth = Growth{
		Today_growth:    tgrowth,
		Yestoday_growth: ygrowth,
		Rate:            fmt.Sprintf("%.2f",float32(tgrowth)/float32(ygrowth)*100),
	}
	search := database.Search_search()
	bc_num := database.Search_bcnum()
	ctx.JSON(200, gin.H{
		"user":data,
		"visit":all_login,
		"growth":growth,
		"search":search,
		"bc_num":bc_num,
	})
}

func (*User) Data_analysis_chart(ctx *gin.Context) {
	type Time struct {
		Begin_date string `json:"begin_date"`
		End_date string `json:"end_date"`
	}
	var time Time
	ctx.BindJSON(&time)
	if time.Begin_date == ""{
		time.Begin_date = "1010-01-01"
	}
	if time.End_date == ""{
		time.End_date = "9090-01-01"
	}
	note := database.Search_loginnotes(time.Begin_date,time.End_date)
	ctx.JSON(200, gin.H{
		"chart":note,
		"time":time,
	})
}


func (*User) Createmould(ctx *gin.Context) {
	type Form struct {
		UserAccount string `json:"user_account"`
		FormInfo struct {
			Card_mould_style			int16			`json:"card_mould_style"`
			Card_mould_name				string			`json:"card_mould_name"`
			Card_mould_company_id		int16 			`json:"card_mould_company_id"`
			Card_mould_company_name		string			`json:"card_mould_company_name"`
			Card_mould_address			string			`json:"card_mould_address"`
			Card_mould_post				string			`json:"card_mould_post"`
			Card_mould_website_style	int16			`json:"card_mould_website_style"`
			Card_mould_website_address	string			`json:"card_mould_website_address"`
			Web_introduction			[]string			`json:"web_introduction"`
			Web_picture					[]string			`json:"web_picture"`
			Web_video					[]string			`json:"web_video"`
			Web_introduction_public		[]string			`json:"web_introduction_public"`
			Web_picture_public			[]string			`json:"web_picture_public"`
			Web_video_public			[]string			`json:"web_video_public"`
		} `json:"fromInfo"`
	}
	var mould Form
	ctx.BindJSON(&mould)
	data := models.Card_mould{
		User_account:               mould.UserAccount,
		Card_mould_style:           mould.FormInfo.Card_mould_style,
		Card_mould_name:            mould.FormInfo.Card_mould_name,
		Card_mould_company_name:    mould.FormInfo.Card_mould_company_name,
		Card_mould_address:         mould.FormInfo.Card_mould_address,
		Card_mould_post:            mould.FormInfo.Card_mould_post,
		Card_mould_website_style:   mould.FormInfo.Card_mould_website_style,
		Web_introduction:			strings.Join(mould.FormInfo.Web_introduction, ","),
		Web_picture:				strings.Join(mould.FormInfo.Web_picture, ","),
		Web_video:					strings.Join(mould.FormInfo.Web_video, ","),
		Web_introduction_public:  	strings.Join(mould.FormInfo.Web_introduction_public, ","),
		Web_picture_public:			strings.Join(mould.FormInfo.Web_picture_public, ","),
		Web_video_public:			strings.Join(mould.FormInfo.Web_video_public, ","),

	}
	err := database.Create_cardmould(data)
	ctx.JSON(200, gin.H{
		"message": err,
	})
}

func (*User) Searchmould(ctx *gin.Context) {
	mouldid := ctx.Query("mould_id")
	mould_id,_ := strconv.Atoi(mouldid)
	res := database.Search_cardmould_bymouldid(mould_id)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Updata_card(ctx *gin.Context) {
	type Form struct {
		Card_id int `json:"card_id"`
		FormInfo struct {
			CompanyAddress string `json:"company_address"`
			Phone          string `json:"phone"`
			Email          string `json:"email"`
			Post           string `json:"post"`
			UserName       string `json:"user_name"`
			Company_name	string `json:"company_name"`
			Introduction	string `json:"introduction"`
			Picture         []string `json:"picture" form:"picture"`
			Video           []string `json:"video" form:"video"`
			Privacy         bool     `json:"privacy" form:"privacy"`
			Avatar			string	`json:"avatar"`
			Mark            bool	`json:"mark"`
			Delete 			[]string `json:"delete"`
		} `json:"fromInfo"`
	}

	var form Form
	//var form map[string]interface{}
	if err := ctx.Bind(&form); err != nil {
		log.Println(err)
	}
	avatar := Download(form.FormInfo.Avatar,form.Card_id)
	log.Println("头像",avatar)
	log.Println("图片", form.FormInfo.Picture)
	log.Println("sad",form.FormInfo.Mark)

	pic := strings.Join(form.FormInfo.Picture, ",")
	video := strings.Join(form.FormInfo.Video, ",")
	database.Delete_uploadFile(form.FormInfo.Delete)
	log.Println(pic)
	data := models.Card_warehouse{
		Card_user_name:          form.FormInfo.UserName,
		Card_company_name:       form.FormInfo.Company_name,
		Card_address:            form.FormInfo.CompanyAddress,
		Card_phone:              form.FormInfo.Phone,
		Card_company_post:       form.FormInfo.Post,
		Card_email:              form.FormInfo.Email,
		Card_introduction:       form.FormInfo.Introduction,
		Card_picture:            pic,
		Card_video:              video,
		Card_privacy:            form.FormInfo.Privacy,
		Mark: 					form.FormInfo.Mark,
		Card_browse:             0,
		Card_collect:            0,
		Website_address:         "",
		Card_mould_id:           0,
		Web_introduction:        "",
		Web_picture:             "",
		Web_video:               "",
		Web_introduction_public: "",
		Web_picture_public:      "",
		Web_video_public:        "",
		Keyword:                 "",
		Avatar:                  avatar,
	}
	log.Println(form.Card_id)
	//card_id,_ := strconv.Atoi(form.Card_id)
	//log.Println(card_id)
	database.Updata_card(form.Card_id,data)

	ctx.JSON(200, gin.H{
		"card_id": form.Card_id,
	})
}

func (*User) Updatapassword(ctx *gin.Context){
	type Admin struct {
		Admin string	`json:"admin"`
		Password string	`json:"password"`
	}
	var admin Admin
	ctx.BindJSON(&admin)
	err := database.Updata_password(admin.Admin,admin.Password)
	ctx.JSON(200, gin.H{
		"message": err,
	})
}

func (*User) Search_notes(ctx *gin.Context){
	type Notes struct {
		Name string `json:"name"`
		Keyword string `json:"keyword"`
		Num string `json:"num"`
		Sdate string `json:"sdate"`
		Edate string `json:"edate"`
	}
	var notes Notes
	ctx.BindJSON(&notes)
	if notes.Sdate == ""{
		notes.Sdate = "1010-1-1"
	}
	if notes.Edate == ""{
		notes.Edate = "8080-1-1"
	}
	log.Println("333",notes.Sdate)
	num,_ := strconv.Atoi(notes.Num)
	res := database.Search_search_notes(notes.Name,notes.Keyword,num,notes.Sdate,notes.Edate)
	ctx.JSON(200, gin.H{
		"message": res,
	})
}

func (*User) Info(ctx *gin.Context) {
	user_id:= ctx.GetString("userid")
	log.Println(user_id)
	res := database.Search_user_byadmin(user_id)
	ctx.JSON(200, gin.H{
		"message": res,
	})

}

func (*User) Nav(ctx *gin.Context) {
	type Meta struct {
		Meta_id		int `json:"meta_id"`
		Title  		string `json:"title"`
		Show		bool `json:"show"`
		Icon		string `json:"icon"`
		Hidden_header_content bool `json:"hiddenHeaderContent"`
		Keep_alive bool `json:"keepAlive"`
	}
	type Res struct {
		Id int `json:"id"`
		Parent_id int `json:"parent_id"`
		Name	string `json:"name"`
		Path    string `json:"path"`
		Meta	Meta `json:"meta"`
		Redirect string `json:"redirect"`
		Component string `json:"component"`
	}


	user_id:= ctx.GetString("userid")
	role := database.Search_user_byadmin(user_id)
	res := strings.Split(role.Roleid, ",")
	log.Println(res)
	var Result [100]Res
	R := Result
	for i := 0;i < len(res);i++{
		x,_ := strconv.Atoi(res[i])
		r := database.Search_role(x)
		meta_id := r.Meta_id
		meta := database.Search_meta(meta_id)
		var m Meta
		m = Meta{
			Meta_id: 	meta.Meta_id,
			Title: 		meta.Title,
			Show: 		meta.Show,
			Icon: 		meta.Icon,
			Hidden_header_content: meta.Hidden_header_content,
			Keep_alive:  meta.Keep_alive,
		}


		R[i] = Res{
			Id:        x,
			Parent_id: r.Parent_id,
			Name:      r.Name,
			Path:      r.Path,
			Meta:      m,
			Redirect:  r.Redirect,
			Component: r.Component,
		}
		log.Println(m)

	}
	//log.Println("asd",R[0:len(res)])
	r1 := R[0:len(res)]
	ctx.JSON(200, gin.H{
		"result":r1,
	})
}

func (*User) Updata_defcard(ctx *gin.Context) {
	type Card struct {
		Card_id int `json:"card_id"`
	}
	var card Card
	ctx.BindJSON(&card)
	defCard := database.Search_card_byid(card.Card_id)
	account := defCard.Card_user_account
	name := defCard.Card_user_name
	avatar := defCard.Avatar
	err := database.Updata_card_mark0(account)
	error := database.Updata_card_mark1(card.Card_id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"err":err,
		})
		if error != nil{
			ctx.JSON(200, gin.H{
				"err":error,
			})
		}
	}else{

		database.Update_user_name(account,name)
		database.Update_user_avatar(account,avatar)
		ctx.JSON(200, gin.H{
			"message":"修改成功",
		})
	}
}

func (*User) Keyword_ranking(ctx *gin.Context) {
	type Query struct {
		User_name string `json:"user_name"`
		Keyword string `json:"keyword"`
		Begin_date string `json:"begin_date"`
		End_date string `json:"end_date"`
	}
	var query Query
	ctx.BindJSON(&query)
	if query.Begin_date == ""{
		query.Begin_date = "1010-01-01"
	}
	if query.End_date == ""{
		query.End_date = "9090-01-01"
	}
	res := database.Keyword_ranking(query.User_name, query.Keyword, query.Begin_date, query.End_date)
	ctx.JSON(200, gin.H{
		"keyword_ranking":res,
	})
}

func (*User) Md5Encode(ctx *gin.Context) {
	type Md5 struct {
		Password string `json:"password"`
	}
	var md5 Md5
	ctx.BindJSON(&md5)
	log.Println(md5.Password)
	res := database.Md5Encode(md5.Password)
	ctx.JSON(200, gin.H{
		"md5":res,
	})
}

func (*User) Encode(ctx *gin.Context) {
	type Encode struct {
		Content string `json:"content"`
	}
	var encode Encode
	ctx.BindJSON(&encode)
	chatCode1 := database.Encode(encode.Content)
	chatCode2 := database.Encode(chatCode1)
	ctx.JSON(200, gin.H{
		"code":chatCode2,
	})
}

func (*User) Decode(ctx *gin.Context) {
	type Encode struct {
		Content string `json:"content"`
	}
	var encode Encode
	ctx.BindJSON(&encode)
	chat1 := database.Decode(encode.Content)
	chat2 := database.Decode(chat1)
	ctx.JSON(200, gin.H{
		"code":chat2,
	})
}



func (*User) Share(ctx *gin.Context) {
	type Share struct {
		Card_id int `json:"card_id"`
		Share_account string `json:"share_account"`
		Visit_account string `json:"visit_account"`
	}
	var share Share
	ctx.BindJSON(&share)
	if share.Share_account != share.Visit_account{
		sname := database.Search_user_byaccount(share.Share_account).Name
		vname := database.Search_user_byaccount(share.Visit_account).Name
		err := database.Create_share(share.Card_id, share.Share_account, share.Visit_account,sname,vname)
		if err != nil{
			ctx.JSON(200, gin.H{
				"err":err,
			})
		}else{

			database.Updata_share(share.Share_account)
			ctx.JSON(200, gin.H{
				"message":"添加分享成功",
			})
		}
	}else{
		ctx.JSON(200, gin.H{
			"message":"查看自己的名片不记录",
		})
	}


}

func (*User) Search_share(ctx *gin.Context) {
	type Share struct {
		Card_id int `json:"card_id"`
		Share_name string `json:"share_name"`
		visit_name string `json:"visit_name"`
		Begin_time string `json:"begin_time"`
		End_time string `json:"end_time"`
	}
	var share Share
	ctx.BindJSON(&share)
	if share.Begin_time == ""{
		share.Begin_time = "1010-01-01"
	}
	if share.End_time == ""{
		share.End_time = "9090-01-01"
	}
	res := database.Search_share(share.Card_id, share.Share_name, share.visit_name, share.Begin_time ,share.End_time)
	ctx.JSON(200, gin.H{
		"share":res,
	})
}




func (*User) Search_collect_byaccount(ctx *gin.Context) {
	type Collect struct {
		Account string `json:"account"`
	}
	var collect Collect
	var collect_card [100]models.Collect_card2
	var returnCard [100]models.Return_card2
	ctx.BindJSON(&collect)
	res := database.Search_collect_byuserid(collect.Account)
	for i := 0 ; i < len(res) ; i++{
		card := database.Search_card_byid(res[i].Card_id)
		collect_card[i] = models.Collect_card2{
			Card_id:      res[i].Card_id,
			Collect_date: res[i].Collect_date.Format("2006-01-02 15:04:05"),
			Username:  		 card.Card_user_name,
			Card:         card,
		}
	}
	return_card := database.Search_return_card(collect.Account)
	for j := 0; j < len(return_card); j++{
		card := database.Search_card_byid(return_card[j].CardId)
		returnCard[j] = models.Return_card2{
			Card_id:     return_card[j].CardId,
			Return_date: return_card[j].Date.Format("2006-01-02 15:04:05"),
			Username:    card.Card_user_name,
			Card:        models.Card_warehouse{},
		}
	}
	ctx.JSON(200, gin.H{
		"collect":collect_card[:len(res)],
		"returnData":returnCard[:len(return_card)],
	})
}

func (*User) Search_keywords_byaccount(ctx *gin.Context) {
	type Account struct {
		Account string `json:"account"`
	}
	var account Account
	ctx.BindJSON(&account)
	log.Println(account.Account)
	public := database.Search_keyword_pulic()
	my := database.Search_keyword_byaccount(account.Account)
	ctx.JSON(200, gin.H{
		"public":public,
		"my":my,
	})
}


func (*User) Code(ctx *gin.Context) {
	type Url struct {
		Path string `json:"path"`
	}
	var addr Url
	ctx.BindJSON(&addr)
	client := &http.Client{Timeout: 5 * time.Second}
	resp2,_ := client.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxdf131fbe9fc9887f&secret=659e5540cf1c36bdadd5f2e08fd877eb",)
	defer resp2.Body.Close()
	var buffer2 [512]byte
	result2 := bytes.NewBuffer(nil)
	for {
		n, err := resp2.Body.Read(buffer2[0:])
		result2.Write(buffer2[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	var resultData2 = result2.String()
	log.Println("res2",resultData2)
	comma := strings.Index(resultData2, ":")
	d := strings.Index(resultData2,",")
	res2 := resultData2[comma+2 :d-1]
	log.Println(res2)

	b ,err := json.Marshal(addr)
	if err != nil {
		log.Println("json format error:", err)
		return
	}
	log.Println(addr.Path)
	log.Println("b",b)
	body := bytes.NewBuffer(b)
	log.Println("body",body)
	resp,_ := client.Post("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="+res2,"content-type:application/json;charset=UTF-8",
		body)
	resp.Header.Set("Content-Type", "application/json; encoding=utf-8")
	defer resp.Body.Close()


	result, err := ioutil.ReadAll(resp.Body)

	/*var resultData = result.String()
	log.Println("res",resultData)
	comma = strings.Index(resultData, ":")
	d = strings.Index(resultData,",")*/
	//res := resultData[comma+2 :d-1]
	log.Println("resp",result)
	/*f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(result)*/
	ctx.JSON(200, gin.H{
		"code":result,
	})
	/*n := rand.Intn(999)
	k := strconv.Itoa(n)
	time := time.Now().Unix()
	name := strconv.Itoa(int(time))
	url := "http://192.168.16.126:8080/upload/"+k+name+fname
	log.Println(url)
	error := ctx.SaveUploadedFile(result,"./upload/"+k+name+fname)*/
}

func (*User) Search_share_bycardid(ctx *gin.Context) {
	type Card struct {
		Card_id string `json:"card_id"`
	}
	var card Card
	ctx.BindJSON(&card)
	card_id,_ := strconv.Atoi(card.Card_id)
	share := database.Search_share_bycardid(card_id)
	ctx.JSON(200, gin.H{
		"share":share,
	})
}

func (*User) Regular(ctx *gin.Context) {
	type String struct {
		Str string `json:"str"`
	}
	var string String
	ctx.BindJSON(&string)
	log.Println(string.Str)
	reg, err := regexp.Compile("[数控铣]")
	if err != nil{
		log.Println(err)
	}
	res := reg.MatchString(string.Str)
	ctx.JSON(200, gin.H{
		"share":res,
	})
}

func (*User) SubscribeMessage(ctx *gin.Context) {
	type SubscribeMessage struct {
		Account string `json:"account"`
		CardId	int `json:"card_id"`
		Template_id string `json:"template_id"`
		Page string `json:"page"`
		Data models.Data `json:"data"`
		Miniprogram_state string `json:"miniprogram_state"`
		Lang string `json:"lang"`
	}
	var subscribeMessage SubscribeMessage
	ctx.BindJSON(&subscribeMessage)
	accessToken := api.GetAccessToken()
	log.Println(accessToken)
	log.Println("接收",subscribeMessage)
	data := models.SubscribeMessage{
		Access_token:      accessToken,
		Touser:            subscribeMessage.Account,
		Template_id:       subscribeMessage.Template_id,
		Page:              subscribeMessage.Page,
		Data:              subscribeMessage.Data,
		Miniprogram_state: subscribeMessage.Miniprogram_state,
		Lang:              subscribeMessage.Lang,
	}
	err := api.SubscribeMessage(accessToken,data)
	returnCard := models.Return_card{
		AcceptAccount: subscribeMessage.Account,
		CardId:        subscribeMessage.CardId,
	}
	type Success struct {
		Errcode int `json:"errcode"`
		Errmsg string `json:"errmsg"`
	}

	if err == "{\"errcode\":0,\"errmsg\":\"ok\"}"{
		database.CreateReturnCard(returnCard)
	}




	log.Println("错误",err)
	ctx.JSON(200, gin.H{
		"err":err,
	})
}


func (*User) Userlist(ctx *gin.Context) {

}
