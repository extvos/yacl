# yacl: Yet Another Configuration Language

Idea from Nginx Configuration File format.

## Format


```
# comments ...
<directive> [param1, ...];
...

<section> [param1,...] {
	<directive> [param1, ...];
	...
	<sub_section> [param1, ...] {
		<directive> [param1, ...];
		...
	}
}
```


## Example

```
### This is an example
set max_mem_size 100m;
set max_file_size 2g;

proxy example.org {
    listen 80;
    path /a/b/c {
        return 404;
    }
}

include another/config.yac;
```