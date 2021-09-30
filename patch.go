package main

import (
	"encoding/json"
	"fmt"
)

func patchCronJob() int {
	cronJob := CronJob{
		Spec: CronJobSpec{
			Schedule: *scheduleP,
			JobTemplate: JobTemplate{
				Spec: JobTemplateSpec{
					Template: Template{
						Spec: PodSpec{
							RestartPolicy: "OnFailure",
							ImagePullSecrets: []struct {
								Name string `json:"name"`
							}{
								{
									Name: *imagepullsecretP,
								},
							},
							Containers: []Container{
								{
									Name:  *cronJobNameP,
									Image: *imageP,
								},
							},
						},
					},
				},
			},
		},
	}
	if secretP != nil {
		volumeMounts := make([]VolumeMount, 0, len(*secretP))
		volumes := make([]Volume, 0, len(*secretP))
		for k, v := range *secretP {
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

	if envP != nil {
		envs := make([]EnvVar, 0, len(*envP))
		for k, v := range *envP {
			env := EnvVar{
				Name:  k,
				Value: v,
			}
			envs = append(envs, env)
		}
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env = envs
	}
	if limitP != nil && len(*limitP) > 0 {
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources.Limits = &Memory{*limitP}
	}
	if requestP != nil && len(*requestP) > 0 {
		cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources.Requests = &Memory{*requestP}
	}

	cronJobJSON, err := json.Marshal(cronJob)
	if err != nil {
		fmt.Printf("err : %v", err)
		return -1
	}
	// fmt.Println(string(cronJobJSON))
	_, err = send(*urlP+"/apis/batch/v1beta1/namespaces/"+*namespaceP+"/cronjobs/"+*cronJobNameP, *userP, *pwP, cronJobJSON, "PATCH")
	if err != nil {
		return -1
	}
	return 0
}
