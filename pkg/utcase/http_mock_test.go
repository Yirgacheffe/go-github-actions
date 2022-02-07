package utcase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var feed = `<?xml version="1.0" encoding="utf-8"?>
<rss>
</rss>
`

func mockServer() *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(200)
			fmt.Fprint(w, feed)
		}))
}

func Test_Download_Mock(t *testing.T) {

	server := mockServer()
	defer server.Close()

	url := server.URL
	statusCode := http.StatusOK

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to  make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v, %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

func Test_Download_404(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 404

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to  make the Get call.", ballotX, err)
			}

			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v, %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
