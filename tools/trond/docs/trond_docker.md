---
layout: manual
permalink: /:path/:basename
---

{% raw %}## trond docker

Commands used for operating docker image, such as:

	1. build java-tron docker image locally
	2. test the built image
	3. pull the latest image from officical docker hub

Please refer to the available commands below.


### Available commands

* [trond docker build](./trond_docker_build)
* [trond docker test](./trond_docker_test)


{% endraw %}
### Examples

{% highlight bash %}{% raw %}
# Help information for docker command
$ ./trond docker

# Build java-tron docker image, defualt: tronprotocol/java-tron:latest
$ ./trond docker build

# Build java-tron docker image with specified org, artifact and version
$ ./trond docker build -o tronprotocol -a java-tron -v latest

# Test java-tron docker image, defualt: tronprotocol/java-tron:latest
$ ./trond docker test

# Test java-tron docker image with specified org, artifact and version
$ ./trond docker test -o tronprotocol -a java-tron -v latest
{% endraw %}{% endhighlight %}

### See also

* [trond](./trond)
