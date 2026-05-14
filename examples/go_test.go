//go:build go || all
// +build go all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/stretchr/testify/require"
)

func TestGoExampleLifecycle(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)

	module := filepath.Join(cwd, "../sdk/go/pulumi-provider-template")
	opts := []opttest.Option{
		opttest.GoModReplacement("github.com/Cogni-AI-OU/pulumi-provider-template/sdk/go/pulumi-provider-template", module),
		opttest.AttachProviderServer("provider-boilerplate", providerFactory),
		opttest.SkipInstall(),
	)

	pt.Preview(t)
	pt.Up(t)
	pt.Destroy(t)
}
