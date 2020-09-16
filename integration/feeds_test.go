package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestGetAllFeeds(t *testing.T) {
	feeds, err := octopusClient.Feeds.GetAll()

	assert.NoError(t, err)
	assert.NotEmpty(t, feeds)

	if err != nil {
		return
	}

	for _, feed := range feeds {
		assert.NotEmpty(t, feed)
	}
}