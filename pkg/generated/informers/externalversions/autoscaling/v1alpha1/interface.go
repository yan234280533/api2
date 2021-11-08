// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gocrane-io/api/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AdvancedHorizontalPodAutoscalers returns a AdvancedHorizontalPodAutoscalerInformer.
	AdvancedHorizontalPodAutoscalers() AdvancedHorizontalPodAutoscalerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AdvancedHorizontalPodAutoscalers returns a AdvancedHorizontalPodAutoscalerInformer.
func (v *version) AdvancedHorizontalPodAutoscalers() AdvancedHorizontalPodAutoscalerInformer {
	return &advancedHorizontalPodAutoscalerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}