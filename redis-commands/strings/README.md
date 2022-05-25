## Store values 
SETEX command is used if we need to set a certain time after which the key will be removed from the database. The time provided is in seconds.

Instead of using the SETEX command, we can also use the SET command with the EX option as shown below
PSETEX time provided is in milliseconds.

SETNX command is used if a key is already present. If we run SET again, it will update value. SETNX will not update value

use the SET command with the NX option.
SET apple 400 NX

STRLEN => find the length of the value for a particular key.
MSET => save multiple records in a single command
MGET => get the value of multiple keys in a single command