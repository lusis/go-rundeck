# example usage
`$ rundeck-run-job <jobid>`

# example usage with options
```
$ rundeck-run-job abcd105c-c358-4e36-b7fb-b748c03a9e8d
non-2xx response (code: 400): Job options were not valid: Option 'firstname' is required.
Option 'password' is required.
Option 'username' is required.

$ rundeck-run-job abcd105c-c358-4e36-b7fb-b748c03a9e8d -- -password foo -username bar -firstname bob
Job 2 is running
```
