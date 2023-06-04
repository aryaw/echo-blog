package migrations

import (
	s "blog/server"
	"fmt"

	"blog/pkg/post"
)

func RunMigrations(server *s.Server) {
	err := server.DB.Debug().AutoMigrate(
		// post
		&post.ModelPost{},
		&post.ModelPostCategory{},
		&post.ModelPostToCategory{},
	)
	if err != nil {
		fmt.Println(err)
	}
}
