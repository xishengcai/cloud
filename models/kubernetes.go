package models

// Kubernetes master nodes
type Kubernetes struct {
	ID                   int    `form:"-"`
	Uid                  string `form:"-"`
	ClusterName          string `form:"clusterName" binding:"required"`
	PrimaryMaster        Host   `form:"primaryMaster" binding:"required"`
	BackendMasters       []Host `form:"backendMasters"`
	NetWorkPlug          string `form:"networkPlug,default=cilium"`
	Registry             string `form:"registry,default=k8s.gcr.io"`
	Version              string `form:"version,default=1.17.2"`
	ControlPlaneEndpoint string `form:"controlPlaneEndpoint" binding:"required"`
	PodCidr              string `form:"podCidr,default=10.244.0.0/16"`
	ServiceCidr          string `form:"serviceCidr,default=10.96.0.0/16"`
	JoinMasterCommand    string `json:"joinMasterCommand"`
}

// KubernetesSlave k8s slave node
type KubernetesSlave struct {
	Version          string `form:"version"`
	Nodes            []Host `form:"nodes"`
	Master           Host   `form:"master"`
	JoinSlaveCommand string `form:"joinSlaveCommand"`
}

// Version kubernetes version
type Version struct {
	Version string `form:"version,default=1.17.2"`
}
