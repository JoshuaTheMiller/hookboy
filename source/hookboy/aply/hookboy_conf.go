package aply

var filename = ".hookboy-conf"
var readMeText = "These hooks have been wrangled by Hookboy! Visit https://github.com/JoshuaTheMiller/hookboy for more information."

type hookboyConf struct {
	README string
	// ShowHelp, when true, a help message about hookboy will be displayed upon every hook execution
	// TODO: https://github.com/JoshuaTheMiller/hookboy/issues/21
	ShowHelp *bool
}

func (hc *hookboyConf) Default() hookboyConf {
	// This value will always be changed back for consistency
	hc.README = readMeText

	if hc.ShowHelp == nil {
		var showHelp = true
		hc.ShowHelp = &showHelp
	}

	return *hc
}
