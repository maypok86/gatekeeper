# Gatekeeper
![example workflow](https://github.com/maypok86/gatekeeper/actions/workflows/test.yml/badge.svg)
![Go Report](https://goreportcard.com/badge/github.com/maypok86/gatekeeper)
![License](https://img.shields.io/badge/license-MIT-green)

## General description
The service is designed to combat password brute-forcing during authorization in any system.
The service is called before the user is authorized and can either allow or block the attempt.
It is assumed that the service is used only for server-server, i.e. it is hidden from the end user.

## Algorithm of work
The service limits the frequency of authorization attempts for different combinations of parameters, for example:
* no more than `RATE_LOGIN = 10` attempts per minute for this login.
* no more than `RATE_PASSWORD = 100` attempts per minute for a given password (protection against reverse brute-force).
* no more than `RATE_IP = 1000` attempts per minute for a given IP (a large number, because NAT).

White/black lists contain lists of network addresses, which are handled in a simpler way.
If incoming ip is in whitelist - service unconditionally allows authorization (ok=true), if in blacklist - rejects (ok=false).

It uses [time/rate](https://pkg.go.dev/golang.org/x/time/rate) package, which implements the [token bucket](https://en.wikipedia.org/wiki/Token_bucket) algorithm.

## Description of API methods

### Authorization attempt
Request:
* login
* password
* ip

Response:
* ok (true/false) - the service should return ok=true if it thinks the request is normal
and ok=false if the request is brute-force.

### Reset bucket by login
* login

Must clear the bucket(s) corresponding to the passed login.

### Reset bucket by IP
* ip

Must clear the bucket(s) corresponding to the passed ip.

### Add IP to blacklist
* subnet (ip + mask)

### Remove IP from blacklist
* subnet (ip + mask)

### Add IP to whitelist
* subnet (ip + mask)

### Remove IP from whitelist
* subnet (ip + mask)

## Command-Line interface
A command-line interface for manual administration of the service is developed.
The CLI allows you to reset the bucket and manage the whitelist/blacklist.
The CLI works through GRPC interface.
