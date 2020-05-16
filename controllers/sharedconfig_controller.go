/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configv1 "github.com/prasadpnm/PrasadController/api/v1"
)

// SharedConfigReconciler reconciles a SharedConfig object
type SharedConfigReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=config.devops.prasad.com,resources=sharedconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.devops.prasad.com,resources=sharedconfigs/status,verbs=get;update;patch

func (r *SharedConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("sharedconfig", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SharedConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&configv1.SharedConfig{}).
		Complete(r)
}
