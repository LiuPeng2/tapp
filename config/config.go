package config

import (
	"k8s.io/klog"
)

var (
	EnableUpdateEtcdInOrder = false
)

func SetEnableUpdateEtcdInOrderFlag(enable bool) {
	EnableUpdateEtcdInOrder = enable
	klog.V(2).Infof("EnableUpdateEtcdInOrder is %v", enable)
}
