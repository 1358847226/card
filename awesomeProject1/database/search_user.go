package database

import (
	"awesomeProject1/connect"
	_ "awesomeProject1/connect"
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	//"awesomeProject1/middleware"
	"awesomeProject1/models"
	_ "fmt"
	//_ "webapp/middleware/jwt"
)

func	Search_user() []models.User {
	var users []models.User
	connect.Db.Debug().Select([]string{"user_name", "user_power", "user_share", "account", "createdate", "avatar", "card_quantity", "max_quantity"}).Find(&users)
	return users
}


func	Search_user_byaccount(account string) models.User {
	user := models.User{}
	connect.Db.Where("account=?",account).Find(&user)
	return user
}

func	Search_user_byname(name string) []models.User{
	user := []models.User{}
	connect.Db.Where("user_name like ?","%"+name+"%").Find(&user)
	return user
}

func	Search_user_bypower(power int) []models.User{
	user := []models.User{}
	connect.Db.Where("user_power=?",power).Find(&user)
	return user
}

func	Search_card() []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Find(&card)
	return card
}

func	Search_card_public() []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Where("card_privacy=?",0).Find(&card)
	return card
}

func	Search_card_bypageing(size int,page int,keyword string,sort string ,desc bool) []models.Card_warehouse {
	var card []models.Card_warehouse
	if sort == ""{
		connect.Db.Order(" card_id").Where("keyword like ?","%"+keyword+"%").Limit(size).Offset((page - 1) * size).Find(&card)
	}else{
		if desc == true{

				connect.Db.Order(sort+" desc").Where("keyword like ?","%"+keyword+"%").Limit(size).Offset((page - 1) * size).Find(&card)


		}else{
				connect.Db.Where("keyword like ?","%"+keyword+"%").Order(sort).Limit(size).Offset((page - 1) * size).Find(&card)
		}

	}


	return card
}


func	Search_card_public_bycomid(id int) []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Where("card_company_id=? and card_privacy=?",id,0).Find(&card)
	return card
}

func	Search_card_bycomid(id int) []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Where("card_company_id=?",id).Find(&card)
	return card
}

func	Search_card_public_byname(name string) []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Where("card_user_name like ? and card_privacy=?","%"+name+"%",0).Find(&card)
	return card
}

func	Search_card_byname(name string) []models.Card_warehouse {
	var card []models.Card_warehouse
	connect.Db.Where("card_user_name like ?","%"+name+"%").Find(&card)
	return card
}

func	Search_browse_byname(name string) []models.BrowseRecord {
	var browse []models.BrowseRecord
	connect.Db.Where("user_name like ?","%"+name+"%").Find(&browse)
	return browse
}

func	Search_browse(account string,card_id int) []models.BrowseRecord {
	var browse []models.BrowseRecord
	connect.Db.Debug().Where("user_account = ? and card_id = ?",account ,card_id).Find(&browse)
	return browse
}

func	Search_browse_bycardid(id int) []models.BrowseRecord {
	var browse []models.BrowseRecord
	connect.Db.Where("card_id=?",id).Find(&browse)
	return browse
}

func	Search_cardmould_bymouldid(id int) models.Card_mould{
	var mould models.Card_mould
	connect.Db.Where("card_mould_id=?",id).Find(&mould)
	return mould
}

func	Search_cardmould_bycomname(name string) []models.Card_mould {
	var	mould []models.Card_mould
	connect.Db.Where("card_mould_company_name like ?","%"+name+"%").Find(&mould)
	return mould
}

func	Search_card_bycomname(name string) []models.Card_warehouse{
	var card []models.Card_warehouse
	connect.Db.Where("card_company_name like ?","%"+name+"%").Find(&card)
	return card
}

func	Search_card_last(account string) models.Card_warehouse{
	var card models.Card_warehouse
	connect.Db.Last(&card)
	return card
}

