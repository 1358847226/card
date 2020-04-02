package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/database"
	_ "awesomeProject1/database"
	"awesomeProject1/middleware"
	"awesomeProject1/static"
	_ "awesomeProject1/static"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func main(){
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			list := database.Search_browse_byname("红")
			c.JSON(200, gin.H{
				"message": "pong",
				"list":list,
			})
		})

	r.GET("/test", func(c *gin.Context) {
		list := database.Search_browse_byname("红")
		c.JSON(200, gin.H{
			"message": "pong",
			"list":list,
		})
	})

	user := new(controller.User)
	r.Use(static.Cors())
	private := r.Group("/v1").Use(middleware.Validate())
	{

			private.POST("/searchuser",user.Search_user)//查询所有用户

			private.POST("/Search_user_byname",user.Search_user_byname)//根据名字模糊查询用户
			private.POST("/Search_card",user.Search_card)//查询所有名片
			private.POST("/Search_card_bycomname",user.Search_card_public_bycomname)//根据公司名模糊查询所有名片
			private.POST("/Search_chat_byuserid",user.Search_chat_byuserid)//根据单个用户id查询所有聊天记录
			private.POST("/Info",user.Info)
			private.POST("/Nav",user.Nav)
			private.POST("/Data_analysis_chart",user.Data_analysis_chart)//用户增长图表
			private.POST("/Data_analysis",user.Data_analysis)//用户增长信息
			private.POST("/Keyword_ranking",user.Keyword_ranking)//查询关键字排行
			private.POST("/Search_share",user.Search_share)
		}

	r.StaticFS("/upload", http.Dir("./upload"))
	r.GET("/user/:name", user.UserInfo)
	r.POST("/login", user.Login)//后台登录接口
	r.POST("/Userregister",user.Userregister)//第一次授权注册
	r.POST("/Userlogin",user.Userlogin)//用户登录
	r.POST("/createuser",user.Createuser)//新增用户接口
	r.POST("/Search_user_byaccount",user.Search_user_byaccount)//根据账号查询用户信息
	r.POST("/Search_card_byaccount",user.Search_card_byaccount)//根据账号查询名片
	r.POST("/Search_defcard_byaccount",user.Search_defcard_byaccount)//根据账号查询默认名片
	r.POST("/Updata_defcard",user.Updata_defcard)//修改默认名片
	r.POST("/Closecard",user.Closecard)//每次退出我的名片界面需要调用
	r.POST("/Search_card_byid",user.Search_card_byid)//根据名片id查询名片
	r.POST("/Search_card_public_bykeyword",user.Search_card_public_bykeyword)//根据名片关键字查询名片
	r.POST("/createcard",user.Createcard)//新增名片
	//r.POST("/Createbrowse",user.Createbrowse)//新增浏览记录
	r.POST("/Updatainformation",user.Updatainformation)//新增名片详细信息
	r.POST("/Createchat",user.Createchat)//新增聊天信息
	r.POST("/Createcollect",user.Createcollect)//收藏，取消收藏
	r.POST("/Createkeyword",user.Createkeyword)//创建标签
	r.POST("/Createneed",user.Createneed)//创建需求
	r.POST("/Search_card_public",user.Search_card_public)//查询所有公开的名片
	r.POST("/Deletecard_bycardid",user.Deletecard_bycardid)//根据cardid删除名片
	r.POST("/Upload",user.Upload)//上传视频
	r.POST("/Deletevideo",user.Deletevideo)//删除视频
	r.POST("/Search_card_public_bycomname",user.Search_card_public_bycomname)//根据公司名模糊查询公开的名片
	r.POST("/Createmould",user.Createmould)//新增名片模板
	r.POST("/Searchmould",user.Searchmould)//根据模板id查询模板信息
	r.POST("/Updata_card",user.Updata_card)//修改名片
	r.POST("/Md5Encode",user.Md5Encode)//测试md5加密
	r.POST("/Encode",user.Encode)//两次base64加密
	r.POST("/Decode",user.Decode)//两次base64解密
	r.POST("/Share",user.Share)//新增分享记录
	r.POST("/Search_collect_byaccount",user.Search_collect_byaccount)//根据用户账号查询收藏名片
	r.POST("/Search_keywords_byaccount",user.Search_keywords_byaccount)//关键字检索
	r.POST("/Code",user.Code)//生成二维码
	r.POST("/Regular",user.Regular)//正则表达式测试
	r.POST("/Search_share_bycardid",user.Search_share_bycardid)//根据名片id查询分享记录
	r.POST("/SubscribeMessage",user.SubscribeMessage)


	r.POST("/Updatapassword",user.Updatapassword)//修改管理员密码
	r.POST("/Search_notes",user.Search_notes)//查询搜索记录
	//r.Run("192.168.16.126:8080") // 监听并在 0.0.0.0:8080 上启动服务
	r.Run("172.17.205.23:8080")//在服务器上运行





		/*data := models.User{
		Name:     "小红",
		Password: "666",
		Power:    0,
		Share:    0,
	}*/
	//database.Createuser(data)
	//r.POST("/createuser",user.Createuser)
	//插入用户信息


