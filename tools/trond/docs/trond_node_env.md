---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond node env

```
trond node env
```

Default environment configuration for node operation:

Current directory: tron-docker/tools/trond

The following files are required:
	- Database directory: ./output-directory (if not exists, will create it)
	- Configuration file(by default, these exist in the current repository directory)
		main network: ../../conf/main_net_config.conf
		nile network: ../../conf/nile_net_config.conf
		private network: ../../conf/private_net_config_*.conf
	- Docker compose file(by default, these exist in the current repository directory)
		single node
			main network: ../../single_node/docker-compose.fullnode.main.yaml
			nile network: ../../single_node/docker-compose.fullnode.nile.yaml
			private network: ../../single_node/docker-compose.witness.private.yaml
		multiple nodes
			private network: ../../private_net/docker-compose.private.yaml
	- Log directory: ./logs (if not exists, will create it)


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Check and configure node local environment
$ ./trond node env
{% endraw %}{% endhighlight %}

### See also

* [trond node](./trond_node)
