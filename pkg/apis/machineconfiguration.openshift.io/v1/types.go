package v1

import (
	igntypes "github.com/coreos/ignition/config/v2_2/types"
	configv1 "github.com/openshift/api/config/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCOConfig describes configuration for MachineConfigOperator.
type MCOConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MCOConfigSpec `json:"spec"`
}

// MCOConfigSpec is the spec for MCOConfig resource.
type MCOConfigSpec struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MCOConfigList is a list of MCOConfig resources
type MCOConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MCOConfig `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfig describes configuration for MachineConfigController.
// This is currently only used to drive the MachineConfig objects generated by the TemplateController.
type ControllerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec ControllerConfigSpec `json:"spec"`
	// +optional
	Status ControllerConfigStatus `json:"status"`
}

// ControllerConfigSpec is the spec for ControllerConfig resource.
type ControllerConfigSpec struct {
	// clusterDNSIP is the cluster DNS IP address
	ClusterDNSIP string `json:"clusterDNSIP"`

	// cloudProviderConfig is the configuration for the given cloud provider
	CloudProviderConfig string `json:"cloudProviderConfig"`

	// TODO: Use PlatformType instead of string

	// The openshift platform, e.g. "libvirt", "openstack", "gcp", "baremetal", "aws", or "none"
	Platform string `json:"platform"`

	// etcdDiscoveryDomain specifies the etcd discovery domain
	EtcdDiscoveryDomain string `json:"etcdDiscoveryDomain"`

	// TODO: Use string for CA data

	// etcdCAData specifies the etcd CA data
	EtcdCAData []byte `json:"etcdCAData"`

	// etcdMetricData specifies the etcd metric CA data
	EtcdMetricCAData []byte `json:"etcdMetricCAData"`

	// rootCAData specifies the root CA data
	RootCAData []byte `json:"rootCAData"`

	// TODO: Investigate using a ConfigMapNameReference for the PullSecret and OSImageURL

	// pullSecret is the default pull secret that needs to be installed
	// on all machines.
	PullSecret *corev1.ObjectReference `json:"pullSecret,omitempty"`

	// images is map of images that are used by the controller to render templates under ./templates/
	Images map[string]string `json:"images"`

	// osImageURL is the location of the container image that contains the OS update payload.
	// Its value is taken from the data.osImageURL field on the machine-config-osimageurl ConfigMap.
	OSImageURL string `json:"osImageURL"`

	// proxy holds the current proxy configuration for the nodes
	Proxy *configv1.ProxyStatus `json:"proxy"`

	// infra holds the infrastructure details
	// TODO this makes platform redundant as everything is contained inside Infra.Status
	Infra *configv1.Infrastructure
}

// ControllerConfigStatus is the status for ControllerConfig
type ControllerConfigStatus struct {
	// observedGeneration represents the generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// conditions represents the latest available observations of current state.
	// +optional
	Conditions []ControllerConfigStatusCondition `json:"conditions"`
}

// ControllerConfigStatusCondition contains condition information for ControllerConfigStatus
type ControllerConfigStatusCondition struct {
	// type specifies the state of the operator's reconciliation functionality.
	Type ControllerConfigStatusConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status object.
	// +nullable
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the reason for the condition's last transition.  Reasons are PascalCase
	Reason string `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	Message string `json:"message,omitempty"`
}

// ControllerConfigStatusConditionType valid conditions of a ControllerConfigStatus
type ControllerConfigStatusConditionType string

const (
	// TemplateControllerRunning means the template controller is currently running.
	TemplateControllerRunning ControllerConfigStatusConditionType = "TemplateControllerRunning"

	// TemplateControllerCompleted means the template controller has completed reconciliation.
	TemplateControllerCompleted ControllerConfigStatusConditionType = "TemplateControllerCompleted"

	// TemplateControllerFailing means the template controller is failing.
	TemplateControllerFailing ControllerConfigStatusConditionType = "TemplateControllerFailing"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfigList is a list of ControllerConfig resources
type ControllerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ControllerConfig `json:"items"`
}

// +genclient
// +genclient:noStatus
// +genclient:nonNamespaced
// +k8s:deepcopy-gen=false

// MachineConfig defines the configuration for a machine
type MachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MachineConfigSpec `json:"spec"`
}

// MachineConfigSpec is the spec for MachineConfig
type MachineConfigSpec struct {
	// OSImageURL specifies the remote location that will be used to
	// fetch the OS.
	OSImageURL string `json:"osImageURL"`
	// Config is a Ignition Config object.
	Config igntypes.Config `json:"config"`

	KernelArguments []string `json:"kernelArguments"`

	Fips bool `json:"fips"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigList is a list of MachineConfig resources
type MachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineConfig `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigPool describes a pool of MachineConfigs.
type MachineConfigPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec MachineConfigPoolSpec `json:"spec"`
	// +optional
	Status MachineConfigPoolStatus `json:"status"`
}

