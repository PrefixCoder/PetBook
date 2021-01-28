package controllers

import (
	"fmt"
	"github.com/dpgolang/PetBook/pkg/logger"
	"github.com/dpgolang/PetBook/pkg/view"
	"net/http"
)

type Stats struct{
	BlogCount int
	UserCount int
	ForumCount int
}

func (c *Controller) GetCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results,err:= c.BlogStore.GetBlogStats()
		if err != nil {
			logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		countblog := results
		fmt.Println(countblog)

		view.GenerateHTML(w, countblog,"stats")
	}
}