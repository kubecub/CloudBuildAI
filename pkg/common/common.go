package common

import "github.com/mitchellh/go-homedir"

// Default dir and file path
const (
	EtcDir                   = "etc"
	DefaultTmpDir            = "/var/lib/kubecub/tmp"
	DefaultLogDir            = "/var/lib/kubecub/log"
	DefaultkubecubDataDir    = "/var/lib/kubecub/data"
	KubeAdminConf            = "/etc/kubernetes/admin.conf"
	DefaultKubectlPath       = "/usr/bin/kubectl"
	ClusterfileName          = "ClusterfileName"
	RenderChartsDir          = "charts"
	RenderManifestsDir       = "manifests"
	KubeLvsCareStaticPodName = "kube-lvscare"
	RegLvsCareStaticPodName  = "reg-lvscare"
	StaticPodDir             = "/etc/kubernetes/manifests"
	LvsCareRepoAndTag        = "kubecubio/lvscare:v1.1.3-beta.8"
)

func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		return "/root"
	}
	return home
}
