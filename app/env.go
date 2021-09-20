package app

import (
	"github.com/handrixn/example-repo/helper"
	"os"
	"strings"
)

func LoadEnvFile(pathFile string) {
	files, err := os.ReadFile(pathFile)
	helper.PanicIfError(err)

	env := strings.Split(string(files), "\n")

	for _, v := range env {
		sliceKeyValueEnv := strings.Split(v, "=")
		err := os.Setenv(sliceKeyValueEnv[0], sliceKeyValueEnv[1])
		helper.PanicIfError(err)
	}
}
