---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond snapshot

Commands used for downloading node's snapshot, such as:
	1. show available snapshot source
	2. list available snapshots in target source
	3. download target snapshot


### Available commands

* [trond snapshot download](./trond_snapshot_download)
* [trond snapshot list](./trond_snapshot_list)
* [trond snapshot source](./trond_snapshot_source)


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Help information for snapshot command
$ ./trond snapshot

# Show available snapshot source
$ ./trond snapshot source

# List available snapshots of target source domain 34.143.247.77
$ ./trond snapshot list -d 34.143.247.77

# Download target backup snapshot (backup20250205 in 34.143.247.77) to current directory
$ ./trond snapshot download -d 34.143.247.77 -b backup20250205 -t lite

# Download latest mainnet lite fullnode snapshot from default source(34.143.247.77) to current directory
$ ./trond snapshot download default-main

# Download latest nile testnet lite fullnode snapshot from default source(database.nileex.io) to current directory
$ ./trond snapshot download default-nile
{% endraw %}{% endhighlight %}

### See also

* [trond](./trond)
