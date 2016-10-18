CloudFoundry-EnvironmentDemo
============================

This is a small, simple, Go application suitable for use as a demo application when deploying on cloud foundry.

Configuration for this application is performed by setting environment variables

| ENV Variable  | Description                                                                        | Default Value |
|---------------|------------------------------------------------------------------------------------|---------------|
| CFENV_PORT    | Set the port for the application to listen to                                      | 8080          |
| CFENV_BGCOLOR | Set the background color for the webpage. Useful for showing blue-green deployment | white         |
