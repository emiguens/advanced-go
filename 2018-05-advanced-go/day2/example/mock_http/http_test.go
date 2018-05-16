package mock_http_test

import (
	"io/ioutil"
	"testing"

	"github.com/mercadolibre/advanced-go/2018-05-advanced-go/day2/example/mock_http"
	"github.com/stretchr/testify/require"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestFetchISO(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	j := `[{"id": 1, "name": "My Great Article"}]`
	mockRes := httpmock.NewStringResponder(200, j)
	httpmock.RegisterResponder("GET", mock_http.ISO, mockRes)

	res, err := mock_http.FetchISO()
	require.NoError(t, err)

	require.Equal(t, 200, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	defer res.Body.Close()

	require.JSONEq(t, j, string(body))
}
