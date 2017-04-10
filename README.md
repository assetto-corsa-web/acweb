![AC](aclogo.png)

# Assetto Corsa server web interface

This tool provides monitoring and management for your Assetto Corsa server instances. You can create multiple configuration profiles, start/stop server instances and watch the status of them.

**STATUS: BETA**

## Screenshots

![Screenshot 1](screenshots/screen1.png)
![Screenshot 2](screenshots/screen2.png)
![Screenshot 3](screenshots/screen3.png)
![Screenshot 4](screenshots/screen4.png)
![Screenshot 5](screenshots/screen5.png)
![Screenshot 6](screenshots/screen6.png)

## Installation

This instruction supposes you to use Linux. You need a MySQL database and rights to upload and execute applications. I recommend to create a user for your web interface installation.

1. download the latest release of acweb
2. upload it to your server and unzip it
3. cd into the folder and edit the config.json:

```
{
    "host": "0.0.0.0:3000", // port the web server will use
    "logfile": "log",       // if set, log is printed to this directory
    "tls_private_key": "",  // you can use TLS by adding paths to private key and cert file
    "tls_cert": "",
    "dbuser": "root",       // database user
    "dbpwd": "",            // database user password
    "dbhost": "",           // database host, e.g. tcp(127.0.0.1:3306)
    "db": "db"              // database name
}
```

4. create the database schema (db/schema.sql) and create the first user:

```
INSERT INTO `user` (`id`, `login`, `email`, `password`, `admin`, `moderator`) VALUES (NULL, 'username', 'user@email.com', 'SHA256_HASH', '1', '0');
```

5. start it ./acweb
6. you can now visit your web interface

## Updating

1. download the latest release
2. upload it to your server and unzip it
3. cp the config.json from the old version to the new version
4. update your MySQL database (db/mig_FROM_TO.sql)
5. start it

## License

MIT
