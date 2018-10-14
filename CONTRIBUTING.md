# Contributing

Thanks for considering contributing to the project!

Here are some things that might help:

## Local Integration Tests

To make development easier, run a local Octopus Deploy server on your machine. You can `vagrant up` [this image](https://github.com/MattHodge/VagrantBoxes/tree/master/OctopusDeployServer) to get a fully working Octopus Deploy Server.

When it comes up, login on [http://localhost:8081](http://localhost:8081) with username `Administrator` and password `OctoVagrant!`.

To get an API to use for local development, go to **Administrator | Profile | My API Keys** and click **New API Key**.

Set the two following environment variables:

```bash
# bash
export OCTOPUS_URL=http://localhost:8081/
export OCTOPUS_APIKEY=API-YOUR-API-KEY
```

```powershell
# PowerShell
$env:OCTOPUS_URL = "http://localhost:8081/"
$env:OCTOPUS_APIKEY = "API-YOUR-API-KEY"
```

You can now run integration tests.

## Tips

* You can open up the Swagger UI of the Octopus Deploy Server at [http://localhost:8081/swaggerui/index.html](http://localhost:8081/swaggerui/index.html). This makes it easier to poke around with the API.

* If you are trying to work out the API calls for something, use Chrome Developer Tools (or similar) to do them in the Octopus Web interface. You can then find out the data that was sent and which endpoints it was sent to.

![Chrome Developer Tools](https://i.imgur.com/TniEjnw.gif)

* It can be useful poking around the [Octopus Client built in .NET](https://github.com/OctopusDeploy/OctopusClients) to work out how things work. There are also some [API examples](https://github.com/OctopusDeploy/OctopusDeploy-Api) available.
