# Amity
Amity client/server

## Installation
```bash
$ git clone git://github.com/mclellac/amity.git
$ cd amity && make deps
$ make 
```

## Configuration
Example amityd.conf configuration file can be found in amity/config/

``` ini
[server]
domainname      = localhost:3000

[database]
username        = dbuser
password        = "dbpass"
hostname        = localhost
databasename    = amity

```


## nginx.conf:
``` nginx
    server {
        listen          80;
        server_name     example.com;
        root            /your/webroot/path;
        charset         utf-8;
        access_log      /var/log/nginx/example.com.access.log  main;
    
        location / {
                proxy_pass      http://127.0.0.1:3000;
        }
 ``` 

 ## Start the amityd service:
 ```bash
$ ./amityd start
 ```

 ## Add a post with the client
 ```bash
 $ ./amity add "I'm the title of the post" "... and I'm the article."
 ```

 ## List all of the posts
 ```bash
 $ ./amity ls
 ```

 ## Delete a post
 ```bash
 $ ./amity rm $ID
 ```