// MachineConfigPoolSpec is the spec for MachineConfigPool resource.
type MachineConfigPoolSpec struct {
	// machineConfigSelector specifies a label selector for MachineConfigs.
	// Refer https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ on how label and selectors work.
	MachineConfigSelector *metav1.LabelSelector `json:"machineConfigSelector,omitempty"`

	// nodeSelector specifies a label selector for Machines
	NodeSelector *metav1.LabelSelector `json:"nodeSelector,omitempty"`

	// paused specifies whether or not changes to this machine config pool should be stopped.
	// This includes generating new desiredMachineConfig and update of machines.
	Paused bool `json:"paused"`

	// maxUnavailable specifies the percentage or constant number of machines that can be updating at any given time.
	// default is 1.
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`

	// The targeted MachineConfig object for the machine config pool.
	Configuration MachineConfigPoolStatusConfiguration `json:"configuration"`
}

// MachineConfigPoolStatus is the status for MachineConfigPool resource.
type MachineConfigPoolStatus struct {
	// observedGeneration represents the generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// configuration represents the current MachineConfig object for the machine config pool.
	Configuration MachineConfigPoolStatusConfiguration `json:"configuration"`

	// machineCount represents the total number of machines in the machine config pool.
	MachineCount int32 `json:"machineCount"`

	// updatedMachineCount represents the total number of machines targeted by the pool that have the CurrentMachineConfig as their config.
	UpdatedMachineCount int32 `json:"updatedMachineCount"`

	// readyMachineCount represents the total number of ready machines targeted by the pool.
	ReadyMachineCount int32 `json:"readyMachineCount"`

	// unavailableMachineCount represents the total number of unavailable (non-ready) machines targeted by the pool.
	// A node is marked unavailable if it is in updating state or NodeReady condition is false.
	UnavailableMachineCount int32 `json:"unavailableMachineCount"`

	// degradedMachineCount represents the total number of machines marked degraded (or unreconcilable).
	// A node is marked degraded if applying a configuration failed..
	DegradedMachineCount int32 `json:"degradedMachineCount"`

	// conditions represents the latest available observations of current state.
	// +optional
	Conditions []MachineConfigPoolCondition `json:"conditions"`
}

// MachineConfigPoolStatusConfiguration stores the current configuration for the pool, and
// optionally also stores the list of MachineConfig objects used to generate the configuration.
type MachineConfigPoolStatusConfiguration struct {
	corev1.ObjectReference `json:",inline"`

	// source is the list of MachineConfig objects that were used to generate the single MachineConfig object specified in `content`.
	// +optional
	Source []corev1.ObjectReference `json:"source,omitempty"`
}

// MachineConfigPoolCondition contains condition information for an MachineConfigPool.
type MachineConfigPoolCondition struct {
	// type of the condition, currently ('Done', 'Updating', 'Failed').
	Type MachineConfigPoolConditionType `json:"type"`

	// status of the condition, one of ('True', 'False', 'Unknown').
	Status corev1.ConditionStatus `json:"status"`

	// lastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	// +nullable
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason"`

	// message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message"`
}

// MachineConfigPoolConditionType valid conditions of a MachineConfigPool
type MachineConfigPoolConditionType string

const (
	// MachineConfigPoolUpdated means MachineConfigPool is updated completely.
	// When the all the machines in the pool are updated to the correct machine config.
	MachineConfigPoolUpdated MachineConfigPoolConditionType = "Updated"

	// MachineConfigPoolUpdating means MachineConfigPool is updating.
	// When at least one of machine is not either not updated or is in the process of updating
	// to the desired machine config.
	MachineConfigPoolUpdating MachineConfigPoolConditionType = "Updating"

	// MachineConfigPoolNodeDegraded means the update for one of the machine is not progressing
	MachineConfigPoolNodeDegraded MachineConfigPoolConditionType = "NodeDegraded"

	// MachineConfigPoolRenderDegraded means the rendered configuration for the pool cannot be generated because of an error
	MachineConfigPoolRenderDegraded MachineConfigPoolConditionType = "RenderDegraded"

	// MachineConfigPoolDegraded is the overall status of the pool based, today, on whether we fail with NodeDegraded or RenderDegraded
	MachineConfigPoolDegraded MachineConfigPoolConditionType = "Degraded"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MachineConfigPoolList is a list of MachineConfigPool resources
type MachineConfigPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MachineConfigPool `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeletConfig describes a customized Kubelet configuration.
type KubeletConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec KubeletConfigSpec `json:"spec"`
	// +optional
	Status KubeletConfigStatus `json:"status"`
}

// KubeletConfigSpec defines the desired state of KubeletConfig
type KubeletConfigSpec struct {
	MachineConfigPoolSelector *metav1.LabelSelector `json:"machineConfigPoolSelector,omitempty"`
	KubeletConfig             *runtime.RawExtension `json:"kubeletConfig,omitempty"`
}

// KubeletConfigStatus defines the observed state of a KubeletConfig
type KubeletConfigStatus struct {
	// observedGeneration represents the generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// conditions represents the latest available observations of current state.
	// +optional
	Conditions []KubeletConfigCondition `json:"conditions"`
}

// KubeletConfigCondition defines the state of the KubeletConfig
type KubeletConfigCondition struct {
	// type specifies the state of the operator's reconciliation functionality.
	Type KubeletConfigStatusConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status object.
	// +nullable
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the reason for the condition's last transition.  Reasons are PascalCase
	Reason string `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	Message string `json:"message,omitempty"`
}

