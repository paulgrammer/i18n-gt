## Command-Line Program for Translating i18n json using Google Cloud Translation API

This is a simple command-line program written in Go that uses the Google Cloud Translation API to translate i18n json with text from one language to another. The program takes in a JSON file containing text to be translated, source language, target language and output JSON file path as required command-line arguments.

### Usage

To use the program, you'll need to have a Google Cloud service account and API key. You can create one by following the instructions in the [Google Cloud documentation](https://cloud.google.com/translate/docs/setup).

Once you have your credentials, you can run the program using the following command:

```
export GOOGLE_APPLICATION_CREDENTIALS="KEY_PATH"
```

```
i18n-gt -input <inputfile> -source <source language> -target <target language> -output <outputfile>
```

where `<inputfile>` is the path to the input JSON file, `<source language>` is the code for the source language (e.g. "en" for English), `<target language>` is the code for the target language, `<outputfile>` is the path to the output JSON file
### Input JSON Format

The input JSON file should contain an object, with property and value

Here's an example input JSON file:

```JSON
{
  "title": "Welcome!",
  "pages": {
    "title": "Pages",
    "profile": {
      "title": "Profile page",
      "login": {
        "pasword": "password"
      }
    }
  }
}
```

### Output JSON Format

The output JSON file will contain translated object

Here's an example output JSON file:

```JSON
{
  "pages": {
    "profile": {
      "login": {
        "pasword": "hasło"
      },
      "title": "Strona profilowa"
    },
    "title": "Strony"
  },
  "title": "Powitanie!"
}
```

### Dependencies

The program uses the following Go packages:

- `cloud.google.com/go/translate`: The Google Cloud Translation API Go client library
- `golang.org/x/text/language`: The Go standard library package for language code parsing
- `google.golang.org/api/option`: The Google Cloud API option library

You can install these packages using the `go get` command:

```
go get cloud.google.com/go/translate
go get golang.org/x/text/language
go get google.golang.org/api/option
```

To build the program for multiple platforms, run the following command:

```sh
make build
```

This will compile the program and generate a binary file in the `build/` directory.

To run the program in development mode, run the following command:

```sh
make dev
```

This will run the program with the `input.json` and `account.json` files provided and generate the `output.json` file.

## Build Targets

The `Makefile` provides the following build targets:

- `build-linux-amd64`: Build the program for Linux x64
- `build-linux-arm`: Build the program for Linux ARM
- `build-darwin-amd64`: Build the program for macOS x64
- `build-windows-amd64`: Build the program for Windows x64
- `build`: Build the program for all platforms
- `clean`: Clean up generated binary files

Here's an example of how to build the program for a specific platform:

```sh
make build-linux-amd64
```

This will compile the program for Linux x64 and generate a binary file in the `build/` directory.

## TODO

- [ ] Set up GitHub release workflow to automate binary file creation and release management.


### License

This program is licensed under the MIT license. See the [LICENSE](https://opensource.org/licenses/MIT) file for details.