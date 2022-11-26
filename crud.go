package star

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// 用户表
type User struct {
	Id        int
	Name      string
	Signature string
	Rank      int
	WaterNum  int
}

// 帖子表
type Post struct {
	Id      int
	Content string
	UserId  int
}

// 楼层表
type Floor struct {
	Id      int
	Content string
	UserId  int
	PostId  int
}
type PostAndUser struct {
	PostId    int
	Content   string
	UserName  string
	Signature string
}

func ProcessError(err error) {
	if err != nil {
		panic("failed")
	}
}

// 数据库名字是test，三张表都在test数据库下
func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		fmt.Println("open db failed err is ", err)
		return nil, err
	}
	return db, nil
}

// 创建一个帖子
func CreatePost(post Post) error {
	db, err := Init()
	if err != nil {
		fmt.Println("CreatePost failed err is", err)
		return err
	}
	sql := "insert into Post values (?,?,?)"

	_, err = db.Exec(sql, post.Id, post.Content, post.UserId)
	if err != nil {
		fmt.Println("db.Exec failed err is", err)
		return err
	}
	return nil
}

// 根据Id查询帖子
func SerchPost(postId int) (*Post, error) {
	db, err := Init()
	if err != nil {
		fmt.Println("SerchPost failed err is", err)
		return nil, err
	}
	sql := "select *from Post where id=?"

	var post Post
	err = db.QueryRow(sql, postId).Scan(&post.Id, &post.Content, &post.UserId)

	if err != nil {
		fmt.Println("SerchPost failed err is", err)
		return nil, err
	}

	return &post, nil
}

// 删除一条帖子
func DeletePost(postId int) error {
	db, err := Init()
	if err != nil {
		fmt.Println("DeletePost failed err is", err)
		return err
	}
	sql := "delete from Post where id=? "

	_, err = db.Exec(sql, postId)
	if err != nil {
		fmt.Println("db.Exec failed err is", err)
		return err
	}
	return nil
}

// 修改一个帖子的内容
func UpdatePost(postId int, content string) error {
	db, err := Init()
	if err != nil {
		fmt.Println("UpdatePost failed err is", err)
		return err
	}
	sql := "update Post set content=? where id=?"

	_, err = db.Exec(sql, content, postId)
	if err != nil {
		fmt.Println("db.Exec failed err is", err)
		return err
	}
	return nil
}

// 根据一个postId可以查询到该post和与之对应的用户的签名
func SearchPostAndUser(postId int) (*PostAndUser, error) {
	db, err := Init()
	if err != nil {
		fmt.Println("SearchPostAndUser failed err is", err)
		return nil, err
	}
	sql := `select p.id,p.content,u.name ,u.signature 
			from Post p 
			inner join User u on p.user_id = u.id 
			where p.id = ?`
	// 包含了post和user的信息
	var ans PostAndUser
	err = db.QueryRow(sql, postId).Scan(&ans.PostId, &ans.Content, &ans.UserName, &ans.Signature)
	if err != nil {
		fmt.Println("SearchPostAndUser failed err is", err)
		return nil, err
	}
	return &ans, nil
}

// 路由我还不是很熟悉，现在就是加一个http服务需要静态设置一个路由，所以就先实现了CreatePost,其他几个操作类似
// 前端的消息是json编码格式
func HttpCreatePost(w http.ResponseWriter, r *http.Request) {
	// 从前端传来的request中获取body，也就是一条post
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)

	if err := CreatePost(post); err != nil {
		fmt.Println("CreatePost failed err is", err)
		w.WriteHeader(500)
		fmt.Fprintln(w, "CreatePost failed")
		return
	}
	fmt.Fprintln(w, "CreatePost successed")
}

// 查找post和对应的用户名字和签名
func HttpSearchPostAndUser(w http.ResponseWriter, r *http.Request) {
	// 从前端传来的request中获取postId
	r.ParseForm()
	postId, _ := strconv.Atoi(r.Form.Get("id"))

	postAndUser, err := SearchPostAndUser(postId)

	if err != nil {
		fmt.Println("HttpSearchPostAndUser failed err is ", err)
		w.WriteHeader(500)
		fmt.Fprintln(w, "HttpSearchPostAndUser failed")
		return
	}

	// 以json形式给前端
	postAndUserJson, _ := json.Marshal(*postAndUser)
	w.Header().Set("Content-Type", "application/json")
	w.Write(postAndUserJson)
}