func	Search_card_public_bycomname(name string) []models.Card_warehouse{
	var card []models.Card_warehouse
	connect.Db.Where("card_company_name like ? and card_privacy=?","%"+name+"%",0).Find(&card)
	return card
}

func	Search_chat_byuserid(id int) []models.Chat{
	var chat []models.Chat
	connect.Db.Where("chat_userA_id=? or chat_userB_id=?",id,id).Find(&chat)
	return chat
}

func	Search_chat_byuserids(id1 int ,id2 int) []models.Chat{
	var chat []models.Chat
	connect.Db.Where("chat_userA_id=? and chat_userB_id=?",id1,id2).Or("chat_userA_id=? and chat_userB_id=?",id2,id1).Find(&chat)
	return chat
}

func	Search_chat_byusernames(name1 string,name2 string) []models.Chat{
	var chat []models.Chat
	connect.Db.Where("chat_userA_name like ? and chat_userB_name like ?","%"+name1+"%","%"+name2+"%").Or("chat_userA_name like ? and chat_userB_name like ?","%"+name2+"%","%"+name1+"%").Find(&chat)
	return chat
}

func	Search_collect_byuserid(account string) []models.Collect_record{
	var collect []models.Collect_record
	connect.Db.Where("user_account = ? and collect_mark = 1",account).Find(&collect)
	return collect
}

func	Search_return_card(account string) []models.Return_card{
	var return_card []models.Return_card
	connect.Db.Where("accept_account = ? ",account).Find(&return_card)
	return return_card
}

func	Search_collect_byids(account string,card_id int) models.Collect_record{
	var collect models.Collect_record
	err := connect.Db.Where("user_account = ? and card_id = ?",account,card_id).Find(&collect).Error
	log.Println(err)
	return collect
}


func Search_collect_byusername(name string)  []models.Collect_record{
	var collect []models.Collect_record
	connect.Db.Where("user_name like ? and collect_mark = 1","%"+name+"%").Find(&collect)
	return collect
}

func	Search_company_bycomid(id int) models.Company_warehouse{
	var company models.Company_warehouse
	connect.Db.Where("company_id = ?",id).Find(&company)
	return company
}

func Search_company_byname(name string)  []models.Company_warehouse{
	var company []models.Company_warehouse
	connect.Db.Where("company_name like ?","%"+name+"%").Find(&company)
	return company

}

/*func Search_keyword() []models.Keyword_warehouse{
	var keywords []models.Keyword_warehouse
	connect.Db.Find(&keywords)
	for _,keyword := range keywords {
		fmt.Println(keyword.Keyword_content)

	}
	return nil

}*/

func Search_keyword() []models.Keyword_warehouse{
	var keywords []models.Keyword_warehouse
	connect.Db.Find(&keywords)
	return keywords

}

func Search_need_byid(id int) []models.Need_warehouse{
	var need []models.Need_warehouse
	connect.Db.Where("need_user_id=?",id).Find(&need)
	return need
}

func Search_need_byusername(name string) []models.Need_warehouse {
	var need  []models.Need_warehouse
	connect.Db.Where("need_user_name like ?","%"+name+"%").Find(&need)
	return need
}

func Search_need_bykeyword(keyword1 string, keyword2 string, keyword3 string) []models.Need_warehouse{
	var need []models.Need_warehouse
	connect.Db.Where("need_keyword = ? or need_keyword = ? or need_keyword = ?",keyword1,keyword2,keyword3).Find(&need)
	return need
}

func Search_product_informaion_bytypePid(t int ,id int) []models.Product_information {
	var product []models.Product_information
	connect.Db.Where("product_information_type=? and product_information_product_id=?",t,id).Find(&product)
	return product
	
}

func Search_product_informaion_bytypePname(t int ,name string) []models.Product_information {
	var product []models.Product_information
	connect.Db.Where("product_information_type=? and product_information_product_name like ?",t,"%"+name+"%").Find(&product)
	return product

}