// KubeletConfigStatusConditionType is the state of the operator's reconciliation functionality.
type KubeletConfigStatusConditionType string

const (
	// KubeletConfigSuccess designates a successful application of a KubeletConfig CR.
	KubeletConfigSuccess KubeletConfigStatusConditionType = "Success"

	// KubeletConfigFailure designates a failure applying a KubeletConfig CR.
	KubeletConfigFailure KubeletConfigStatusConditionType = "Failure"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeletConfigList is a list of KubeletConfig resources
type KubeletConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KubeletConfig `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerRuntimeConfig describes a customized Container Runtime configuration.
type ContainerRuntimeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec ContainerRuntimeConfigSpec `json:"spec"`
	// +optional
	Status ContainerRuntimeConfigStatus `json:"status"`
}

// ContainerRuntimeConfigSpec defines the desired state of ContainerRuntimeConfig
type ContainerRuntimeConfigSpec struct {
	MachineConfigPoolSelector *metav1.LabelSelector          `json:"machineConfigPoolSelector,omitempty"`
	ContainerRuntimeConfig    *ContainerRuntimeConfiguration `json:"containerRuntimeConfig,omitempty"`
}

// ContainerRuntimeConfiguration defines the tuneables of the container runtime
type ContainerRuntimeConfiguration struct {
	// pidsLimit specifies the maximum number of processes allowed in a container
	PidsLimit int64 `json:"pidsLimit,omitempty"`

	// logLevel specifies the verbosity of the logs based on the level it is set to.
	// Options are fatal, panic, error, warn, info, and debug.
	LogLevel string `json:"logLevel,omitempty"`

	// logSizeMax specifies the Maximum size allowed for the container log file.
	// Negative numbers indicate that no size limit is imposed.
	// If it is positive, it must be >= 8192 to match/exceed conmon's read buffer.
	LogSizeMax resource.Quantity `json:"logSizeMax"`

	// overlaySize specifies the maximum size of a container image.
	// This flag can be used to set quota on the size of container images. (default: 10GB)
	OverlaySize resource.Quantity `json:"overlaySize"`
}

// ContainerRuntimeConfigStatus defines the observed state of a ContainerRuntimeConfig
type ContainerRuntimeConfigStatus struct {
	// observedGeneration represents the generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// conditions represents the latest available observations of current state.
	// +optional
	Conditions []ContainerRuntimeConfigCondition `json:"conditions"`
}

// ContainerRuntimeConfigCondition defines the state of the ContainerRuntimeConfig
type ContainerRuntimeConfigCondition struct {
	// type specifies the state of the operator's reconciliation functionality.
	Type ContainerRuntimeConfigStatusConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status object.
	// +nullable
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the reason for the condition's last transition.  Reasons are PascalCase
	Reason string `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	Message string `json:"message,omitempty"`
}

// ContainerRuntimeConfigStatusConditionType is the state of the operator's reconciliation functionality.
type ContainerRuntimeConfigStatusConditionType string

const (
	// ContainerRuntimeConfigSuccess designates a successful application of a ContainerRuntimeConfig CR.
	ContainerRuntimeConfigSuccess ContainerRuntimeConfigStatusConditionType = "Success"

	// ContainerRuntimeConfigFailure designates a failure applying a ContainerRuntimeConfig CR.
	ContainerRuntimeConfigFailure ContainerRuntimeConfigStatusConditionType = "Failure"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerRuntimeConfigList is a list of ContainerRuntimeConfig resources
type ContainerRuntimeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ContainerRuntimeConfig `json:"items"`
}
