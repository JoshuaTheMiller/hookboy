package conf

// RecognizedHooks files and folder names that will be automatically recognized
// during processing
var RecognizedHooks = [...]string{
	"applypatch-msg",
	"commit-msg",
	"fsmonitor-watchman",
	"post-update",
	"pre-applypatch",
	"pre-commit",
	"pre-merge-commit",
	"pre-push",
	"pre-rebase",
	"pre-receive",
	"prepare-commit-msg",
	"update"}

// ActualGitHooksDir the default directory for where hooks
// should be applied to
var ActualGitHooksDir = ".git/hooks/"

// GrappleCacheDir the location where temporary files should be stored
var GrappleCacheDir = ".grapple-cache"

// DefaultLocalHooksDir the directory where local hooks can be found
var DefaultLocalHooksDir = "./hooks"