func Search_product_byid(id int) []models.Product_warehouse{
	var product []models.Product_warehouse
	connect.Db.Where("product_id = ?",id).Find(&product)
	return product
}

func Search_product_bycomid(id int) []models.Product_warehouse{
	var product []models.Product_warehouse
	connect.Db.Where("product_company_id = ?",id).Find(&product)
	return product
}

func Search_product_bycomname(name string) []models.Product_warehouse{
	var product []models.Product_warehouse
	connect.Db.Where("product_company_name like ?","%"+name+"%").Find(&product)
	return product
}

func Search_product_byPname(name string) []models.Product_warehouse{
	var product []models.Product_warehouse
	connect.Db.Where("product_name like ?","%"+name+"%").Find(&product)
	return product
}

func Search_product_bykeyword(keyword string) []models.Product_warehouse{
	var product []models.Product_warehouse
	connect.Db.Where("product_keyword = ?",keyword).Find(&product)
	return product
}

func Search_keyword_bycon(keyword string) []models.Keyword_warehouse {
	var content []models.Keyword_warehouse
	connect.Db.Where("keyword_content = ?",keyword).Find(&content)
	return content
}

func Login(user string, password string) bool {
	var login []models.User
	connect.Db.Debug().Where("admin_id = ? and user_password= ? and user_power=2", user, password).Find(&login)
	if len(login) > 0 {
		return true
	}else{
		return false
	}


}

func Search_card_byuseraccount(account string) []models.Card_warehouse{
	var card []models.Card_warehouse
	connect.Db.Where("card_user_account=?",account).Find(&card)
	return card
}
func Search_defcard_byuseraccount(account string) models.Card_warehouse{
	var card models.Card_warehouse
	connect.Db.Where("card_user_account=? and mark = 1",account).Find(&card)
	return card
}

func Search_card_byid(id int) models.Card_warehouse{
	var card models.Card_warehouse
	connect.Db.Where("card_id = ?",id).Find(&card)
	return card
}

func Search_user_bytime() []models.User{
	var user []models.User
	t := time.Now()
	today := time.Now().Format("2006-01-02 15:04:05")
	yesterday := t.AddDate(0,0,-1).Format("2006-01-02 15:04:05")
	connect.Db.Debug().Where("createdate BETWEEN ? AND ?", yesterday,today).Find(&user)

	return user
}

func Search_user_bytimey() []models.User{
	var user []models.User
	t := time.Now()
	lastday := t.AddDate(0,0,-2).Format("2006-01-02 15:04:05")
	yesterday := t.AddDate(0,0,-1).Format("2006-01-02 15:04:05")
	connect.Db.Debug().Where("createdate BETWEEN ? AND ?", lastday,yesterday).Find(&user)

	return user
}
type Note struct {
	Time string `json:"time"`
	VisitNum int `json:"visit_num"`
	GrowthNum	int `json:"growth_num"`
}
func Search_loginnotes(begin_date string, end_date string) ([12] Note) {
	var notes1 []models.Login_notes
	var growth []models.User
	var n[12] Note
	for i := 0;i < 12;i++{
		connect.Db.Debug().Where("time BETWEEN ? AND ? and hour >= ? and hour < ?",begin_date,end_date,(i+1)*2-2,(i+1)*2).Find(&notes1)
		connect.Db.Debug().Where("createdate BETWEEN ? AND ? and hour >= ? and hour < ?",begin_date,end_date,(i+1)*2-2,(i+1)*2).Find(&growth)
		n[i] = Note{
			Time:strconv.Itoa((i+1)*2-2)+":00-"+strconv.Itoa((i+1)*2)+":00",
			VisitNum: len(notes1),
			GrowthNum:len(growth),
			}
	}


	return n
}



