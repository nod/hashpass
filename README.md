
```
 _               _                         
| |__   __ _ ___| |__  _ __   __ _ ___ ___ 
| '_ \ / _` / __| '_ \| '_ \ / _` / __/ __|
| | | | (_| \__ \ | | | |_) | (_| \__ \__ \
|_| |_|\__,_|___/_| |_| .__/ \__,_|___/___/
                      |_|                  
```

# hashpass

quick wrapper around bcrypt to hash passwords

## why

Why do this?  there exist multiple packages for more complex password hashers,
or they get bundled with other applications.  I wanted something I could drop
in a directory and forget about.

## usage


To see help:

```
hashpass -h                        : prints this help and exits,
```


standard usage to hash a password:

```
hashpass - or <password to hash>
```

You can put the password on the command line (often a bad idea as it may show
up in process lists) or specify `-` and then feed the password via stdin like
so:

```
% echo supers3kret > mypassfile
% ./hashpass - < mypassfile 
$2a$04$AdSMH9eNuFXeWwjUlC2J5.1/veJzu1MNy5FdM6HPLh/VM6/MWcuQi
```

finally, you can verify a hash with the following format

```
build/hashpass -verify <password> <hash> : verifies the pass is the hash,
```


