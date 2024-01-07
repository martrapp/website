package pages

import (
	"html/template"
	"log"

	"github.com/labstack/echo/v4"

	"isabelroses.com/lib"
)

type BlogProps struct {
	Posts []lib.Post
}

func Blog(c echo.Context) error {
	props := BlogProps{
		Posts: lib.GetBlogPosts(),
	}

	templates := []string{
		"./templates/layouts/base.html",
		"./templates/partials/header.html",
		"./templates/partials/blogpreview.html",
		"./templates/pages/blog.html",
	}

	ts, err := template.ParseFiles(templates...)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return ts.ExecuteTemplate(c.Response().Writer, "base", props)
}
