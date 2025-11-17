package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	GenerateTestData(db)
	posts := findPostAndCommentByUserId(1, db)

	fmt.Println(posts)

	p2 := test02()
	fmt.Println(p2)

	delete01(8, db)
}

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment"`
	Name      string
	Age       uint
	Post      []Post `gorm:"foreignKey:UserID"`
	PostCount uint
}

type Post struct {
	ID       uint64 `gorm:"primary_key:auto_increment"`
	Content  string
	UserID   uint64
	User     User      `gorm:"foreignKey:UserID"` // 帖子属于用户
	Comments []Comment `gorm:"foreignKey:PostID"`
	Status   string
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {

	var user User
	err1 := tx.First(&user, "id = ?", post.UserID).Error
	if err1 != nil {
		return err1
	}

	user.PostCount++

	err2 := tx.Save(&user).Error

	if err2 != nil {
		return err2
	}

	return nil
}

type Comment struct {
	ID      uint64 `gorm:"primary_key:auto_increment"`
	Comment string
	PostID  uint64
	Post    Post `gorm:"foreignKey:PostID"` // 评论属于帖子
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {

	var post Post
	err1 := tx.Preload("Comments").Take(&post, "id = ?", c.PostID).Error
	if err1 != nil {
		panic(err1)
	}

	if len(post.Comments) == 1 {
		err2 := tx.Model(&Post{}).Where("id = ?", post.ID).Update("status", "无评论").Error
		fmt.Println("执行 Comment AfterDelete")
		if err2 != nil {
			panic(err2)
		}
	}

	return nil
}

var db *gorm.DB

func init() {
	db1, err := gorm.Open(mysql.Open("root:szoscar55@tcp(10.0.47.57:3306)/test001?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = db1

	//表迁移
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	//数据初始化
	//GenerateTestData(db)

}

func findPostAndCommentByUserId(userID uint64, db *gorm.DB) []Post {

	var user User

	err := db.Debug().Preload("Post.Comments").Take(&user, "id = ?", userID).Error

	if err != nil {
		fmt.Println(err)
	}

	return user.Post

}

func test02() Post {

	var post Post
	err := db.Debug().Preload("Comments").Select("posts.id, count(1) as count").Joins("left join comments ON comments.post_id = posts.id").Group("posts.id").Order("count desc").First(&post)

	if err != nil {
		fmt.Println(err)
	}

	return post
}

func delete01(commentId uint64, db *gorm.DB) {

	var comment Comment
	if err := db.First(&comment, commentId).Error; err != nil {
		fmt.Printf("   评论 %d 不存在\n", commentId)
		return
	}

	err := db.Debug().Delete(&comment).Error
	if err != nil {
		fmt.Println(err)
	}

}

// 生成测试数据
func GenerateTestData(db *gorm.DB) error {
	// 创建用户
	users := []User{
		{Name: "张三", Age: 25},
		{Name: "李四", Age: 30},
		{Name: "王五", Age: 28},
		{Name: "赵六", Age: 35},
		{Name: "孙七", Age: 22},
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}

	// 创建帖子
	posts := []Post{
		// 用户1的帖子
		{Content: "今天天气真好！", UserID: users[0].ID},
		{Content: "学习Golang真有趣", UserID: users[0].ID},

		// 用户2的帖子
		{Content: "分享一个技术文章", UserID: users[1].ID},
		{Content: "周末去哪里玩？", UserID: users[1].ID},

		// 用户3的帖子
		{Content: "我的新项目开始了", UserID: users[2].ID},

		// 用户4的帖子
		{Content: "读书笔记分享", UserID: users[3].ID},
		{Content: "电影推荐", UserID: users[3].ID},

		// 用户5的帖子
		{Content: "健身打卡第一天", UserID: users[4].ID},
	}

	if err := db.Create(&posts).Error; err != nil {
		return err
	}

	// 创建评论
	comments := []Comment{
		// 帖子1的评论
		{Comment: "确实不错！", PostID: posts[0].ID},
		{Comment: "我也觉得很好", PostID: posts[0].ID},

		// 帖子2的评论
		{Comment: "Golang确实很棒", PostID: posts[1].ID},
		{Comment: "学习了，谢谢分享", PostID: posts[1].ID},

		// 帖子3的评论
		{Comment: "文章很有深度", PostID: posts[2].ID},

		// 帖子4的评论
		{Comment: "建议去爬山", PostID: posts[3].ID},
		{Comment: "电影院不错", PostID: posts[3].ID},

		// 帖子5的评论
		{Comment: "项目加油！", PostID: posts[4].ID},

		// 帖子6的评论
		{Comment: "这本书我也看过", PostID: posts[5].ID},

		// 帖子7的评论
		{Comment: "什么电影？", PostID: posts[6].ID},

		// 帖子8的评论
		{Comment: "一起健身！", PostID: posts[7].ID},
		{Comment: "坚持就是胜利", PostID: posts[7].ID},
	}

	if err := db.Create(&comments).Error; err != nil {
		return err
	}

	return nil
}