/*	data := models.Card_warehouse{
		Card_company_id:   0,
		Card_user_id:      24,
		Card_user_name:    "小李",
		Card_company_name: "AA模塑",
		Card_address:      "黄岩新前",
		Card_phone:        "154825698521",
		Card_company_post: "CEO",
		Card_email:        "454789@qq.com",
		Card_introduction: "大家好，我是小李",
		Card_picture:      "",
		Card_video:        "",
		Card_privacy:      false,
		Card_browse:       0,
		Card_collect:      0,
		Website_address:   "",
		Card_mould_id:     0,
	}
	database.Createcard(data)*/
	 //新增名片

	 /*data := models.BrowseRecord{

		 Userid:   23,
		 Username: "小王",
		 Cardid:   2,
		 Date:     time.Now(),
	 }
	 database.Createbrowse(data)*/
	 //新增浏览记录

	/* data := models.Card_information{
	 	 Card_information_type:    0,
	 	 Card_information_card_id: 5,
	 	 Card_information_cont:    "这是小李的名片3",
	 	 Card_information_picture: "",
	 	 Card_information_video:   "",
	 	 Card_information_pdf:     "",
	  }
	 database.Createinformation(data)*/
	 //新增名片信息

	/*data := models.Chat{
		Chat_userA_id:   25,
		Chat_userB_id:   24,
		Chat_userA_name: "小红",
		Chat_userB_name: "小李",
		Chat_content:    "小红说，李总好",
		Chat_date:       time.Now(),
	}
	database.Createchat(data)*/
	//新增聊天信息

	/*data := models.Collect_record{

		User_id:      24,
		User_name:    "小李",
		Card_id:      1,
		Collect_mark: true,
		Collect_date: time.Now(),
	}
	database.Createcollect(data)*/
	//新增收藏信息

	/*data := models.Company_warehouse{
		Company_name:    "BB金属",
		Company_address: "黄岩BBBB",
		Company_phone:   "16485254874",
	}
	database.Createcompany(data)*/
	//新增公司基本信息

	/*data := models.Keyword_warehouse{
		Keyword_content: "螺帽",
	}
	database.Createkeyword(data)*/
	//新增标签

	/*data := models.Need_warehouse{
		Need_user_id:   24,
		Need_user_name: "小李",
		Need_name:      "螺帽",
		Need_content:   "我需要结实的螺帽",
		Need_keyword:   "螺帽",
		Need_starttime: time.Now(),
		Need_endtime:   time.Now(),
	}
	database.Createneed(data)*/
	//新增需求












	//fmt.Println(database.Search_user())
	//查询所有用户

	//fmt.Println(database.Search_user_byid(25))
	//根据id查询用户

	//fmt.Println(database.Search_user_byname("张"))
	//根据姓名模糊查询

	//fmt.Println(database.Search_user_bypower(0))
	//根据用户权限查询


	//fmt.Println(database.Search_card())
	//查询所有名片库中的名片

	//fmt.Println(database.Search_card_public())
	//查询名片库中公开的名片

	//fmt.Println(database.Search_card_public_bycomid(1))
	//根据公司id查询公开的名片

	//fmt.Println(database.Search_card_bycomid(1))
	//根据公司id查询所有名片

	//fmt.Println(database.Search_card_bycomname("塑"))
	//根据公司名模糊查询所有名片

	//fmt.Println(database.Search_card_public_bycomname("天"))
	//根据公司名模糊查询所有公开的名片

	//fmt.Println(database.Search_card_public_byname("红"))
	//根据姓名模糊查询公开的名片

	//fmt.Println(database.Search_card_byname("红"))
	//根据姓名模糊查询所有名片


	//fmt.Println(database.Search_browse_byname("红"))
	//根据姓名模糊查询浏览记录

	//fmt.Println(database.Search_browse_bycardid(2))
	//根据名片id搜索浏览记录


	//fmt.Println(database.Search_cardmould_bymouldid(1))
	//根据企业名片模板id查询企业名片模板

	//fmt.Println(database.Search_cardmould_bycomname("天"))
	//根据公司名模糊查询企业名片模板

	//fmt.Println(database.Search_chat_byuserid(24))
	//根据单个用户id查询聊天记录

	//fmt.Println(database.Search_chat_byuserids(24,23))
	//根据聊天双方id查询聊天记录

	//fmt.Println(database.Search_chat_byusernames("小","李"))
	//根据聊天双方的姓名模糊查询聊天记录


	//fmt.Println(database.Search_collect_byuserid(24))
	//根据用户id查询收藏记录

	//fmt.Println(database.Search_collect_byusername("李"))
	//根据用户名模糊查询收藏记录

	//fmt.Println(database.Search_company_bycomid(1))
	//根据公司id查询公司信息

	//fmt.Println(database.Search_company_byname("A"))
	//根据公司名模糊查询公司信息

	//fmt.Println(database.Search_keyword())
	//查询所有标签

	//fmt.Println(database.Search_keyword_bycon("螺帽"))
	//根据标签内容查询标签

	//fmt.Println(database.Search_need_byid(20))
	//根据用户id查询需求

	//fmt.Println(database.Search_need_byusername("张"))
	//根据用户姓名模糊查询需求

	//fmt.Println(database.Search_need_bykeyword("灯","车","车灯"))
	//根据关键字查询需求

	//fmt.Println(database.Search_product_informaion_bytypePid(1,1))
	//根据信息类型和产品id查询产品详细信息

	//fmt.Println(database.Search_product_informaion_bytypePname(1,"车"))
	//根据信息类型和产品名模糊搜索产品详细信息

	//fmt.Println(database.Search_product_byid(1))
	//根据产品id查询产品基本信息

	//fmt.Println(database.Search_product_bycomid(2))
	//根据公司id查询产品基本信息

	//fmt.Println(database.Search_product_bycomname("AA"))
	//根据公司名模糊搜索产品基本信息

	//fmt.Println(database.Search_product_byPname("车"))
	//根据产品名模糊搜索产品基本信息

	//fmt.Println(database.Search_product_bykeyword("车灯"))
	//根据关键字搜索产品基本信息

	//fmt.Println(database.Login(20,"123"))
	//登录验证
























	//database.Delete(21)
	//根据id删除用户

	//database.Deletecard_bycardid(5)
	//根据card_id删除名片










	/*
	data :=models.User{
		Name 		: "小王",
		Password 	: "999",
		Power		: 2,
		Share		: 2,

	}
	database.Update_user(data,23)
	 */
	//根据用户id更新用户信息
}

