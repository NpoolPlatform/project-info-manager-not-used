package api

import (
	"os"
	"strconv"
	"testing"

	_ "github.com/NpoolPlatform/go-service-framework/pkg/version"
	"github.com/go-resty/resty/v2"
	"github.com/test-go/testify/assert"
)

func TestVersion(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	_, err := cli.R().
		Post("http://localhost:50310/version")
	assert.Nil(t, err)
}
