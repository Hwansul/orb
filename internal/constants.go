package internal

import (
	"os"

	"github.com/adrg/xdg"
)

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = xdg.ConfigHome + "/but"

// github username and repo in order to fetch snippets.
var Owner = "hoehwa"
var Repository = "sandbox"

// a url for my private websites.
const GardenBaseurl = "https://mindulle.github.io/garden"

// urls for origin contents.
const WebdevBaseurl = "https://web.dev"

// urls for sandbox providers.
const FiddleBaseurl = "https://jsfiddle.net/gh/get/library/pure/"
const ComponentBaseurl = "https://stackblitz.com/github/"
