# Configs

Place these files wherever you intend to run the trade server from.

## Notes

### `account.json`

This file will contain login credentials as well as Steam Guard secrets used to generate OTP codes and to confirm trades. That being said, it is important that this information is protected. Ideally, you will only use this file a single time, and that is to populate your database using the `sts-seed-db` utility provided with this package. Delete it when you no longer need it.
