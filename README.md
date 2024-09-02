An implementation of the `tree` Unix command.

Example:

```
$ go run . /somefolder
```

```
wtf
├── .git
│   ├── HEAD
│   ├── config
│   ├── description
│   ├── hooks
│   │   ├── applypatch-msg.sample
│   │   ├── commit-msg.sample
│   │   ├── fsmonitor-watchman.sample
│   │   ├── post-update.sample
│   │   ├── pre-applypatch.sample
│   │   ├── pre-commit.sample
│   │   ├── pre-merge-commit.sample
│   │   ├── pre-push.sample
│   │   ├── pre-rebase.sample
│   │   ├── pre-receive.sample
│   │   ├── prepare-commit-msg.sample
│   │   ├── push-to-checkout.sample
│   │   └── update.sample
│   ├── index
│   ├── info
│   │   └── exclude
│   ├── logs
│   │   ├── HEAD
│   │   └── refs
│   │   │   ├── heads
│   │   │   │   └── main
│   │   │   └── remotes
│   │   │   │   └── origin
│   │   │   │   │   └── HEAD
│   ├── objects
│   │   ├── info
│   │   └── pack
│   │   │   ├── pack-b46d10d05578a47b02d33810fdf6c39490d7bffe.idx
│   │   │   └── pack-b46d10d05578a47b02d33810fdf6c39490d7bffe.pack
│   ├── packed-refs
│   └── refs
│   │   ├── heads
│   │   │   └── main
│   │   ├── remotes
│   │   │   └── origin
│   │   │   │   └── HEAD
│   │   └── tags
├── .gitattributes
├── .github
│   └── workflows
│   │   ├── deploy.yml
│   │   └── test.yml
├── .gitignore
├── LICENSE
├── Makefile
├── README.md
├── auth.go
├── cmd
│   ├── wtf
│   │   ├── dial.go
│   │   ├── dial_create.go
│   │   ├── dial_delete.go
│   │   ├── dial_list.go
│   │   ├── dial_members.go
│   │   ├── dial_set.go
│   │   └── main.go
│   ├── wtf-storybook
│   │   └── main.go
│   └── wtfd
│   │   ├── auth_test.go
│   │   ├── dial_test.go
│   │   ├── main.go
│   │   └── main_test.go
├── context.go
├── csv
│   └── dial.go
├── dial.go
├── dial_membership.go
├── error.go
├── event.go
├── go.mod
├── go.sum
├── http
│   ├── assets
│   │   ├── assets.go
│   │   ├── assets_development.go
│   │   ├── assets_production.go
│   │   ├── css
│   │   │   ├── fontawesome.css
│   │   │   └── theme.css
│   │   ├── fonts
│   │   │   ├── fa-brands-400.eot
│   │   │   ├── fa-brands-400.svg
│   │   │   ├── fa-brands-400.ttf
│   │   │   ├── fa-brands-400.woff
│   │   │   ├── fa-brands-400.woff2
│   │   │   ├── fa-regular-400.eot
│   │   │   ├── fa-regular-400.svg
│   │   │   ├── fa-regular-400.ttf
│   │   │   ├── fa-regular-400.woff
│   │   │   ├── fa-regular-400.woff2
│   │   │   ├── fa-solid-900.eot
│   │   │   ├── fa-solid-900.svg
│   │   │   ├── fa-solid-900.ttf
│   │   │   ├── fa-solid-900.woff
│   │   │   └── fa-solid-900.woff2
│   │   ├── images
│   │   │   ├── background.jpg
│   │   │   └── screenshot.png
│   │   ├── index.html
│   │   └── scripts
│   │   │   ├── main.js
│   │   │   └── reconnecting-websocket.js
│   ├── auth.go
│   ├── auth_test.go
│   ├── dial.go
│   ├── dial_membership.go
│   ├── dial_test.go
│   ├── event.go
│   ├── html
│   │   ├── app.ego
│   │   ├── dial.edit.ego
│   │   ├── dial.index.ego
│   │   ├── dial.view.ego
│   │   ├── dial_membership.create.ego
│   │   ├── error.ego
│   │   ├── html.go
│   │   ├── index.ego
│   │   ├── login.ego
│   │   └── settings.ego
│   ├── http.go
│   ├── server.go
│   └── server_test.go
├── inmem
│   ├── event.go
│   └── event_test.go
├── mock
│   ├── auth.go
│   ├── dial.go
│   ├── dial_membership.go
│   ├── event.go
│   └── user.go
├── scripts
│   └── provision.sh
├── sqlite
│   ├── auth.go
│   ├── auth_test.go
│   ├── dial.go
│   ├── dial_membership.go
│   ├── dial_membership_test.go
│   ├── dial_test.go
│   ├── migration
│   │   └── 00000000.sql
│   ├── sqlite.go
│   ├── sqlite_test.go
│   ├── user.go
│   └── user_test.go
├── tools.go
├── user.go
└── wtf.go
```