func Search_loginall() models.User_visit {
		var today_visit []models.Login_notes
		var yesterday_visit []models.Login_notes
		var all_visit []models.Login_notes
		t := time.Now()
		today := time.Now().Format("2006-01-02 15:04:05")
		yesterday := time.Now().Format("2006-01-02")
		lastday := t.AddDate(0,0,-1).Format("2006-01-02")
		connect.Db.Debug().Where("time BETWEEN ? AND ?", yesterday,today).Find(&today_visit)
		connect.Db.Debug().Where("time BETWEEN ? AND ?", lastday,yesterday).Find(&yesterday_visit)
		connect.Db.Find(&all_visit)
		var rate float32
		if len(yesterday_visit) == 0{
			rate = 0
		}else{
			rate = float32(len(today_visit))/float32(len(yesterday_visit))*100 - 100
		}
		var login models.User_visit
		login = models.User_visit{
			All_visit: len(all_visit),
			Today_visit: len(today_visit),
			Yestoday_visit: len(yesterday_visit),
			Rate	: fmt.Sprintf("%.2f",rate),
		}
		return login

}


func Search_card_public_bykeyword(keyword string) []models.Card_warehouse {
	var card []models.Card_warehouse
	 connect.Db.Where("keyword like ?","%"+keyword+"%").Find(&card)
	return card
}

func Search_search_notes(name string, keyword string, num int, stime string, etime string) []models.Search_notes {
	var notes []models.Search_notes
	log.Printf("222"+stime)
	log.Printf("222"+etime)
	connect.Db.Where("time BETWEEN ? AND ? and name like ? and keyword like ?",stime,etime,"%"+name+"%","%"+keyword+"%").Find(&notes)
	return notes
}

func Search_search_note(account string,keyword string) []models.Search_notes {
	var notes []models.Search_notes
	connect.Db.Where("account = ? and keyword = ?",account,keyword).Find(&notes)
	return notes
}

func Search_search_note_account(account string) []models.Search_notes {
	var notes []models.Search_notes
	connect.Db.Where("account = ?",account).Find(&notes)
	return notes
}

type Keyword_search struct {
	Search_user int `json:"search_user"`
	Per_capita float32 `json:"per_capita"`
	Keyword []models.Search_keyword_num
}

func Search_search() Keyword_search {
	var notes []models.Search_notes
	var search_user []models.Search_notes
	var search_keyword []models.Search_keyword_num
	var keyword_search Keyword_search
	search_num := 0
	connect.Db.Find(&notes)
	for i := 0;i < len(notes);i++{
		search_num = search_num + notes[i].Num
	}
	connect.Db.Where("first_search = ?",1).Find(&search_user)
	connect.Db.Order("num desc").Find(&search_keyword)
	log.Println(search_keyword)
	per_capita := float32(search_num)/float32(len(search_user))
	log.Println(per_capita)
	keyword_search = Keyword_search{
		Search_user: len(search_user),
		Per_capita:         per_capita,
		Keyword:            search_keyword,
	}
	return keyword_search
}

func Search_bcnum() []models.Card_warehouse {
	/*var bcnum []models.Bc_num
	connect.Db.Find(&bcnum)
	return bcnum*/
	var results []models.Card_warehouse
	connect.Db.Debug().Where("card_browse > ?",0).Find(&results)
	return results
}

func Search_user_byadmin(admin string) models.User {
	var user models.User
	connect.Db.Where("admin_id = ?",admin).Find(&user)
	return user
}

func Search_role(role_id int) models.Role {
	var role models.Role
	connect.Db.Where("id = ?",role_id ).Find(&role)
	return role
}

func Search_meta(meta_id int) models.Meta {
	var meta models.Meta
	connect.Db.Where("meta_id = ?",meta_id).Find(&meta)
	return meta
}

func S2a(s string) []string{
	var a []string
	if s == ""{
		return a
	}else{
		a = strings.Split(s, ",")
		return a
	}
}

