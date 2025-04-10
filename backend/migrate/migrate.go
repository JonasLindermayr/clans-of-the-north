package migrate

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func init() {
	utils.LoadEnvs();
	utils.ConnectDB();
}

func main() {
	utils.DB.AutoMigrate(&models.User{})
}