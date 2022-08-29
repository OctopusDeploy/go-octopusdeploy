## `executionsapi` is a Temporary Package

The current design has the root `client` own an instance of `releases.ReleaseService`, so there's a
package dependency `client->releases`. 

As this code refers back to the `client`, it can't be in the `releases` package because 
that would create a circular dependency.

In the API client v3 we want to get rid of all the services, which would break 
the cycle, allowing us to put CreateReleaseV1 in the releases package, 
RunRunbookV1 in the runbooks package, etc (where they should be)... and delete this folder 
