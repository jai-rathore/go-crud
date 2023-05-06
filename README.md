# lists

> lists - service to manage lists, etc.


- using [go modules](https://blog.golang.org/using-go-modules)
- [go-gin](https://github.com/gin-gonic/gin) for the web framework
- [swagger docs](http://localhost:3333/swagger/index.html) for api docs
- databases: mysql with support for flyway

## Contact
Squad: jaiadityarathore@gmail.com

## Project Structure
- `/src` - base code directory
- `src/config` - project configurations
- `src/config/secrets.json` - copy secrets json from fuze secrets
- `src/config/config.json` - application config
- `src/controllers` - api endpoints
- `src/docs` - swagger docs, need not be changed, will be updated via the make command below
- `src/middlewares` - everything to be added to the HTTP pipeline
- `src/services` - application / business logic
- `src/repositories` - sql
- `src/server` - router setup, services, repository initialization

## Setting Up

Install [Go](https://golang.org/dl/) >= `v1.20`

Recommended [VS Code](https://code.visualstudio.com/)

Install [Go for VS Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

Clone the `lists` repo

```bash
# Create your local `app_configs.json` file and update the configuration values appropriately
cp ./config/app_configs.template.json ./config/app_configs.json
# Install build dependencies
make install
```

## Replacing Application Secrets

Copy the whole secrets json from Fuze to the file `secrets.json`
Run the command below to replace the secrets from the `secrets.json` file in the `config.json` file

```bash
# in the root directory (will automatically download dependencies from go.mod file)
make configWithSecrets

```

You may need to format the config file `⌘K ⌘F` (Visual Studio Code)

## Creating Flyway Scripts

Creating a new flyway script using the bash script will automatically create a new file in `configs/flyway/sql/`
run the shell below or use makefile
```bash
# <script_name> usually formated FirstLastInitial_Ticket-Number_ShortName
# eg: LL_WCS-123_init
./new_flyway.sh <script_name>
```

flyway make example 
```bash
make flyway SCRIPTNAME=lists-4567_SmallNameHere
```
## Generating Swagger Docs

After adding a new route with [swagger comments](https://github.com/swaggo/swag#api-operation), regenerate the swagger to reflect those changes by running:

```bash
make generate
```

In case of error , export path
```bash
export PATH=$(go env GOPATH)/bin:$PATH
```

## Running the application locally

```bash
# in the root directory (will automatically download dependencies from go.mod file)
go run main.go

```

## Testing

In the root directory run :
```bash
go test -v ./...
```
this will run tests for all _test.go files in all subdirectories

To check unit-test coverage 
```bash
open test-coverage/coverage.html
```

or use the makefile command
```bash
make test-coverage
```