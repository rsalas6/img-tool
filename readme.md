# **Image Resizer CLI Tool**

A versatile command-line tool written in Go that offers bulk resizing of images within a specified directory. Supports a variety of image formats and provides several customization options.

## **Features:**

- Bulk resize images in a directory.
- Supports various image formats: JPEG, PNG, BMP, TIFF, WebP.
- Ability to target individual files.
- Intelligently skips already-processed images.
- Customizable dimensions and verbosity levels.

## **Usage:**

```
go run main.go [OPTIONS]
```

### **Options:**

- **\-w** or **\--width**: Width of the resized image. Default is **100**.
- **\-h** or **\--height**: Height of the resized image. Default is **100**.
- **\-f** or **\--file**: Specify a particular file to resize.
- **\-p** or **\--path**: Directory path containing images to resize. Default is **./imgs**.
- **\-v** or **\--verbose**: Enable verbose mode to display more detailed logs.

## **Example:**

To resize all images in the **./imgs** directory to a width and height of **200**:

```
go run main.go --width 200 --height 200
```

To resize a specific image **sample.jpg** to a width of **300**:

```
go run main.go --file sample.jpg --width 300
```

To resize images in a specific directory "C:\example\folder" to a width and height of 820 and 450 respectively:

```
go run main.go --path "C:\example\folder" --width 820 --height 450
```

## **Installation:**

1.  Ensure you have Go installed on your machine.
2.  Clone this repository.
3.  Navigate to the directory and run the tool as described in the usage section.


## **Compilation:**

To compile the `main.go` file and generate an executable, run the following command in your terminal:

```
go build -o img-tool main.go
```

This command will compile the Go code and generate an executable named img-tool.