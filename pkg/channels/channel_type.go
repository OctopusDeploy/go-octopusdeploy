package channels

// ChannelType is the type of channel.
type ChannelType string

const (
	ChannelTypeEphemeral ChannelType = "EphemeralEnvironment"
	ChannelTypeLifecycle ChannelType = "Lifecycle"
)
