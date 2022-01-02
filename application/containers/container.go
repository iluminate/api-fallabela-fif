package containers

import (
	"api-fallabela-fif/application/controllers"
	"api-fallabela-fif/application/repositories"
	"api-fallabela-fif/application/services"
	"api-fallabela-fif/application/utils"
	"api-fallabela-fif/helpers/database"
	"log"
)

func mongoConfig() map[string]string {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return map[string]string{
		"uri":      config.Mongodb.Uri,
		"database": config.Mongodb.Database,
	}
}

func MongodbHelper() *database.MongodbHelper {
	mongoHelper := database.NewMongodbHelper(mongoConfig())
	err := mongoHelper.Open()
	if err != nil {
		log.Fatalln("cannot connect to mongo:", err)
	}
	return mongoHelper
}

func BeerRepository() repositories.IBeerRepository {
	return repositories.NewBeerRepository(MongodbHelper())
}

func BeerService() services.IBeerService {
	return services.NewBeerService(BeerRepository())
}
func BeerHandler() *controllers.BeerHandler {
	return controllers.NewBeerHandler(BeerService(), ExchangeService())
}

func ExchangeService() services.IExchangeService {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	url := config.Currency.Url
	endpoints := map[string]string{"live": config.Currency.Endpoints.Live}
	token := config.Currency.ApiKey
	return services.NewExchangeService(url, endpoints, token)
}
