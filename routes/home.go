package routes

import (
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
	"net/http"
)

func IndexHandler(rnd render.Render, r *http.Request, db *mgo.Database) {

	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(inMemorySession.Get(cookie.Value))
	}

	postDocuments := []documents.PostDocument{}
	postCollection := db.C("Posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	rnd.HTML(200, "index", posts)
}
