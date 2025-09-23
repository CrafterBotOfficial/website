package main

import "html/template"

type PageVars interface { }

type IndexPageData struct {
	ServerTime string
	Messages []string
	Socials []Social
}

type ArticlePostPageData struct {
	Title string
	Author string
	Content template.HTML 
}

type ArticlesPageData struct {
	Blogs []Blog
}
