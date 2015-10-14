# v13
- Added support for username/password auth as non-admins cannot use api tokens. Precendence order is as follows:
  - If `RUNDECK_TOKEN` is set, that will be used
  - If `RUNDECK_TOKEN` is not set, `RUNDECK_USERNAME`/`RUNDECK_PASSWORD` are required.
  - Every request will attempt to reauthenticate with the rundeck server when username/password auth is being used

## Core API
- Changed API for `client.<method>` (i.e. `GET`,`PUT`,`POST`,`DELETE`) to be more flexible in the fact of rundeck API changes

Instead of passing the struct you want the XML unmarshaled into to each function, you are now required to pass in a pointer to a `[]byte` and handle unmarshalling yourself. 

- Removed the `client.RawGet` function as it's no longer neccessary
- Changed `ListExecutions`/DeleteAllExecutionsFor` to `ListProjectExecutions`/`DeleteAllExecutionsForProject`. See reasoning under bundled utilities changes.
## Bundled utilities
- Updated bundled utilities to use `v13` of the Rundeck API. There was only a single change between `v12` and `v13`:

```
Version 13:

New endpoints
/api/13/project/[PROJECT]/readme.md and /api/13/project/[PROJECT]/motd.md
Project Readme File (GET, PUT, DELETE)
```

which did not impact existing functionality.

- Renamed `rundeck-list-executions` to `rundeck-list-project-executions` as there are actually two execution list scopes that can be used in the Rundeck API
- Renamed `rundeck-delete-executions-for` to `rundeck-delete-executions-for-project`
- Some arguments to bins MAY have changed as I migrated from a basic argv approach to using kingpin.

# v12
Initial release
