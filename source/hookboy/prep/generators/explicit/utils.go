package explicit

import "net/url"

func pathIsNonLocalPath(path string) bool {
	u, err := url.Parse(path)

	// TODO: this should most likely be expanded. There are edge cases this
	// will not catch. Decent starting point though.
	return err == nil && u.Scheme != "" && u.Host != ""
}
