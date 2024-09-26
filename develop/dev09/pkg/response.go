package pkg

import "net/http"

func getResponse(url string) (*http.Response, error) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
