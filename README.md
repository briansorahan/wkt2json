# wkt2json

Convert wkt to geojson.

Usage:
```
echo 'POINT(0 0)' | wkt2json
```
outputs:
```
{"type":"Feature","geometry":{"type":"Point","coordinates":[0,0]},"properties":null}
```
