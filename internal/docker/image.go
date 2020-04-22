package docker

type image struct {
	name          string
	gaugeVersion  string
	chromeVersion string
	goVersion     string
	nodeVersion   string
	taikoVersion  string
}

func newImage(name string) *image {
	return &image{name,
		gaugeVersion(name),
		chromeVersion(name),
		goVersion(name),
		nodeVersion(name),
		taikoVersion(name)}
}
