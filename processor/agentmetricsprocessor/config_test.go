// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agentmetricsprocessor

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/otelcol/otelcoltest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := otelcoltest.NopFactories()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Processors[componentType] = factory

	cfg, err := otelcoltest.LoadConfigAndValidate(path.Join(".", "testdata", "config.yaml"), factories)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	p1 := cfg.Processors[component.NewID(componentType)]

	want := factory.CreateDefaultConfig()
	want.(*Config).BlankLabelMetrics = []string{"system.cpu.time"}

	assert.Equal(t, want, p1)
}
