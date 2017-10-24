# Changelog

## 1.3.0

* added: Postgres support
* added: textfields for custom server_cfg/entry_list ini

## 1.2.1

* fixed: -1 value for tires out
* added: acweb now creates a default user when non exists "root" "root"
* added: license information on about page

## 1.2.0

* added configuration import
* added windows build (can't be used with docker)

## 1.1.0

* fixed: some default config values
* fixed: missing system and content directories for checksum calculation
* added changing dynamic track conditions
* added wind parameters
* added pitstop window configuration
* added fixed setups
* added ability to download configuration files of configurations and running instances
* added automatic calculation of sun angle
* added better logging and log level configuration

Note that you have to update the database schema. A migration script can be found within the db/ directory.

## 1.0.3

Now providing a docker image! Configuration is now done via environment variables, see README.md for details.

## 1.0.2

* fixed: added missing columns to schema.sql

## 1.0.1

* added cars from red pack
* added windows 64 bit build

## 1.0.0

* initial release
