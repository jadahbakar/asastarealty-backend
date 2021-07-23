package utils

import (
	"log"

	"github.com/jadahbakar/asastarealty-backend/pkg/version"
)

func AppName() {
	appName := `
 	 __ _ ____    ____  __   ___ __ _ ____ __ _ ____
 	(  / |  _ \__(  _ \/ _\ / __|  / |  __|  ( (    \
	 )  ( )   (___) _ (    ( (__ )  ( ) _)/    /) D (
	(__\_|__\_)  (____|_/\_/\___|__\_|____)_)__|____/
	`
	log.Println(appName + version.Version)
}
