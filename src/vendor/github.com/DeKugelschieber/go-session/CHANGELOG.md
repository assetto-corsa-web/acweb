# Changelog

## 2.0.1

* better error logging

## 2.0.0

* initial release

## 1.1.6

* bugfix: GC start in Init()

## 1.1.5

* added in package scope manager instance
    - use Init() function to create the instance
    - use Get() function to obtain it

## 1.1.4

* added http arguments to redirect function signature

## 1.1.3

* added middleware to check if session is set
* improved documentation

## 1.1.2

* changed NewMgoProvider return type, so that it returns errors
* session now checks if manager is set (to prevent nil ptr errors)
* remove function for session data

## 1.1.1

* added MongoDB provider

## 1.0.1

* bugfix: expire date when deleting sessions
* bugfix: create cookie no matter it exists already on CreateSession()

## 1.0.0

* initial release
