# Path reader

The application allows you to get a list of all files in a specified path, given the files of subfolders.  
Then 3 parallel threads process the resulting list. File handling is determining the size of the file.  
The result of the work is saved into a text file result.txt, where each line consists of two columns separated by a tabulation character. The first column is the path to the file, the second is the file size in bytes.  
At the end, the total size of all files in bytes is indicated.  

*The program may not have access to certain files and folders on the disk, such folders are ignored, and the files are ignored in the calculation of the total size, but they are present in the resulting file, their size is indicated as "error".*

---
**To specify the path, you must specify the flag (-path Path) when starting the program.**  
