# dellog

A simple tool to delete history log files

## Usage

* Write rule files

  ```
  # /etc/dellog.d/example
  # keep 3 days for /example/**/*.log
  /example/**/*.log:3
  ```

* Execute `dellog`

  ```shell
  dellog
  ```

## Credits

GUO YANKE, MIT License