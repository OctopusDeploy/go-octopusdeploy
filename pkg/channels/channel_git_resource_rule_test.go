package channels

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/gitdependencies"
	"github.com/stretchr/testify/require"
)

func TestChannelValidateDivesIntoGitResourceRules(t *testing.T) {
	newChannelWithGitDependency := func(dep gitdependencies.DeploymentActionGitDependency) *Channel {
		channel := NewChannel("my-channel", "Projects-1")
		channel.GitResourceRules = []ChannelGitResourceRule{
			{GitDependencyActions: []gitdependencies.DeploymentActionGitDependency{dep}},
		}
		return channel
	}

	t.Run("valid nested git dependency passes", func(t *testing.T) {
		channel := newChannelWithGitDependency(gitdependencies.DeploymentActionGitDependency{
			DeploymentActionSlug: "deploy-action",
			GitDependencyName:    "",
		})
		require.NoError(t, channel.Validate())
	})

	t.Run("empty deployment action slug fails", func(t *testing.T) {
		channel := newChannelWithGitDependency(gitdependencies.DeploymentActionGitDependency{
			DeploymentActionSlug: "",
			GitDependencyName:    "my-dependency",
		})
		require.Error(t, channel.Validate())
	})

	t.Run("whitespace-only deployment action slug fails", func(t *testing.T) {
		channel := newChannelWithGitDependency(gitdependencies.DeploymentActionGitDependency{
			DeploymentActionSlug: "   ",
			GitDependencyName:    "my-dependency",
		})
		require.Error(t, channel.Validate())
	})
}
