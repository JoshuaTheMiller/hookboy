package hookboy

var recognizedHooks = [...]string{
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

var actualGitHooksDir = ".git/hooks/"
var grappleCacheDir = ".grapple-cache"
var defaultLocalHooksDir = "./hooks"
var hooksInstalledMessage = "Hooks installed!"

// TODO: modify so that it checks a variety
// of locations and returns the first match it finds.
// Global config should used ONLY IF no other items
// are matched (when global config is implemented)
func retrieveConfigPath() string {
	return ".gitgrapple.yml"
}
