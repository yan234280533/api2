package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName specifies the group name used to register the objects.
const GroupName = "ensurance.crane.io"

// GroupVersion specifies the group and the version used to register the objects.
var GroupVersion = v1.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// SchemeGroupVersion is group version used to register these objects
// Deprecated: use GroupVersion instead.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	// Depreciated: use Install instead
	AddToScheme = localSchemeBuilder.AddToScheme
	Install     = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes, addDefaultingFuncs)
}

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&PodQOSEnsurancePolicy{},
		&PodQOSEnsurancePolicyList{},
		&NodeQOSEnsurancePolicy{},
		&NodeQOSEnsurancePolicyList{},
		&AvoidanceAction{},
		&AvoidanceActionList{},
	)

	// AddToGroupVersion allows the serialization of client types like ListOptions.
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}
