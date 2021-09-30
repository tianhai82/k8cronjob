package main

import (
	"encoding/json"
	"fmt"
)

func createCronJob() int {
	cronJob := CronJob{
		APIVersion: "batch/v1beta1",
		Kind:       "CronJob",
		Metadata: &Metadata{
			Name:      *cronJobName,
			Namespace: *namespace,
		},
		Spec: CronJobSpec{
			Schedule: *schedule,
			JobTemplate: JobTemplate{
				Spec: JobTemplateSpec{
					Template: Template{
						Spec: PodSpec{
							RestartPolicy: "OnFailure",
							ImagePullSecrets: []struct {
								Name string `json:"name"`
							}{
								{
									Name: *imagepullsecret,
								},
							},
							Containers: []Container{
								{
									Name:  *cronJobName,
									Image: *image,
								},
							},
						},
					},
				},
			},
		},
	}
	if secret != nil {
		volumeMounts := make([]VolumeMount, 0, len(*secret))
		volumes := make([]Volume, 0, len(*secret))
		for k, v := range *secret {
			volumeMount := VolumeMount{
				Name:      k,
				MountPath: v,
				ReadOnly:  true,
			}
			volume := Volume{
				Name: k,
				Secret: Secret{
					SecretName: k,
				},
			}
			volumeMounts = append(volumeMounts, volumeMount)
			volumes = append(volumes, volume)
		}
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Volumes = volumes
	}

	if env != nil {
		envs := make([]EnvVar, 0, len(*env))
		for k, v := range *env {
			env := EnvVar{
				Name:  k,
				Value: v,
			}
			envs = append(envs, env)
		}
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env = envs
	}
	if limit != nil && len(*limit) > 0 {
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources.Limits = &Memory{*limit}
	}
	if request != nil && len(*request) > 0 {
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources.Requests = &Memory{*request}
	}

	cronJobJSON, err := json.Marshal(cronJob)
	if err != nil {
		fmt.Printf("err : %v", err)
		return -1
	}
	// fmt.Println(string(cronJobJSON))
	_, err = send(*url+"/apis/batch/v1beta1/namespaces/"+*namespace+"/cronjobs", *user, *pw, cronJobJSON, "POST")
	if err != nil {
		return -1
	}
	return 0
}
