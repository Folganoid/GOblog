package routes

import (
	"../db/documents"
	"../models"
	"../utils"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
	"net/http"
)

func WriteHandler(rnd render.Render) {
	post := models.Post{}
	rnd.HTML(200, "write", post)
}

func EditHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {

	postsCollection := db.C("posts")

	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	rnd.HTML(200, "write", post)
}

func SavePostHandler(rnd render.Render, r *http.Request, db *mgo.Database) {

	postsCollection := db.C("posts")

	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkDownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{id, title, contentHtml, contentMarkdown}
	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/")
}

func DeleteHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database) {

	postsCollection := db.C("posts")

	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkDownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}
