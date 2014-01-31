Forked from https://github.com/surma/gobox.

I updated the code so that it runs with the current Go tools and compilers (Go 1.2)

The following is a stripped down version of the original README (+ updated installation)


GoBox v0.3.1
============
GoBox is supposed to be something like [BusyBox](http://www.busybox.net). I.e.
a single, preferably small executable which bundles all important shell tools.
A swiss army knife for the command line, if you will.
It is being developed with a focus on [Amazon EC2](http://aws.amazon.com) or as
a small footprint basis for an [OpenVZ](http://www.openvz.org) template.

In order to keep the source code and executable small, I have cut a lot of options
you might be used to from [GNU Coreutils](http://www.gnu.org/software/coreutils/) or
similar. I might even have less options than BusyBox itself. I certainly have
fewer applets right now, and probably ever will. But I consider that a good thing.

The current development status can be seen [here](https://trello.com/board/gobox/4ed265f07e5ffd00002b0aed)

Pitfalls
--------
- The shell is *not* a bash, sh or zsh. It is something original, written by me and
  is fairly limited. It does the job of acting as a shell, it‘s hardly adequate for
  scripting, though.
- Telnetd has no authentication mechanism right now. It’s noting more than a
  network-capable pipe.

Installation
------------

	go get github.com/tiborvass/gobox

Why is there not real shell?
----------------------------
I got this question a lot and I have 2 main reasons:

- I seriously did not want to implement the broken and god-awful syntax of bash
  or any other currently used shell!
- You have Go. Do you need anything *more* lightweight? The philosohpy behind this
  project is, that it is cheap to (re)build and deploy. So you don’t really use
  scripting anymore. If you need to automate some process, write an applet in Go and
  integrate it with GoBox and push it.

Bugs
----
Probably

Thanks
------
- Thanks to [Andreas Krennmair](https://github.com/akrennmair) for `grep`, `gzip` and `gunzip`.
- Thanks to [ukaszg](https://github.com/ukaszg) for `head`.

Credits
-------
(c) 2011 Alexander "Surma" Surma <surma@asdf-systems.de>
