---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond node run-single

```
trond node run-single [flags]
```

You need to make sure the local environment is ready before running the node.
Use this command "./trond node env" to check the environment before starting the node.

The following files are required:
	- Database directory: ./output-directory
	- Configuration file(by default, these exist in the current repository directory)
		main network: ../../conf/main_net_config.conf
		nile network: ../../conf/nile_net_config.conf
		private network: ../../conf/private_net_config_*.conf
	- Docker compose file(by default, these exist in the current repository directory)
		main network: ../../single_node/docker-compose.fullnode.main.yaml
		nile network: ../../single_node/docker-compose.fullnode.nile.yaml
		private network: ../../single_node/docker-compose.witness.private.yaml
	- Log directory: ./logs


### Options


<dl class="flags">
	<dt><code>-s</code>, 
		<code>--service &lt;string&gt;</code></dt>
	<dd>Service name you want to start in docker-compose.*.yaml (optional, default: tron-node)</dd>

	<dt><code>-t</code>, 
		<code>--type &lt;string&gt;</code></dt>
	<dd>Node type you want to deploy (required, available: full-main, full-nile, witness-private)</dd>
</dl>


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Run single java-tron fullnode for main network
$ ./trond node run-single -t full-main

# Run single java-tron fullnode for nile network
$ ./trond node run-single -t full-nile

# Run single java-tron witness node for private network
$ ./trond node run-single -t witness-private
{% endraw %}{% endhighlight %}

### See also

* [trond node](./trond_node)
