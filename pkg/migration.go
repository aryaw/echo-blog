package pkg

import (
	gmEngine "github.com/ShkrutDenis/go-migrations/engine"
	gmStore "github.com/ShkrutDenis/go-migrations/engine/store"
)

func RunMigration() {
	e := gmEngine.NewEngine()
	e.Run(GetMigrationsList())
}

func GetMigrationsList() []gmStore.Migratable {
	return []gmStore.Migratable{
		// &list.CreateUserTable{},
		// &list.CreatePostTable{},
		// &list.UpdateUserTable{},
	}
}
