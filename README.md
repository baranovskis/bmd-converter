## Introduction :memo:
**BmdConverter** is a MuOnline client file converting tool.

With this tool, you may simply convert your client files into **.csv** file format, and edit them in Microsoft Excel.

Supported Versions :white_check_mark:
--------
- Season 1 (tested on 1.01T+)
- Season 9 (from IGCN team)

## Synopsis :mag:
```sh
NAME:
   bmd-converter - tool for .bmd file conversion

USAGE:
   bmd-converter.exe [global options] command [command options] [arguments...]
   * bmd-converter.exe <client season> <bmd type> <file path>
   * bmd-converter.exe s1 text text.bmd - decode file to .csv
   * bmd-converter.exe s1 text text.csv - encode file to .bmd

VERSION:
   1.5

AUTHOR:
   NexT <info@baranovskis.dev>

COMMANDS:
   s1       Run season 1 file conversion
   s9       Run season 9 file conversion
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```