curl -O -J -L -u $MAXMIND_ACCOUNT:$MAXMIND_LICENSE 'https://download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz'
tar zxf GeoLite2-Country_*.tar.gz
go run . --db GeoLite2-Country_*/GeoLite2-Country.mmdb --destDir /output
