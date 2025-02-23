This was made only for me to learn Go.

# How does it work?

When you execute the project, you provide the path to the folder you want to organize as an argument.

The program will:

1. Traverse through all the files in the folder.
2. Create a subfolder for each file extension (e.g., '.txt', '.jpg').
3. Move files into their respective subfolders.
4. If a file has no extension, it will be placed in a 'MISC' folder.

### Example:

*Before organizing*:
```bash
MessyFolder/
├── file1.txt
├── file2.jpg
├── file3
├── notes.txt
├── image.png
```

*After organizing*:
```bash
MessyFolder/
├── TXT/
│   ├── file1.txt
│   ├── notes.txt
├── JPG/
│   ├── file2.jpg
├── PNG/
│   ├── image.png
├── MISC/
│   ├── file3
```

# How to use it?

First, clone the repository:

```bash
git clone github.com/felipemdutra/folder-organizer.git
cd folder-organizer
```

Build the project:

```bash
go build .
```

Then type the path to the folder you want to organize as an argument for the executable, for example:

```bash
./folder-organizer ~/ExampleFolder/MessyFolder
```

In this example, "MessyFolder" is going to be organized.

