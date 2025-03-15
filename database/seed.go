package main

import (
	"github.com/Lunaticsatoshi/go-task/database/seeders"
	"github.com/Lunaticsatoshi/go-task/initializers"
)

func main() {
	initializers.MigrateDB()
	initializers.ConnectDB()
	seeders.UserSeeder(initializers.DB)
	seeders.TaskSeeder(initializers.DB)
}
