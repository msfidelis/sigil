# Sigil, The Command Line of Door's


> Sigil doesn't matter, and it's not impressive. Actually, it's quite disappointing. It's just a nice wrapper for the lsof command intended to show occupied ports and their respective processes on the operating system.

> Just Lazy

# Installation Guide

> Sigil only works on Linux and MacOS. 

## Usage 

```
Usage:
  sigil [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  listening   Show host ports that are active and accepting connections

Flags:
  -h, --help   help for sigil

Use "sigil [command] --help" for more information about a command.
```

### Show Active Ports that are serving TCP connections

```bash
sigil listening

+-------+-------+---------+----------+
|  PID  | PORT  | PROCESS | PROTOCOL |
+-------+-------+---------+----------+
|  3731 | 49508 | Notion  | TCP      |
+-------+-------+---------+----------+
| 10634 |  8080 | php     | TCP      |
+-------+-------+---------+----------+
```

### Show Active Ports that are serving UDP connections

```bash
sigil listening --udp

+-----+-------+----------+----------+
| PID | PORT  | PROCESS  | PROTOCOL |
+-----+-------+----------+----------+
| 471 | 62524 | chronod  | UDP      |
+-----+-------+----------+----------+
| 483 | 63720 | sharingd | UDP      |
```