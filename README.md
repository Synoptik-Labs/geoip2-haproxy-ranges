# GeoIP to HAProxy

## Generate ACL files per country for use with HAProxy

Inspired by https://github.com/berkayunal/haproxy-geoip-iprange, but for use with GeoIP2/GeoLite2 dbs.

Download the free versions from https://dev.maxmind.com/geoip/geoip2/geolite2/.

* [Introduction to HAProxy ACLs](https://www.haproxy.com/blog/introduction-to-haproxy-acls/)
* [Dovecot and HAProxy](https://wiki.dovecot.org/HAProxy)

## Running

```sh
# if Debian/Ubuntu
sudo apt install golang
# if RedHat familly
sudo dnf install golang
git clone https://github.com/Synoptik-Labs/geoip2-haproxy-ranges
cd geoip2-haproxy-ranges
go run . --db path/to/GeoLite2-Country.mmdb --destDir output
```
