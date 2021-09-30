package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	create          = kingpin.Command("create", "Create new CronJob.")
	user            = create.Flag("user", "User Id for Kubernetes").Default("admin").String()
	pw              = create.Flag("password", "Password for Kubernetes").Required().String()
	url             = create.Flag("url", "API Url").Required().String()
	namespace       = create.Flag("ns", "Namespace").Default("default").String()
	schedule        = create.Flag("schedule", "Schedule").Required().String()
	imagepullpolicy = create.Flag("imagepullpolicy", "Image Pull Policy").Default("Always").String()
	cronJobName     = create.Flag("name", "CronJob Name").Required().String()
	image           = create.Flag("image", "Docker Image").Required().String()
	imagepullsecret = create.Flag("imagepullsecret", "Image Pull Secret").Required().String()
	secret          = create.Flag("secret", "Set secret name and mountpath").StringMap()
	env             = create.Flag("env", "Add env variable and it's value").StringMap()
	limit           = create.Flag("limit", "Add memory limit").String()
	request         = create.Flag("request", "Add memory request").String()

	patch            = kingpin.Command("patch", "Patch a current CronJob.")
	userP            = patch.Flag("user", "User Id for Kubernetes").Default("admin").String()
	pwP              = patch.Flag("password", "Password for Kubernetes").Required().String()
	urlP             = patch.Flag("url", "API Url").Required().String()
	namespaceP       = patch.Flag("ns", "Namespace").Default("default").String()
	scheduleP        = patch.Flag("schedule", "Schedule").Required().String()
	imagepullpolicyP = patch.Flag("imagepullpolicy", "Image Pull Policy").Default("Always").String()
	cronJobNameP     = patch.Flag("name", "CronJob Name").Required().String()
	imageP           = patch.Flag("image", "Docker Image").Required().String()
	imagepullsecretP = patch.Flag("imagepullsecret", "Image Pull Secret").Required().String()
	secretP          = patch.Flag("secret", "Set secret name and mountpath").StringMap()
	envP             = patch.Flag("env", "Add env variable and it's value").StringMap()
	limitP           = patch.Flag("limit", "Add memory limit").String()
	requestP         = patch.Flag("request", "Add memory request").String()

	del          = kingpin.Command("delete", "Delete a current CronJob.")
	userD        = del.Flag("user", "User Id for Kubernetes").Default("admin").String()
	pwD          = del.Flag("password", "Password for Kubernetes").Required().String()
	urlD         = del.Flag("url", "API Url").Required().String()
	namespaceD   = del.Flag("ns", "Namespace").Default("default").String()
	cronJobNameD = del.Flag("name", "CronJob Name").Required().String()
)

func main() {
	code := 0
	switch kingpin.Parse() {
	case "create":
		code = createCronJob()
	case "delete":
		code = deleteCronJob()
	case "patch":
		code = patchCronJob()
	}
	os.Exit(code)
}
