# Changelog

## [2.116.0](https://github.com/OctopusDeploy/go-octopusdeploy/compare/v2.115.0...v2.116.0) (2026-08-21)


### Features

* add GetDashboard and assert every query field carries a uri tag ([3f475dd](https://github.com/OctopusDeploy/go-octopusdeploy/commit/3f475dd30918a126d7124d4b1bace5fa0c64eb53)), closes [#442](https://github.com/OctopusDeploy/go-octopusdeploy/issues/442)
* add GetDynamicDashboard to DashboardService ([3f545ad](https://github.com/OctopusDeploy/go-octopusdeploy/commit/3f545ad03a01058faaf8f166e8936888dd6238ca))
* add interruptions.InterruptionType and use it on DashboardItem ([b263a38](https://github.com/OctopusDeploy/go-octopusdeploy/commit/b263a38fef044fcdc468c1a48e4f9bde5c70aa27))
* add Priority to the shared create-execution command ([c39c4a4](https://github.com/OctopusDeploy/go-octopusdeploy/commit/c39c4a475029c075ce14dc900a9a5c2499936658))
* add runbooks.Get for space-wide runbook queries ([9cb294e](https://github.com/OctopusDeploy/go-octopusdeploy/commit/9cb294e536618e1a8aa50a36b904b060679d35d8))
* add server health, timezones and document counts to serverstatus ([d4d2984](https://github.com/OctopusDeploy/go-octopusdeploy/commit/d4d2984ec8f42224b71b580568501bd5c170fbf7)), closes [#47](https://github.com/OctopusDeploy/go-octopusdeploy/issues/47)
* add system info, recent logs and system report to serverstatus ([9020da2](https://github.com/OctopusDeploy/go-octopusdeploy/commit/9020da27f867270f08cac89063cbcdda728005ab)), closes [#47](https://github.com/OctopusDeploy/go-octopusdeploy/issues/47)
* convert Amazon ECS cluster endpoints through EndpointResource ([18bea46](https://github.com/OctopusDeploy/go-octopusdeploy/commit/18bea4622d118d26c3f20b62144ee5be4dc6a2e5))
* deserialise Amazon ECS cluster target endpoints ([c902a51](https://github.com/OctopusDeploy/go-octopusdeploy/commit/c902a5153aff6db9535ca8a491f2b5dcf2aff500))


### Bug Fixes

* add actiontemplates.GetByQuery so collection filters reach the server ([cd2b594](https://github.com/OctopusDeploy/go-octopusdeploy/commit/cd2b5946e2cd80e6727850cc5428882b2c891b50)), closes [#437](https://github.com/OctopusDeploy/go-octopusdeploy/issues/437)
* complete DashboardItem and document that the dashboard filters on IDs ([fb3d95e](https://github.com/OctopusDeploy/go-octopusdeploy/commit/fb3d95e9ea3ce77d7cf58850b6d39908688adad5))
* model environment Links and add dashboard e2e coverage ([0518886](https://github.com/OctopusDeploy/go-octopusdeploy/commit/051888622310948d3f44865dbc30c23e519bf14b))
* tolerate a null ServiceDeskProjectName from the server ([16b9fae](https://github.com/OctopusDeploy/go-octopusdeploy/commit/16b9fae94f88612b69e41f1e78e8243b6e3a0b84))
* tolerate a null StandardChangeTemplateName from the server ([d50e716](https://github.com/OctopusDeploy/go-octopusdeploy/commit/d50e71640b1049603237741f32b166beda44c2d7))
