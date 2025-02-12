---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond node

Commands used for operating node, such as:

	1. check and configure node local environment
	2. deploy single Fullnode/SR for different environment(main network, nile network, private network)
	3. stop single node

	** coming soon **
	4. deploy multiple nodes in local single machine
	5. stop multiple nodes in local single machine
	6. deploy multiple nodes in different local machines with ssh(one node on one machine)
	7. deploy wallet-cli

Please refer to the available commands below.


### Available commands

* [trond node env](./trond_node_env)
* [trond node run-single](./trond_node_run-single)


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Help information for node command
$ ./trond node

# Check and configure node local environment
$ ./trond node env

# Run single java-tron fullnode for main network
$ ./trond node run-single -t full-main

# Run single java-tron fullnode for nile network
$ ./trond node run-single -t full-nile

# Run single java-tron witness node for private network
$ ./trond node run-single -t witness-private
{% endraw %}{% endhighlight %}

### See also

* [trond](./trond)
