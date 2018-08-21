# NfsShares

## Caveats:
- FreeNAS requires that all POST and PUT API calls have a trailing slash.
- FreeNAS API describes quota, refquota, reservation and refreservation as string, but they return int. You can apparently send things like 10GiB and get an integer response
