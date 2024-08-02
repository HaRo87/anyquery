---
title: Functions
description: Learn about the functions supported by AnyQuery
---

AnyQuery supports all the functions provided by SQLite as well as some additional functions. The functions are divided into two categories: SQLite functions and additional functions.

## SQLite functions

### Main functions

| Function name                | Reference                                                                  |
| ---------------------------- | -------------------------------------------------------------------------- |
| abs(X)                       | [Doc](https://www.sqlite.org/lang_corefunc.html#abs)                       |
| changes()                    | [Doc](https://www.sqlite.org/lang_corefunc.html#changes)                   |
| char(X1,X2,...,XN)           | [Doc](https://www.sqlite.org/lang_corefunc.html#char)                      |
| coalesce(X,Y,...)            | [Doc](https://www.sqlite.org/lang_corefunc.html#coalesce)                  |
| concat(X,...)                | [Doc](https://www.sqlite.org/lang_corefunc.html#concat)                    |
| concat_ws(SEP,X,...)         | [Doc](https://www.sqlite.org/lang_corefunc.html#concat_ws)                 |
| format(FORMAT,...)           | [Doc](https://www.sqlite.org/lang_corefunc.html#format)                    |
| glob(X,Y)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#glob)                      |
| hex(X)                       | [Doc](https://www.sqlite.org/lang_corefunc.html#hex)                       |
| ifnull(X,Y)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#ifnull)                    |
| iif(X,Y,Z)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#iif)                       |
| instr(X,Y)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#instr)                     |
| last_insert_rowid()          | [Doc](https://www.sqlite.org/lang_corefunc.html#last_insert_rowid)         |
| length(X)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#length)                    |
| like(X,Y)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#like)                      |
| like(X,Y,Z)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#like)                      |
| likelihood(X,Y)              | [Doc](https://www.sqlite.org/lang_corefunc.html#likelihood)                |
| likely(X)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#likely)                    |
| load_extension(X)            | [Doc](https://www.sqlite.org/lang_corefunc.html#load_extension)            |
| load_extension(X,Y)          | [Doc](https://www.sqlite.org/lang_corefunc.html#load_extension)            |
| lower(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#lower)                     |
| ltrim(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#ltrim)                     |
| ltrim(X,Y)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#ltrim)                     |
| max(X,Y,...)                 | [Doc](https://www.sqlite.org/lang_corefunc.html#max_scalar)                |
| min(X,Y,...)                 | [Doc](https://www.sqlite.org/lang_corefunc.html#min_scalar)                |
| nullif(X,Y)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#nullif)                    |
| octet_length(X)              | [Doc](https://www.sqlite.org/lang_corefunc.html#octet_length)              |
| printf(FORMAT,...)           | [Doc](https://www.sqlite.org/lang_corefunc.html#printf)                    |
| quote(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#quote)                     |
| random()                     | [Doc](https://www.sqlite.org/lang_corefunc.html#random)                    |
| randomblob(N)                | [Doc](https://www.sqlite.org/lang_corefunc.html#randomblob)                |
| replace(X,Y,Z)               | [Doc](https://www.sqlite.org/lang_corefunc.html#replace)                   |
| round(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#round)                     |
| round(X,Y)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#round)                     |
| rtrim(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#rtrim)                     |
| rtrim(X,Y)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#rtrim)                     |
| sign(X)                      | [Doc](https://www.sqlite.org/lang_corefunc.html#sign)                      |
| soundex(X)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#soundex)                   |
| sqlite_compileoption_get(N)  | [Doc](https://www.sqlite.org/lang_corefunc.html#sqlite_compileoption_get)  |
| sqlite_compileoption_used(X) | [Doc](https://www.sqlite.org/lang_corefunc.html#sqlite_compileoption_used) |
| sqlite_offset(X)             | [Doc](https://www.sqlite.org/lang_corefunc.html#sqlite_offset)             |
| sqlite_source_id()           | [Doc](https://www.sqlite.org/lang_corefunc.html#sqlite_source_id)          |
| sqlite_version()             | [Doc](https://www.sqlite.org/lang_corefunc.html#sqlite_version)            |
| substr(X,Y)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#substr)                    |
| substr(X,Y,Z)                | [Doc](https://www.sqlite.org/lang_corefunc.html#substr)                    |
| substring(X,Y)               | [Doc](https://www.sqlite.org/lang_corefunc.html#substr)                    |
| substring(X,Y,Z)             | [Doc](https://www.sqlite.org/lang_corefunc.html#substr)                    |
| total_changes()              | [Doc](https://www.sqlite.org/lang_corefunc.html#total_changes)             |
| trim(X)                      | [Doc](https://www.sqlite.org/lang_corefunc.html#trim)                      |
| trim(X,Y)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#trim)                      |
| typeof(X)                    | [Doc](https://www.sqlite.org/lang_corefunc.html#typeof)                    |
| unhex(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#unhex)                     |
| unhex(X,Y)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#unhex)                     |
| unicode(X)                   | [Doc](https://www.sqlite.org/lang_corefunc.html#unicode)                   |
| unlikely(X)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#unlikely)                  |
| upper(X)                     | [Doc](https://www.sqlite.org/lang_corefunc.html#upper)                     |
| zeroblob(N)                  | [Doc](https://www.sqlite.org/lang_corefunc.html#zeroblob)                  |


### Math functions

| Function name | Reference                                                |
| ------------- | -------------------------------------------------------- |
| acos(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#acos)    |
| acosh(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#acosh)   |
| asin(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#asin)    |
| asinh(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#asinh)   |
| atan(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#atan)    |
| atan2(Y,X)    | [Doc](https://www.sqlite.org/lang_mathfunc.html#atan2)   |
| atanh(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#atanh)   |
| ceil(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#ceil)    |
| ceiling(X)    | [Doc](https://www.sqlite.org/lang_mathfunc.html#ceil)    |
| cos(X)        | [Doc](https://www.sqlite.org/lang_mathfunc.html#cos)     |
| cosh(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#cosh)    |
| degrees(X)    | [Doc](https://www.sqlite.org/lang_mathfunc.html#degrees) |
| exp(X)        | [Doc](https://www.sqlite.org/lang_mathfunc.html#exp)     |
| floor(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#floor)   |
| ln(X)         | [Doc](https://www.sqlite.org/lang_mathfunc.html#ln)      |
| log(B,X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#log)     |
| log(X)        | [Doc](https://www.sqlite.org/lang_mathfunc.html#log)     |
| log10(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#log)     |
| log2(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#log2)    |
| mod(X,Y)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#mod)     |
| pi()          | [Doc](https://www.sqlite.org/lang_mathfunc.html#pi)      |
| pow(X,Y)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#pow)     |
| power(X,Y)    | [Doc](https://www.sqlite.org/lang_mathfunc.html#pow)     |
| radians(X)    | [Doc](https://www.sqlite.org/lang_mathfunc.html#radians) |
| sin(X)        | [Doc](https://www.sqlite.org/lang_mathfunc.html#sin)     |
| sinh(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#sinh)    |
| sqrt(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#sqrt)    |
| tan(X)        | [Doc](https://www.sqlite.org/lang_mathfunc.html#tan)     |
| tanh(X)       | [Doc](https://www.sqlite.org/lang_mathfunc.html#tanh)    |
| trunc(X)      | [Doc](https://www.sqlite.org/lang_mathfunc.html#trunc)   |

## Additional functions

### String functions

| Function name         | Usage                                                             | Alias                         |
| --------------------- | ----------------------------------------------------------------- | ----------------------------- |
| ascii(X)              | Returns the ASCII value of the first character of X.              | ord                           |
| bin(X)                | Returns the binary representation of X (string or integer).       |                               |
| bit_length(X)         | Returns the number of bits in X.                                  |                               |
| chr(X)                | Returns the character with the ASCII value of X.                  | char                          |
| length(X)             | Returns the length of X.                                          | char_length, character_length |
| elt(X,Y,...)          | Returns the Y-th element of the list X.                           |                               |
| elt_word(X,Y,delim)   | Returns the Y-th word of the string X.                            | split_part                    |
| field(X,Y,...)        | Returns the index of X in the list Y.                             |                               |
| find_in_set(X,Y)      | Returns the index of X in the list Y (comma-separated).           |                               |
| to_char(X,Y)          | Converts X to a string using the format Y.                        |                               |
| from_base64(X)        | Decodes the base64-encoded string X.                              |                               |
| to_base64(X)          | Encodes the string X to base64.                                   |                               |
| to_hex(X)             | Converts X to a hexadecimal string.                               |                               |
| from_hex(X)           | Converts the hexadecimal string X to a string.                    |                               |
| decode(X,Y)           | Decodes the string X using the encoding Y (base64, hex).          |                               |
| encode(X,Y)           | Encodes the string X using the encoding Y (base64, hex).          |                               |
| insert(X,Y,Z,N)       | Inserts the string N into X at position Y with length Z.          |                               |
| locate(X,Y,Z)         | Returns the position of X in Y starting from Z(optional).         | position, instr(from SQLite)  |
| lcase(X)              | Converts X to lowercase.                                          | lower                         |
| ucase(X)              | Converts X to uppercase.                                          | upper                         |
| left(X,Y)             | Returns the leftmost Y characters of X.                           |                               |
| right(X,Y)            | Returns the rightmost Y characters of X.                          |                               |
| load_file(X)          | Reads the file X and returns its content.                         |                               |
| load_file_bytes(X)    | Reads the file X and returns its content as bytes.                |                               |
| lpad(X,Y,Z)           | Pads the string X to length Y with Z on the left.                 |                               |
| rpad(X,Y,Z)           | Pads the string X to length Y with Z on the right.                |                               |
| octet_length(X)       | Returns the length of X in bytes.                                 |                               |
| to_octal(X)           | Converts X to an octal string.                                    |                               |
| regexp_replace(X,Y,Z) | Replaces the regular expression Y in X with Z.                    |                               |
| regexp_substr(X,Y,Z)  | Returns the substring of X that matches the regular expression Y. |                               |
| repeat(X,Y)           | Repeats the string X Y times.                                     |                               |
| reverse(X)            | Reverses the string X.                                            |                               |
| space(X)              | Returns a string of X spaces.                                     |                               |

### URL functions

| Function name      | Usage                                                           | Alias                                                    |
| ------------------ | --------------------------------------------------------------- | -------------------------------------------------------- |
| url_encode(X)      | Encodes the string X to a URL-encoded form.                     | urlEncode                                                |
| url_decode(X)      | Decodes the URL-encoded string X.                               | urlDecode                                                |
| domain(X)          | Returns the domain of the URL X.                                | urlDomain, url_domain                                    |
| path(X)            | Returns the path of the URL X.                                  | urlPath, url_path                                        |
| port(X)            | Returns the port of the URL X.                                  | urlPort, url_port                                        |
| url_query(X)       | Returns the query of the URL X.                                 | urlQuery                                                 |
| url_parameter(X,Y) | Returns the value of the parameter Y in the query of the URL X. | urlParameter, extract_url_parameter, extractUrlParameter |
| protocol(X)        | Returns the protocol of the URL X.                              | urlProtocol, url_protocol                                |

### Crypto functions

| Function name  | Usage                                                                 | Alias                                     |
| -------------- | --------------------------------------------------------------------- | ----------------------------------------- |
| md5(X)         | Returns the MD5 hash of the string X.                                 |                                           |
| sha1(X)        | Returns the SHA-1 hash of the string X.                               |                                           |
| sha256(X)      | Returns the SHA-256 hash of the string X.                             |                                           |
| sha384(X)      | Returns the SHA-384 hash of the string X.                             |                                           |
| sha512(X)      | Returns the SHA-512 hash of the string X.                             |                                           |
| blake2b(X)     | Returns the BLAKE2b hash of the string X.                             |                                           |
| blake2b_384(X) | Returns the BLAKE2b-384 hash of the string X.                         |                                           |
| blake2b_512(X) | Returns the BLAKE2b-512 hash of the string X.                         |                                           |
| random_float   | Returns a random float between 0 and 1.                               | random_real, random_double, randCanonical |
| rand           | Returns a random integer up to 4 294 967 295 (2^32 - 1).              | random_int                                |
| randn          | Returns a random integer between 0 and N.                             | random_intn                               |
| rand64         | Returns a random integer up to 18 446 744 073 709 551 615 (2^64 - 1). | random_int64                              |
| randn64        | Returns a random integer between 0 and N.                             | random_int64n                             |

### Other functions

| Function name         | Usage                                             | Alias |
| --------------------- | ------------------------------------------------- | ----- |
| clear_file_cache      | Clears the file cache for read_* table functions. |       |
| clear_plugin_cache(X) | Clears the plugin cache for the plugin X.         |       |