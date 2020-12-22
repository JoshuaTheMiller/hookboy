<div align="center">

# Grapple

![An image of a hook, representing the end of a grappling hook](./docs/images/hook.svg)

[![Go](https://github.com/JoshuaTheMiller/Grapple/workflows/Go/badge.svg)](https://github.com/JoshuaTheMiller/Grapple/actions?query=workflow%3AGo) [![codecov](https://codecov.io/gh/JoshuaTheMiller/hookboy/branch/main/graph/badge.svg?token=EJ36OJYYF1)](https://codecov.io/gh/JoshuaTheMiller/hookboy)

</div>

## Welcome!

Similar to the ever popular [grappling hook](https://terraria.gamepedia.com/Hooks), Grapple is an essential tool for anyone looking to increase the mobility (i.e. shareability) of [Git Hooks](https://git-scm.com/docs/githooks). Additionally, it makes using them a ton easier!

**Why should you care?** First and foremost, Git Hooks provide a way to automate your own workflows. Properly leveraging such automation provides [other benefits](#Benefits-of-Git-Hooks), such as forcing yourself to get better at commit messages, linting your source code upon commit, etc. Unfortunately, git does not offer a great out-of-box experience for maintaining, installing, and sharing Hooks. Grapple aims to provide this great experience!

## So... what actually is Grapple?

Currently, Grapple is a tiny command line tool that allows users to "install" Git Hooks, either from a local file, or from some file hosted out on the internet.

Under normal circumstances, to "Install" a Hook, one would have to navigate to the `.git/hooks` subdirectory of a repository, and modify/overwrite files there. This quickly gets tiresome with many files and many repositories, as Git does not natively offer an easy way to transfer Hooks around. A quick search around the internet will show you that there are other tools available today that offer a similar benefit to Grapple, and you should [check them out](#Other-Tools)! However, Grapple shines because of its small footprint, the drive to support any and all tech stacks (i.e. not just Javascript, Python, etc), and its stance towards user friendliness (i.e. be user friendly).

## Benefits of Git Hooks

* Reduce manual, repetitive tasks (of course).
   * For example, do you run a linter *one last time* before committing your code to make sure everything is up to par? Why not use a `pre-commit` Hook that runs a linter before accepting a commit?
* Force yourself to get better.
   * Did you tell yourself you would write better commit messages, only to continue to say, "I'll amend them later?" Well, you can use a Hook that inspects commit messages, and rejects them if they are bad! (bad is whatever you define it to be, everyone has different standards)
* Shift left
   * If you have validation that normally runs on the server/remote to check commits for *bad stuff*™ (e.g. bad formatting, spelling, breaking of tests, etc), you can run them locally so that you can find a bad commit before you push!
   * Tip: if you think you want to block the pushing of commits to a remote, you should at least provide a helpful message as to why so the pusher can quickly fix their code.
   * Honestly, this is just a rewording of the top two reasons... Having three bullet points are nice though.

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

- [ ] Have script add colons to **recognized** tags (less typing for me)
- [ ] Support usage of files over HTTP
  * Needs to support caching (.grapple-cache directory?)
  * Needs to support setting of execution rights after downloading
- [x] ~~Error handling in MyNameWillChange.sh~~
- [x] ~~Rewrite CaptainHook.go as Bash?~~

## Other Tools

* https://github.com/rycus86/githooks
* https://github.com/Autohook/Autohook
* https://github.com/sds/overcommit
* https://pre-commit.com/
* https://www.npmjs.com/package/node-hooks
* https://pythonhosted.org/jig/
* https://github.com/typicode/husky