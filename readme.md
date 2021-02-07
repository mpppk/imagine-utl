# imagine-utl
![GitHub Actions](https://github.com/mpppk/imagine-utl/workflows/Go/badge.svg)

Utility tool for [imagine-app](https://github.com/mpppk/imagine)

## Installation

Download from [GitHub Releases](https://github.com/mpppk/imagine-utl/releases) 

If you are using macOS or Linux, you can also download by below command.

```
# If you are using linux, replace "darwin" to "linux" of URL.
$ curl -L https://github.com/mpppk/imagine-utl/releases/download/v0.1.0/imagine-utl_0.1.0_darwin_amd64.tar.gz \
  | tar xv
```

## Usage

### Load tags based on directory names

Suppose you have "Dog&Cat" dataset as below structure.

```
$ tree dataset
dataset/
├── dog
│   └── dog1.png
└── cat
    └── cat1.png
```

To label each image as "dog" or "cat", you can use `imagine-utl load` command.

```shell
$ imagine-utl load --dir path/to/dataset --depth 1
{ "name": "dog1.png", "path": "dog/dog1.png", "boundingBoxes": [{ "TagName": "dog" }] }
{ "name": "cat.png", "path": "cat/cat.png", "boundingBoxes": [{ "TagName": "cat" }] }
...
```

`--depth 1` means "Use directory name as label. The directory hierarchy going back is one level".
So if you specify `--depth 2`, output will look like this.

```shell
$ imagine-utl load --dir path/to/dataset --depth 2
{ "name": "dog1.png", "path": "dog/dog1.png", "boundingBoxes": [{ "TagName": "dog" }, {"TagName": "dataset"}] }
{ "name": "cat.png", "path": "cat/cat.png", "boundingBoxes": [{ "TagName": "cat" }, {"TagName": "dataset"}] }
```

There is little point in setting the depth to 2 or higher in this example.
However, it will be useful if the image is contained under a directory such as "test" or "train".