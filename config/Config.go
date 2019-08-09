package config

type AppConfig struct{
	DbHost   string `split_words:"true"`
	DbPort   string `default:"3306" split_words:"true"`
	DbUser   string `split_words:"true"`
	DbPass   string `split_words:"true"`
	DbName   string `default:"dockeraja" split_words:"true"`
	AppsPort string `default:"1323" split_words:"true"`
}
