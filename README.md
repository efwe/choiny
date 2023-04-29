Welcome to Choiny
=================

* checkout wipu db on mongo and analyze what to migrate
* checkout postGIS and setup (in podman)
* write migration-tool
* write new api-server
* write new frontend

Status Quo Wipu
===============




Podman Essentials
=================

```shell
podman run --name mongo -v /home/fw/Devel/mongo/db:/data/db -v /home/fw/Devel/mongo/backup:/data/backup  -p 27017:27017 -d 336f61db
5f26
podman exec -it mongo bash
```

Mongo Essentials
================

Again I lost valuable lifetime because of `mongorestore`

```shell
root@eb5cf392f6bf:/data# mongorestore --verbose --drop --stopOnError --objcheck --dryRun --nsInclude=wipu.* backup/
```

* we need to assume _sub-directories_ in the `backup` directory
* we need to give `--verbose` to see that the `dry-run` would work

Rustbelt
========

As we're starting from zero here it also does not matter what to note here

* https://crates.io/crates/geozero
* https://github.com/georust/geo
* https://docs.rs/crate/postgis/latest
* https://postgis.net/docs/PostGIS_Special_Functions_Index.html#PostGIS_GeographyFunctions
* https://postgis.net/docs/using_postgis_dbmanagement.html#EWKB_EWKT
* https://epsg.io/