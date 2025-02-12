---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond docker test

```
trond docker test [flags]
```

Test java-tron docker image locally.
Default, will test the "tronprotocol/java-tron:latest" image. You can specify the flags to test the image you want.
The test includes the following tasks:
	1. Perform port checks
	2. Verify whether block synchronization is functioning normally


### Options


<dl class="flags">
	<dt><code>-a</code>, 
		<code>--artifact &lt;string&gt; (default &#34;java-tron&#34;)</code></dt>
	<dd>ArtifactName for the docker image</dd>

	<dt><code>-o</code>, 
		<code>--org &lt;string&gt; (default &#34;tronprotocol&#34;)</code></dt>
	<dd>OrgName for the docker image</dd>

	<dt><code>-v</code>, 
		<code>--version &lt;string&gt; (default &#34;latest&#34;)</code></dt>
	<dd>Release version for the docker image</dd>
</dl>


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Build java-tron docker image, defualt: tronprotocol/java-tron:latest
$ ./trond docker test

# Build java-tron docker image with specified org, artifact and version
$ ./trond docker test -o tronprotocol -a java-tron -v latest
{% endraw %}{% endhighlight %}

### See also

* [trond docker](./trond_docker)
