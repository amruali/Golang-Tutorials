package main

import (
	_ "database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"time"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "a12345"
	dbname   = "postgres"
	searchPath  = "TEMP"
)


type Profile struct {
	ProfileID     uint              `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	FirstName     string            `sql:"type:VARCHAR(55);NOT NULL"`
	LastName      string            `sql:"type:VARCHAR(55);NOT NULL"`
	UserID        uint              `sql:"FOREIGN_KEY;UNIQUE;NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type User struct {
	UserID        uint              `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserName      string            `gorm:"type:VARCHAR(55);UNIQUE;NOT NULL"`
	Password      string            `gorm:"type:VARCHAR(55);NOT NULL"`
	Email         string            `gorm:"type:VARCHAR(55);UNIQUE;NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Profile       Profile           `gorm:"foreign_key:UserID"`
	Posts         []Post            `gorm:"foreign_key:UserID"`
	PostLikes     []PostLikes           `gorm:"foreign_key:UserID"`
	PostComments  []PostComments        `gorm:"foreign_key:UserID"`
	PostCommentLikes []PostCommentLikes  `gorm:"foreign_key:UserID"`
	PostCommentReplies []PostCommentReplies   `gorm:"foreign_key:UserID"`
	PostCommentRepliesLikes[]PostCommentRepliesLikes  `gorm:"foreign_key:UserID"`
	PostShares             []PostShares                 `gorm:"foreign_key:UserID"`
}

type Post struct {
	PostID        uint              `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Description   string            `sql:"type:VARCHAR(400);NOT NULL"`
	UserID        uint              `sql:"NOT NULL"`
	PostLikes     []PostLikes
	PostComments  []PostComments
	CreatedAt     time.Time
	UpdatedAt     time.Time
}


//UNIQUE Group Constraint
type PostLikes struct {
	PostLikeID    uint              `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserID        uint              `sql:"NOT NULL"`
	PostID        uint              `sql:"NOT NULL"`
	CreatedAt     time.Time
}

type PostComments struct {
	PostCommentID  			uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Description             string    `sql:"type:VARCHAR(400);NOT NULL"`
	UserID        		 	uint      `sql:"NOT NULL;"`
	PostID        		    uint      `sql:"NOT NULL"`
	PostCommentLikes        []PostCommentLikes
	PostCommentReplies      []PostCommentReplies
}

type PostCommentLikes struct {
	PostCommentLikeID       uint          `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	CreatedAt               time.Time
	UserID                  uint          `sql:"NOT NULL"`
	PostCommentID           uint          `sql:"NOT NULL"`
}

type PostCommentReplies struct {
	PostCommentReplyID      uint        `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Description             string      `sql:"type:VARCHAR(400);NOT NULL"`
    UserID                  uint        `sql:"NOT NULL"`
	PostCommentID           uint        `sql:"NOT NULL"`
	PostCommentRepliesLikes []PostCommentRepliesLikes
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type PostCommentRepliesLikes struct {
	PostCommentRepliesLikeID uint          `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserID                   uint          `sql:"NOT NULL"`
	PostCommentReplyID       uint          `sql:"NOT NULL"`
	CreatedAt                time.Time
}

type PostShares struct {
	PostShareID               uint           `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Description               string         `sql:"type:VARCHAR(400);NOT NULL"`
	UserID                    uint           `sql:"NOT NULL "`
	PostID                    uint           `sql:"NOT NULL"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

type Follow struct {

}






func main () {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, searchPath)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec(`set search_path='TEMP'`)

	db.AutoMigrate( &Profile{} , &User{} , &Post{} ,  &PostLikes{} ,  &PostComments{} ,
	                &PostShares{} , &PostCommentLikes{} , &PostCommentReplies{} ,
	                &PostCommentRepliesLikes{})


	// ProfileForeignKeys
	db.Model(&Profile{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")

	// PostForeignKeys
	db.Model(&Post{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")

	// PostLikesForeignKeys
	db.Model(&PostLikes{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostLikes{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT")


	// PostCommentsForeignKeys
	db.Model(&PostComments{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostComments{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT")

	//PostSharesForeignKeys
	db.Model(&PostShares{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostShares{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT")


	// PostCommentLikesForeignKeys
	db.Model(&PostCommentLikes{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostCommentLikes{}).AddForeignKey("post_comment_id", "post_comments(post_comment_id)", "RESTRICT", "RESTRICT")

	// PostCommentRepliesForeignKeys
	db.Model(&PostCommentReplies{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostCommentReplies{}).AddForeignKey("post_comment_id", "post_comments(post_comment_id)", "RESTRICT", "RESTRICT")


	// PostCommentRepliesLikesForeignKeys
	db.Model(&PostCommentRepliesLikes{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&PostCommentRepliesLikes{}).AddForeignKey("post_comment_reply_id", "post_comment_replies(post_comment_reply_id)", "RESTRICT", "RESTRICT")


	db.AutoMigrate( &Profile{} , &User{} , &Post{} ,  &PostLikes{} ,  &PostComments{} ,
		&PostShares{} , &PostCommentLikes{} , &PostCommentReplies{} ,
		&PostCommentRepliesLikes{})
}