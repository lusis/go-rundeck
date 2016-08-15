# `rundeck-http-get`
This utility allows you to grab data in a bit more flexible way than `rundeck-xml-get`.
The primary use case was for grabbing data where you need to specify an alternate content-type (as in key storage).
Otherwise it behaves the same as `rundeck-xml-get` by default.

## Usage
```
usage: rundeck-http-get [<flags>] <path>

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).
  --query_params=QUERY_PARAMS ...  
          key=value query parameter. specify multiple times if neccessary
  --content_type="application/xml"  
          an alternate content type if neccessary

Args:
  <path>  path to dump (e.g. executions/12234)
```

## Examples

### default format (xml)
```
$ bin/rundeck-http-get storage/keys/jv-pubkey | xmllint --format -

<?xml version="1.0"?>
<resource path="keys/jv-pubkey" type="file" url="http://rundeck/api/17/storage/keys/jv-pubkey" name="jv-pubkey">
  <resource-meta>
    <Rundeck-content-type>application/pgp-keys</Rundeck-content-type>
    <Rundeck-content-size>759</Rundeck-content-size>
    <Rundeck-content-creation-time>2016-08-15T13:14:31Z</Rundeck-content-creation-time>
    <Rundeck-content-modify-time>2016-08-15T13:14:31Z</Rundeck-content-modify-time>
    <Rundeck-auth-created-username>jvincent</Rundeck-auth-created-username>
    <Rundeck-auth-modified-username>jvincent</Rundeck-auth-modified-username>
    <Rundeck-key-type>public</Rundeck-key-type>
  </resource-meta>
</resource>
```

### json format
```
$ bin/rundeck-http-get storage/keys/jv-pubkey --content_type=application/json | python -mjson.tool

{
    "meta": {
        "Rundeck-auth-created-username": "jvincent",
        "Rundeck-auth-modified-username": "jvincent",
        "Rundeck-content-creation-time": "2016-08-15T13:14:31Z",
        "Rundeck-content-modify-time": "2016-08-15T13:14:31Z",
        "Rundeck-content-size": "759",
        "Rundeck-content-type": "application/pgp-keys",
        "Rundeck-key-type": "public"
    },
    "name": "jv-pubkey",
    "path": "keys/jv-pubkey",
    "type": "file",
    "url": "http://rundeck/api/17/storage/keys/jv-pubkey"
}
```

### raw data
Using the `Rundeck-context-type` value above, we can see what content-type we need to specify to get the actual data:

```
$ bin/rundeck-http-get storage/keys/jv-pubkey --content_type=application/pgp-keys

ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDgwJUtMLQTqfU9ICRmeFVi28vfLnlnBfEqzjuqpJT1dfR0gSA2+VT5wUu04wAaZ/3fmHS5GYBRDb7S494zLTDrD7ErkTdD0TU7I3CF3Bm4vrQkaxwcUR+0WacTjZW4dkSYjTi5h/Hohr654vX2hRXKCqTKh2jVGknd9BAWnuXINa0QnQgqVlLlt9tEfQI82JCaog95zUcO1NMw4HS5Eq0oaR5osLAZ5DKYgLRh9M4jOB7YRXgZgkHMQut2meTyYuNTLje0YNz8PtF+LDqNNrm+GQsp6yEhegjkhKD4gC4bRdpvyuDqZ0PxRWmFoWel1alWlzfXbKLU4HMTMYIH3I4Mll+oRKKtUVM+2IVCXEUSIxDTEY+0ivT8DkNdowYBWYrJy2dpDJ6C40d4crHpE++Pu5rrxm6Z7jP9oLSDN6yD8m8TQUvImFsZw0aO8iLm/C1haG0gsh5z4elL77wmGeJHxzyeO95y8UEqi6EeAtLJAIt752osBect8i9J0ceVleZG3+DJ71TaMEjYgEf1CJbRl6ZK+uEISAc/EG+oRNLek4DT2vTsth/95OS9D/ZmCJf3zRvUy9qimZMk7TmamKQRFR6v1BMvZ7s9d6cdhPZqfyIZHFkOvn26zapwny8Z82JL0fyyhzLZzjI0AVzgON5g9w72o2gHH6qvSjW7lhUVjQ==
```
