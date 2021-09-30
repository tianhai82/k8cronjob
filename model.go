package main

type CronJob struct {
	APIVersion string      `json:"apiVersion,omitempty"`
	Kind       string      `json:"kind,omitempty"`
	Metadata   *Metadata   `json:"metadata,omitempty"`
	Spec       CronJobSpec `json:"spec"`
}
type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}
type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Memory struct {
	Memory string `json:"memory"`
}
type Resources struct {
	Limits   *Memory `json:"limits,omitempty"`
	Requests *Memory `json:"requests,omitempty"`
}
type Container struct {
	Name         string        `json:"name"`
	Image        string        `json:"image"`
	Args         []string      `json:"args,omitempty"`
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
	Env          []EnvVar      `json:"env,omitempty"`
	Resources    Resources     `json:"resources,omitempty"`
}
type Secret struct {
	SecretName string `json:"secretName"`
}
type Volume struct {
	Name   string `json:"name"`
	Secret Secret `json:"secret,omitempty"`
}
type PodSpec struct {
	Containers       []Container `json:"containers"`
	RestartPolicy    string      `json:"restartPolicy"`
	Volumes          []Volume    `json:"volumes,omitempty"`
	ImagePullSecrets []struct {
		Name string `json:"name"`
	} `json:"imagePullSecrets"`
}
type Template struct {
	Spec PodSpec `json:"spec"`
}
type JobTemplateSpec struct {
	Template Template `json:"template"`
}
type JobTemplate struct {
	Spec JobTemplateSpec `json:"spec"`
}
type CronJobSpec struct {
	Schedule    string      `json:"schedule"`
	JobTemplate JobTemplate `json:"jobTemplate"`
}
