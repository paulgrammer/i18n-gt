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

To run the program in development mode, run the following command:

```sh
make dev
```

To build the program, run the following command:

```sh
make build
```

This will compile the program and generate a binary file in the `build/` directory.

To build a Linux binary, run the following command:

```sh
make build-linux
```

This will compile the program for a Linux environment and generate a binary file in the `build/` directory.

### License

This program is licensed under the MIT license. See the [LICENSE](https://opensource.org/licenses/MIT) file for details.