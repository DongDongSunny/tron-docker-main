---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond snapshot download

```
trond snapshot download [flags]
```

Refer to the snapshot source domain and backup name you input, the available backup snapshot will be downloaded to the local directory.

Note:
- because some snapshot sources have multiple snapshot types, you need to specify the type(full, lite) of snapshot you want to download.
- the snapshot is large, it may need a long time to finish the download, depends on your network performance.


### Available commands

* [trond snapshot download default-main](./trond_snapshot_download_default-main)
* [trond snapshot download default-nile](./trond_snapshot_download_default-nile)


### Options


<dl class="flags">
	<dt><code>-b</code>, 
		<code>--backup &lt;string&gt;</code></dt>
	<dd>Backup name(required).
Please run command &#34;./trond snapshot list&#34; to get the available backup name under target source domains.</dd>

	<dt><code>-d</code>, 
		<code>--domain &lt;string&gt;</code></dt>
	<dd>Domain for target snapshot source(required).
Please run command &#34;./trond snapshot source&#34; to get the available snapshot source domains.</dd>

	<dt><code>-t</code>, 
		<code>--type &lt;string&gt;</code></dt>
	<dd>Node type of the snapshot(required, available: full, lite).</dd>
</dl>


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Download target backup snapshot (backup20250205 in 34.143.247.77) to current directory
$ ./trond snapshot download -d 34.143.247.77 -b backup20250205 -t lite
{% endraw %}{% endhighlight %}

### See also

* [trond snapshot](./trond_snapshot)
