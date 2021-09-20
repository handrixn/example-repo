package app

import (
	"os"
	"strings"

	"github.com/handrixn/example-repo/helper"
)

func LoadEnvFile(pathFile string) {
	file, err := os.ReadFile(pathFile)
	helper.PanicIfError(err)

	env := strings.Split(string(file), "\n")

	for _, v := range env {
		sliceKeyValueEnv := strings.Split(v, "=")
		err := os.Setenv(sliceKeyValueEnv[0], sliceKeyValueEnv[1])
		helper.PanicIfError(err)
	}
}
