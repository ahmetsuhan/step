# step
Easy SSH management tool with encryption, you can think of alias for ssh connections with encryption

**Arguments**

````shell
    step {ALIAS} {COMMAND}
````

**Example**

````shell
    step customer_1 "cd /path/to/file/ ; bash update.sh"
````

**Save New Alias**

````shell
    step -s -i /full/path/key/ IP@server
````
````shell
    -i flag is optional
````

**Need a hint ?**
````shell
    step -h
````

**Todo**
* Use Tab key to search alias names