func Card_stirng2array(data []models.Card_warehouse) []models.Card {
	var res [100]models.Card
	result := data
	for i := 0; i < len(result); i++ {
		pic := S2a(result[i].Card_picture)
		log.Println(pic)
		video := S2a(result[i].Card_video)
		web_introduction_public := S2a(result[i].Web_introduction_public)
		web_picture_public := S2a(result[i].Web_picture_public)
		web_video_public := S2a(result[i].Web_video_public)
		//web_introduction := strings.Split(res[i].Web_introduction, ",")
		web_picture := S2a(result[i].Web_picture)
		web_video := S2a(result[i].Web_video)
		keyword := S2a(result[i].Keyword)
		res[i] = models.Card{
			Card_id:                 result[i].Card_id,
			Card_company_id:         result[i].Card_company_id,
			Card_user_account:       result[i].Card_user_account,
			Card_user_name:          result[i].Card_user_name,
			Card_company_name:       result[i].Card_company_name,
			Card_address:            result[i].Card_address,
			Card_phone:              result[i].Card_phone,
			Card_company_post:       result[i].Card_company_post,
			Card_email:              result[i].Card_email,
			Card_introduction:       result[i].Card_introduction,
			Card_picture:            pic,
			Card_video:              video,
			Card_privacy:            result[i].Card_privacy,
			Card_browse:             result[i].Card_browse,
			Card_collect:            result[i].Card_collect,
			Website_address:         result[i].Website_address,
			Card_mould_id:           result[i].Card_mould_id,
			Mark:                    result[i].Mark,
			Web_introduction:        result[i].Web_introduction,
			Web_picture:             web_picture,
			Web_video:               web_video,
			Web_introduction_public: web_introduction_public,
			Web_picture_public:      web_picture_public,
			Web_video_public:        web_video_public,
			Keyword:                 keyword,
			Avatar:                  result[i].Avatar,
		}
	}
	return res[0:len(result)]
}

func Search_keyword_num(keyword string) []models.Search_keyword_num {
	var num []models.Search_keyword_num
	connect.Db.Where("keyword = ?",keyword).Find(&num)
	return num
}

func Keyword_ranking(name string, keyword string, begin_date string, end_date string) []models.Search_notes {
	var search_notes []models.Search_notes
	connect.Db.Where("name like ? and keyword like ? and time >= ? and time < ?","%"+name+"%" , "%"+keyword+"%" , begin_date , end_date).Order("num desc").Find(&search_notes)
	return search_notes
}

func Search_share(card_id int, share_name string, visit_name string, begin_date string, end_date string) []models.Share {
	var share []models.Share
	if card_id == 0{
		connect.Db.Where("share_name like ? and visit_name like ? and time >= ? and time < ?",share_name, visit_name, begin_date, end_date).Find(&share)
	}else{
		connect.Db.Where("card_id = ?",card_id).Find(&share)
	}
	return share
}

func Search_keyword_pulic() []models.Keyword_warehouse {
	var keyword_pubilc []models.Keyword_warehouse
	connect.Db.Where("account = ?","admin").Find(&keyword_pubilc)
	return keyword_pubilc
}

func Search_keyword_byaccount(account string) []models.Search_notes {
	var keyword_my []models.Search_notes
	connect.Db.Where("account = ?",account).Order("num desc").Limit(3).Find(&keyword_my)
	log.Println(account)
	return keyword_my
}

func Search_share_bycardid(card_id int) []models.Share_2 {
	var share []models.Share
	connect.Db.Where("card_id = ? ",card_id).Find(&share)
	var share_2 [1000]models.Share_2
	for i := 0 ; i < len(share) ; i ++{
		user := Search_defcard_byuseraccount(share[i].Visit_account)
		avatar := user.Avatar
		name := user.Card_user_name
		time := share[i].Time.Format("2006-01-02 15:04:05")
		share_2[i] = models.Share_2{
			Card_id:       share[i].Card_id,
			Share_account: share[i].Share_account,
			Share_name:    share[i].Share_name,
			Visit_account: share[i].Visit_account,
			Visit_name:    name,
			Time:          time,
			Avatar:        avatar,
		}
	}
	return share_2[:len(share)]
}

//
func Md5Encode(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
