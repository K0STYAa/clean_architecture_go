package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"forgoproject/services/contact/internal/useCase"
)

// @title slurm contact service on clean architecture
// @version 1.0
// @description contact service on clean architecture
// @license.name kolyadkons

// @contact.name API Support
// @contact.email kolyadkons@gmail.com

// @BasePath /

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("HTTP_PORT", 80)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
	viper.SetDefault("IS_PRODUCTION", "false")
}

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	router    *gin.Engine

	options Options
}

type Options struct{}

func New(ucContact useCase.Contact, ucGroup useCase.Group, options Options) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf("%s:%d", viper.GetString("HTTP_HOST"), uint16(viper.GetUint("HTTP_PORT"))))
}

func checkAuth(c *gin.Context) {
	c.Next()
}
