package models

// Kubernetes master nodes
type Kubernetes struct {
	ID             int    `json:"-"`
	Uid            string `json:"-"`
	Name           string `json:"name" validate:"required" default:"test"`
	PrimaryMaster  Host   `json:"primaryMaster" validate:"required"`
	BackendMasters []Host `json:"backendMasters"`
	NetWorkPlug    string `json:"networkPlug" default:"cilium"`
	// registry.aliyuncs.com/google_containers， k8s.gcr.io
	Registry             string `json:"registry" default:"registry.aliyuncs.com/google_containers"`
	Version              string `json:"version" default:"1.21.0"`
	ControlPlaneEndpoint string `json:"controlPlaneEndpoint" validate:"required"`
	PodCidr              string `json:"podCidr" default:"10.244.0.0/16"`
	ServiceCidr          string `json:"serviceCidr" default:"10.96.0.0/16"`
	JoinMasterCommand    string `json:"-"`
}

// KubernetesSlave k8s slave node
type KubernetesSlave struct {
	Version          string `form:"version"`
	Nodes            []Host `form:"nodes"`
	Master           Host   `form:"master"`
	JoinSlaveCommand string `form:"joinSlaveCommand"`
}

// Version cluster version
type Version struct {
	Version string `form:"version" default:"1.21.0"`
}
