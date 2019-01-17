# MUZSH :zombie:

## A Multi-User Zombie SHooter

### About

This is an early concept for a [MUD](https://en.wikipedia.org/wiki/MUD) I intend to finish some day.

It is written in Go and, as it was just intended to be a quick concept, is just single-player.

### Future

When I pick this project up again, there are multiple things I want to add:

* An actual parser to recognise player input
* A pluggable network system - WebSockets for web client, telnet for traditional MUD clients
* A module to construct the world from a JSON file so the world/story can be made using a web application
* Persistent state - probably using PostgreSQL
* Actual multi-user logic (there's already *some* concept of sessions built in)

This initial concept version will stay on the [v0](https://github.com/kscarlett/muzsh/tree/v0) branch.

### Getting Started

If you want to play what's implemented now, you can - you just need the Go toolchain installed.

Download the contents of this repo with `git clone git@github.com:kscarlett/muzsh.git` or download the zip file.

In the directory of your copy, simply run `go run main.go` and you're off!

You can even modify the game by adding rooms to the game world, as well as items. A good place to get started with that is in `game/world.go`. In the actually good version of the MUD, the world won't be entirely defined in the constructor, don't worry.
