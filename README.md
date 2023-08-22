# dellog

A simple tool to delete history log files

## Usage

- Write rule files

  ```
  # /etc/dellog.d/example.yaml
  
  match: /example/**/*.log
  size: 4g
  days: 4
  ---
  match: 
    - /example2/**/*.log
    - /example3/**/*.log
  size: 4g
  days: 4
  ```

- Execute `dellog`

  ```shell
  dellog
  ```

## Credits

GUO YANKE, MIT License
