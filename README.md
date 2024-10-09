# GeoIP to HAProxy

> Original work from [andyjack](https://github.com/andyjack/geoip2-haproxy-ranges)

This tool allows you to create HAProxy compatible IP lists for each country using MAxMind databases. You will get a bunch of files that you can use to allow/block from frontend block in your HAProxy configuration.

## Configuration

The current image and commands use the free version (`GeoLite2-Country`). You will need an account and a license key. You can create your account [here](https://www.maxmind.com/en/geolite2/signup) and get the license key [here](https://www.maxmind.com/en/accounts/current/license-key).

If you have a membership or paid to get the `GeoIP2` data, you can replace `GeoLite2` inside the `curl` link.

## Running

### With docker

If you don't have `go` and don't want to install it, you can use the docker image. It will download the file using your informations and generate the files inside the `output` folder on your machine.

```sh
mkidr output
docker run --rm -e MAXMIND_ACCOUNT=ID -e MAXMIND_LICENSE="LICENSE" -v ./output:/output nioupola/geoip2-haproxy-ranges:latest
```

### Without docker

```sh
# if Debian/Ubuntu
sudo apt install golang
# if RedHat familly
sudo dnf install golang
curl -O -J -L -u $MAXMIND_ACCOUNT:$MAXMIND_LICENSE 'https://download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz'
tar zxf GeoLite2-Country_*.tar.gz
git clone https://github.com/Synoptik-Labs/geoip2-haproxy-ranges
cd geoip2-haproxy-ranges
go run . --db path/to/GeoLite2-Country.mmdb --destDir output
```

## References

* [Introduction to HAProxy ACLs](https://www.haproxy.com/blog/introduction-to-haproxy-acls)