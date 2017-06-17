// image_test.go

package profitbricks

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

var imgid string

func TestListImages(t *testing.T) {
	setupTestEnv()
	want := 200
	resp := ListImages()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
		t.Errorf(resp.Response)
	}

	imgid = resp.Items[0].Id
	assert.True(t, len(resp.Items) > 0)
}

func TestGetImage(t *testing.T) {
	want := 200
	resp := GetImage(imgid)

	if resp.StatusCode != want {
		if resp.StatusCode == 403 {
			fmt.Println(bad_status(want, resp.StatusCode))
			fmt.Println("This error might be due to user's permission level ")
		}
	}

	assert.Equal(t, resp.Id, imgid)
	assert.Equal(t, resp.Type, "image")
}
