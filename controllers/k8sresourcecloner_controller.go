/*
Copyright 2020 The Kubernetes authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	cachev1alpha1 "github.com/girishg4t/k8sResourceCloner/api/v1alpha1"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// K8sResourceClonerReconciler reconciles a K8sResourceCloner object
type K8sResourceClonerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cache.example.org,resources=k8sresourcecloners,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.example.org,resources=k8sresourcecloners/status,verbs=get;update;patch

func (r *K8sResourceClonerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("k8sresourcecloner", req.NamespacedName)

	// Fetch the K8sResourceCloner instance
	k8sResoureCloner := &cachev1alpha1.K8sResourceCloner{}
	err := r.Get(ctx, req.NamespacedName, k8sResoureCloner)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("K8sResourceCloner resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get K8sResourceCloner")
		return ctrl.Result{}, err
	}

	configMapInstance := &corev1.ConfigMap{}
	for j := 0; j < len(k8sResoureCloner.Spec.ResourceNames); j++ {
		err = r.Get(ctx, types.NamespacedName{Name: k8sResoureCloner.Spec.ResourceNames[j],
			Namespace: k8sResoureCloner.Spec.FromNamespace}, configMapInstance)
		if err == nil {
			for i := 0; i < len(k8sResoureCloner.Spec.ToNamespaces); i++ {
				if err == nil {
					newConfigMap := newConfigMap(configMapInstance, k8sResoureCloner.Spec.ToNamespaces[i])
					err = r.Create(ctx, newConfigMap)
					if err != nil {
						return ctrl.Result{}, err
					}
				}
			}
			break
		}
	}

	secretInstance := &corev1.Secret{}
	for j := 0; j < len(k8sResoureCloner.Spec.ResourceNames); j++ {
		err = r.Get(ctx, types.NamespacedName{Name: k8sResoureCloner.Spec.ResourceNames[j],
			Namespace: k8sResoureCloner.Spec.FromNamespace}, secretInstance)
		if err == nil {
			for i := 0; i < len(k8sResoureCloner.Spec.ToNamespaces); i++ {
				if err == nil {
					newSecret := newSecret(secretInstance, k8sResoureCloner.Spec.ToNamespaces[i])
					err = r.Create(ctx, newSecret)
					if err != nil {
						return ctrl.Result{}, err
					}
				}
			}
			break
		}
	}
	return ctrl.Result{}, nil
}

func newConfigMap(cm *corev1.ConfigMap, ns string) *corev1.ConfigMap {
	labels := map[string]string{
		"app": cm.Name,
	}
	data := make(map[string]string)
	for k, v := range cm.Data {
		data[k] = v
	}

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cm.Name,
			Namespace: ns,
			Labels:    labels,
		},
		Data: data,
	}
}

func newSecret(sc *corev1.Secret, ns string) *corev1.Secret {
	labels := map[string]string{
		"app": sc.Name,
	}
	data := make(map[string][]byte)
	for k, v := range sc.Data {
		data[k] = v
	}

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sc.Name,
			Namespace: ns,
			Labels:    labels,
		},
		Data: data,
	}
}

func (r *K8sResourceClonerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.K8sResourceCloner{}).
		Owns(&corev1.Secret{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
