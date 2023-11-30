package cluster

type Cluster struct {
	ID          string `bson:"_id" json:"-"`
	Name        string `bson:"name" json:"name" validate:"required" default:"test"`
	Master      []Host `bson:"master" json:"master" validate:"required"`
	NetWorkPlug string `bson:"networkPlug" json:"networkPlug" default:"cilium"`
	// registry.aliyuncs.com/google_containers， k8s.gcr.io
	Registry             string `bson:"registry" json:"registry" default:"registry.aliyuncs.com/google_containers"`
	Version              string `bson:"version" json:"version" default:"1.22.15"`
	ControlPlaneEndpoint string `bson:"controlPlaneEndpoint" json:"controlPlaneEndpoint" validate:"required"`
	PodCidr              string `bson:"podCidr" json:"podCidr" default:"10.244.0.0/16"`
	ServiceCidr          string `bson:"serviceCidr" json:"serviceCidr" default:"10.96.0.0/16"`
	WorkNodes            []Host `bson:"workNodes" json:"workNodes"`
}

type Host struct {
	IP       string `json:"ip" bson:"ip"`
	User     string `json:"user" default:"root" bson:"user"`
	Port     int    `json:"port" bson:"port" default:"22"`
	Password string `json:"password" bson:"password"`
}
