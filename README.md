<div align="center">

<h1>
  <img src="https://raw.githubusercontent.com/rfyiamcool/golang_logo/refs/heads/master/gif/stop.gif" alt="stop" height="100px" style="vertical-align: middle; margin-left: 10px;" />
<!-- <a href="https://github.com/1hehaq/burf">BURF</a> -->
</h1>

</div>


>[!NOTE]
>**_A Go tool that generates backup file path lists for hostnames by appending common backup extensions. Ideal for fuzzing with tools like ffuf and similar utilities!_**
>
>Inspired from - [**`@h4x0r_dz`**](https://x.com/h4x0r_dz/status/1573318682230530048)

<!--
<h3>Features</h3>

- Generates common paths for specified hostnames.
- Supports intermediate variants for hostnames with dots (TLDs and subdomains).
- Allows users to provide extra file extensions.
- Can read hostnames and extensions from standard input.
- Allows reading hostnames and extra extensions from stdin or command line arguments
- Displays generated URLs in the console
- Error handling for incorrect input and usage messages
- Supports parsing comma-separated extra file extensions via the <kbd>-ext</kbd> flag
- Supports piping input to the tool via the <kbd>-ext</kbd> `pipe` option, allowing you to specify extra extensions using another file or command output.


<h3>Installation</h3>

-->


```bash
go install github.com/1hehaq/burf@latest
```

>[!TIP]
> ```bash
> cat extensions.txt | burf -host github.com -ext pipe | ffuf -w - -u https://github.com/FUZZ
> ```

### Flags

- <kbd>-host</kbd>: specify the hostname **(e.g: `github` or `sub.github.com`)**. If not provided, Burf will read from standard input!
- <kbd>-ext</kbd>: provide extra file extensions as a comma separated list or use `pipe` to read from standard input.
- <kbd>-both</kbd>: generate intermediate variants for hostnames with dots.
- <kbd>-help</kbd>: show help documentation.


<!--

<h3>Future plan</h3>

1. **Support for other file formats**: The current list of predefined file extensions is limited to certain types. Expanding the list to include other common formats such as PDF,
DOCX, XLSX, and audio/video files would increase its utility.
2. **Automatic detection of existing files**: Implementing a function that checks if the generated URLs point to existing files on the server could help prevent unnecessary time
spent downloading non-existent files.
3. **Recursive directory scanning**: Scanning recursively through directories to find all the files with specific extensions can save users from having to specify each file
individually.
4. **Output format customization**: Offering an option to modify the output format (e.g., JSON, CSV) could make it easier for users to process and analyze the generated URLs in other
applications.
5. **Integration with popular web browsers or download managers**: Integrating with web browsers or download managers would provide a seamless experience for users, allowing them to
directly open or save files from the console output.
6. **Support for wildcard characters**: Allowing users to specify wildcard characters (e.g., `*.html` or `*.{pdf,docx}`) in the file extension list could make it easier to match
multiple file types at once.
7. **Option to exclude specific file extensions**: Offering an option to exclude certain file extensions would give users more control over the output and help filter out unwanted
results.
8. **Error handling for duplicate URLs**: Implementing a check to prevent duplicates in the generated URLs could save time and reduce clutter in the console output.
9. **Option to specify custom hostnames**: Allowing users to specify custom hostnames would make the tool more versatile and useful for various use cases.
10. **Integration with version control systems (VCS)**: Integrating this tool with popular VCS like Git, SVN or Mercurial could help developers easily generate URLs for code
repositories, branches, commits, or tags.

sed 's/^/./'

-->

