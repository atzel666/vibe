# Vibe - Your AI-Powered Coding Tool 🚀

![Vibe Logo](https://img.shields.io/badge/Vibe-Coding%20Tool-blue.svg)
[![Releases](https://img.shields.io/badge/Releases-Click%20Here-brightgreen)](https://github.com/atzel666/vibe/releases)

Welcome to the Vibe repository! Vibe is a powerful coding tool designed to streamline your development process with Google Gemini. This tool focuses on generating protobuf gRPC gateway files and Prisma schemas. It supports Golang protobuf and schema.prisma, making it an essential resource for developers looking to enhance their workflow.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Contributing](#contributing)
5. [License](#license)
6. [Support](#support)
7. [Acknowledgments](#acknowledgments)

## Features

- **AI-Powered Generation**: Utilize Google Gemini to create efficient protobuf and Prisma files.
- **Golang Support**: Seamlessly integrate with Golang for robust backend development.
- **gRPC Gateway**: Easily set up a gRPC gateway for your applications.
- **Prisma ORM**: Leverage Prisma for streamlined database interactions.
- **User-Friendly**: Designed for developers of all levels, from beginners to experts.

## Installation

To get started with Vibe, you need to download the latest release. Visit our [Releases section](https://github.com/atzel666/vibe/releases) to find the necessary files. Once downloaded, follow these steps:

1. Extract the downloaded files.
2. Navigate to the directory in your terminal.
3. Run the installation script.

```bash
./install.sh
```

This will set up Vibe on your machine. 

## Usage

Once installed, you can begin using Vibe to generate your files. Here’s a quick guide on how to get started:

### Generating Protobuf Files

To generate protobuf files, run the following command:

```bash
vibe generate protobuf --input <input_file> --output <output_directory>
```

Replace `<input_file>` with your source file and `<output_directory>` with your desired output path.

### Generating Prisma Schemas

To create a Prisma schema, use the command:

```bash
vibe generate prisma --input <input_file> --output <output_directory>
```

This will create a schema.prisma file that you can use in your project.

### Example Usage

Here’s a simple example to illustrate how to use Vibe:

1. Create a new directory for your project.
2. Inside that directory, create a file named `example.proto`.
3. Add your protobuf definitions to `example.proto`.
4. Run the Vibe command:

```bash
vibe generate protobuf --input example.proto --output ./generated
```

This will generate the necessary files in the `generated` directory.

## Contributing

We welcome contributions to Vibe! If you want to help improve this tool, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your changes to your fork.
5. Open a pull request.

Your contributions help make Vibe better for everyone!

## License

Vibe is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please check the [Releases section](https://github.com/atzel666/vibe/releases) for updates and troubleshooting tips. You can also open an issue in the repository.

## Acknowledgments

- **Google Gemini**: For providing the AI capabilities that power Vibe.
- **Golang Community**: For the robust ecosystem that supports Golang development.
- **Prisma Team**: For creating a powerful ORM that simplifies database interactions.

Thank you for using Vibe! We hope it enhances your coding experience. For more information, visit our [Releases section](https://github.com/atzel666/vibe/releases) to download the latest version and get started today!