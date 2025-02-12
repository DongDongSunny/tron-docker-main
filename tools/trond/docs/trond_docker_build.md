---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond docker build

```
trond docker build [flags]
```

Build java-tron docker image locally.
The master branch of java-tron repository will be built by default, using jdk1.8.0_202.


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
$ ./trond docker build

# Build java-tron docker image with specified org, artifact and version
$ ./trond docker build -o tronprotocol -a java-tron -v latest
{% endraw %}{% endhighlight %}

### See also

* [trond docker](./trond_docker)
