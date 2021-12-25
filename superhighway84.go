package main

import (
	"embed"
	"time"

	"log"
	"os"

	"github.com/mrusme/superhighway84/models"
	"github.com/mrusme/superhighway84/tui"
)

//go:embed superhighway84.jpeg
var EMBEDFS embed.FS

func main() {
  // ctx, cancel := context.WithCancel(context.Background())
  // defer cancel()

  dbInit := false
  dbInitValue := os.Getenv("SUPERHIGHWAY84_DB_INIT")
  if dbInitValue == "1" {
    dbInit = true
  }

  dbURI := os.Getenv("SUPERHIGHWAY84_DB_URI")
  if dbInit == false && dbURI == "" {
    log.Panicln("SUPERHIGHWAY84_DB_URI missing!")
  }

  dbCache := os.Getenv("SUPERHIGHWAY84_DB_CACHE")
  if dbCache == "" {
    log.Panicln("SUPERHIGHWAY84_DB_CACHE missing!")
  }

  // logger, err := zap.NewDevelopment()
  // if err != nil {
  //   log.Panicln(err)
  // }
  //
  //
  // db, err := database.NewDatabase(ctx, dbURI, dbCache, dbInit, logger)
  // if err != nil {
  //   log.Panicln(err)
  // }
  // defer db.Disconnect()
  // db.Connect(func() {
  //   //TUI.App.Stop()
  // })
  //
  // articles, _ := db.ListArticles()
  var articles []models.Article

  go func() {
    for i := 0; i < 10; i++ {
      time.Sleep(time.Second * 3)
      art1 := *models.NewArticle()
      art1.Subject = "This is a test"
      art1.Body = "This is just a test article\nWhat's up there?"
      art1.From = "test@example.com"
      art1.Newsgroup = "comp.alt.test"

      articles = append(articles, art1)
    }
  }()

  TUI := tui.Init(&EMBEDFS, &articles)

  go func() {
    time.Sleep(time.Second * 2)
    TUI.SetView("mainscreen")
    TUI.Refresh()
  }()
  TUI.Launch()

  // var input string
  // for {
  //   fmt.Scanln(&input)
  //
  //   switch input {
  //   case "q":
  //     return
  //   case "g":
  //     fmt.Scanln(&input)
  //     article, err := db.GetArticleByID(input)
  //     if err != nil {
  //       log.Println(err)
  //     } else {
  //       log.Println(article)
  //     }
  //   case "p":
  //     article := models.NewArticle()
  //     article.From = "test@example.com"
  //     article.Newsgroup = "comp.test"
  //     article.Subject = "This is a test!"
  //     article.Body = "Hey there, this is a test!"
  //
  //     err = db.SubmitArticle(article)
  //     if err != nil {
  //       log.Println(err)
  //     } else {
  //       log.Println(article)
  //     }
  //   case "l":
  //     articles, err := db.ListArticles()
  //     if err != nil {
  //       log.Println(err)
  //     } else {
  //       log.Println(articles)
  //     }
  //   }
  //
  // }
}

