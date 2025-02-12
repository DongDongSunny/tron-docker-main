---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond snapshot list

```
trond snapshot list [flags]
```

Refer to the snapshot source domain you input, the available backup snapshots will be showen below.
Note: different domain may have different snapshots that can be downloaded.


### Options


<dl class="flags">
	<dt><code>-d</code>, 
		<code>--domain &lt;string&gt;</code></dt>
	<dd>Domain for target snapshot source (required)
Please run command &#34;./trond snapshot source&#34; to get the available snapshot source domains</dd>
</dl>


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# List available snapshots of target source domain 34.143.247.77
$ ./trond snapshot list -d 34.143.247.77
{% endraw %}{% endhighlight %}

### See also

* [trond snapshot](./trond_snapshot)
