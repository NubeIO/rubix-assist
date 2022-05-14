package cmd

import (
	bac "github.com/NubeDev/bacnet/cmd/cmd"
)

func init() {
	//import bacnet commands
	bac.Execute()
}
