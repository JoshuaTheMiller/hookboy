# Hooks

## First Line

`#!/usr/bin/env python`

## Available Git Hooks

* applypatch-msg
* commit-msg
* fsmonitor-watchman
* post-update
* pre-applypatch
* pre-commit
* pre-merge-commit
* pre-push
* pre-rebase
* pre-receive
* prepare-commit-msg
* update

## Wants+Improvements

- [ ] Error handling in MyNameWillChange.sh
- [ ] Rewrite CaptainHook.go as Bash?
- [ ] Have script add colons to **recognized** tags (less typing for me)

## Quick Notes

After poking around the internet a bit more, it looks like someone already has made something similar to what I wanted (almost exactly, and for the exact same reasons): https://github.com/git-hooks/git-hooks/wiki/Thoughts

They even got to the git-hooks GitHub Org! Ah well.

I believe there are enough difference in what I want to do to warrant this being a fun side project. Namely, I do not wish to use symbolic links, as I have seen folks run into issues with them.

I will continue this project with Go, as getting up and running with Go on any system is simple, and the footprint is tiny.

## Related Projects

* https://github.com/rycus86/githooks
* https://github.com/Autohook/Autohook
* https://github.com/sds/overcommit
* https://pre-commit.com/
* https://www.npmjs.com/package/node-hooks
* https://pythonhosted.org/jig/
* https://github.com/typicode/husky