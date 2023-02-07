package tradeurl

import (
	"errors"
	"net/url"
	"regexp"
)

func Parse(tradeURL string) ([]string, error) {
	// https://steamcommunity.com/tradeoffer/new/?partner=118985620&token=cNjbWUDT

	re := regexp.MustCompile(`(?m)https://steamcommunity\.com/tradeoffer/new/\?partner=[0-9]{9}&token=[A-Za-z]{8}`)
	if !re.MatchString(tradeURL) {
		return nil, errors.New("not a valid trade url")
	}

	uri, err := url.Parse(tradeURL)
	if err != nil {
		return nil, err
	}

	params := uri.Query()

	partner := params.Get("partner")
	token := params.Get("token")

	return []string{partner, token}, nil
}
