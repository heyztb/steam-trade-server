# Configs

This document details what each of the config files is used for.

## Notes

### `account.json`

This file will contain login credentials as well as Steam Guard secrets used to generate OTP codes and to confirm trades. That being said, it is important that this information is protected. Ideally, you will only use this file a single time, and that is to populate your database using the `sts-seed-db` utility provided with this package. Delete it when you no longer need it.

### `haproxy.cfg`

This file contains an example HAProxy configuration suitable for Steam Trade Server. Take a look at it and modify it to your needs, or leave it as is. I've done my best to create a reasonable setup that should cater to a large number of concurrent users. Don't know what you're doing? Read the docs: <https://www.haproxy.com/documentation/hapee/latest/> (yes I know this is the enterprise documentation, but the differences between community and enterprise edition are not felt by this project.
