package main

import (
	"fmt"
	"strings"
)

func deleteCronJob() int {
	send(*urlD+"/apis/batch/v1beta1/namespaces/"+*namespaceD+"/cronjobs/"+*cronJobNameD, *userD, *pwD, []byte(""), "DELETE")
	list, err := send(*urlD+"/apis/batch/v1/namespaces/"+*namespaceD+"/jobs", *userD, *pwD, []byte(""), "GET")

	if err != nil {
		return -2
	}
	var name string
	found := false
	if list == nil || list["items"] == nil {
		return -2
	}
	toMatch := fmt.Sprintf("%s-", *cronJobNameD)
	jobs := list["items"].([]interface{})
	for _, job := range jobs {
		meta := job.(map[string]interface{})["metadata"].(map[string]interface{})
		labels := meta["labels"].(map[string]interface{})
		app := labels["job-name"].(string)
		if strings.HasPrefix(app, toMatch) {
			name = meta["name"].(string)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Job for " + *cronJobNameD + " not found")
		return -2
	}
	_, err = send(*urlD+"/apis/batch/v1/namespaces/"+*namespaceD+"/jobs/"+name, *userD, *pwD, []byte(""), "DELETE")
	if err != nil {
		return -1
	}
	return 0
}
