# Go Training

A comprehensive training resource for learning Go programming with a focus on HTTP, APIs, and testing.

## About

This repository contains documentation and examples for learning Go programming. The content is organized into sections
covering HTTP basics, API interactions, and testing practices.

## Documentation

The documentation is built with [MkDocs](https://www.mkdocs.org/) using the Material theme. To view the documentation:

1. Install Python, MkDocs and the Material theme:
   ```
   brew install pyenv
   pyenv install 3.13.2 && pyenv global 3.13.2

   pip install mkdocs mkdocs-material
   ```

2. Run the documentation server locally:
   ```
   mkdocs serve
   ```

3. Open your browser and navigate to `http://localhost:8000`

## Content

The training materials cover:

- **HTTP Basics**: Making HTTP requests, parsing responses, and working with JSON
- **Testing**: Writing unit tests, table-driven tests, and running tests effectively

## Getting Started

Start with the "HTTP Basics" section to learn about making HTTP requests and processing responses. When you're ready to
learn about testing, check out the "Testing" section for comprehensive guides on writing and running tests in Go.