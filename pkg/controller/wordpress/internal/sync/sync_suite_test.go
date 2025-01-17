/*
Copyright 2019 Pressinfra SRL.

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

package sync

import (
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func TestPodTemplate(t *testing.T) {
	logf.SetLogger(logf.ZapLoggerTo(GinkgoWriter, true))
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Wordpress Sync Test Suite", []Reporter{envtest.NewlineReporter{}})
}
