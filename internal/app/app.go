package app

import (
	"fmt"

	"github.com/YurcheuskiRadzivon/telemetrics-back/config"
)

func Run(cfg *config.Config) {
	fmt.Println(cfg.App.Name)
}
