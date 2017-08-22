// +build windows

package containerd

import (
	"context"
	"strconv"

	"github.com/containerd/containerd/containers"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

const newLine = "\r\n"

func generateSpec(ctx context.Context, client *Client, opts ...SpecOpts) (*specs.Spec, error) {
	spec, err := GenerateSpec(ctx, client, opts...)
	if err != nil {
		return nil, err
	}

	spec.Windows.LayerFolders = dockerLayerFolders

	return spec, nil
}

func withExitStatus(es int) SpecOpts {
	return func(_ context.Context, _ *Client, s *specs.Spec) error {
		s.Process.Args = []string{"powershell", "-noprofile", "exit", strconv.Itoa(es)}
		return nil
	}
}

func withProcessArgs(args ...string) SpecOpts {
	return WithProcessArgs(append([]string{"powershell", "-noprofile"}, args...)...)
}

func withCat() SpecOpts {
	return WithProcessArgs("cmd", "/c", "more")
}

func withTrue() SpecOpts {
	return WithProcessArgs("cmd", "/c")
}

func withExecExitStatus(s *specs.Process, es int) {
	s.Args = []string{"powershell", "-noprofile", "exit", strconv.Itoa(es)}
}

func withExecArgs(s *specs.Process, args ...string) {
	s.Args = append([]string{"powershell", "-noprofile"}, args...)
}

func withImageConfig(i Image) SpecOpts {
	// TODO: when windows has a snapshotter remove the withImageConfig helper
	return withNoop
}

func withNewSnapshot(id string, i Image) NewContainerOpts {
	// TODO: when windows has a snapshotter remove the withNewSnapshot helper
	return func(ctx context.Context, client *Client, c *containers.Container) error {
		return nil
	}
}

func withUserNamespace(u, g, s uint32) SpecOpts {
	return withNoop
}

func withRemappedSnapshot(id string, i Image, u, g uint32) NewContainerOpts {
	return func(ctx context.Context, client *Client, c *containers.Container) error {
		return nil
	}
}

func withNoop(_ context.Context, _ *Client, _ *specs.Spec) error {
	return nil
}
