Manage alerts from the comfort of your tty.

### Install

Download [official binaries](https://github.com/ttyops/cli/releases/latest) for Linux,
macOS, NetBSD, FreeBSD, OpenBSD, Windows and Solaris.

### Configure

Create a `~/.config/ttyops/config.toml` with the following:

```
token = "Bearer <token>"
endpoint = "https://ttyops.com/api/v1"
```

using a `token` from [your team dashboard](http://ttyops.com/team).

### Examples

Available commands are shown with `ttyops help`.  The most common
commands are listing alerts:

```
% ttyops alert list
ID		TITLE		STATUS		SERVICE		ASSIGNED
c2d16fed	InstanceDown	firing		default		bob@ttyops.com
e69fee0b	InstanceDown	acknowledged	default		bob@ttyops.com
```

seeing who's on-call:

```
% ttyops schedule list
NAME		ON-CALL
Core infra	bob@ttyops.com
```

and creating alerts:

```
% ttyops alert create "Core infra" "On fire"
alert created
```
