# Sigil, The Command Line of Door's

![Sigil](.github/assets/img/sigil.png)

> This is a passive-agressive project. Sigil doesn't matter, and it's not impressive. Actually, it's quite disappointing. It's just a nice wrapper for the lsof command intended to show occupied ports and their respective processes on the operating system.

> "Don't be Lazy or Sad. Be Lazy and Sad" - Matheus Fidelis (@fidelissauro)

> “How to avoid google every time you want to check a port” - Sergio Soares (@sergsoares)

# Installation Guide

> Sigil only works on Linux and MacOS. 

## MacOS amd64

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_darwin_arm64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```


## MacOS arm64

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_darwin_amd64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```

## Linux amd64 

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_linux_amd64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```

## Linux arm64 

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_linux_arm64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```

## Freebsd amd64 

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_freebsd_amd64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```

## Freebsd arm64 

```bash
wget https://github.com/msfidelis/sigil/releases/download/v0.0.2/sigil_0.0.2_freebsd_arm64 -O sigil 
mv sigil /usr/local/bin 
chmod +x /usr/local/bin/sigil
```

## Windows 

```bash
kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk
```

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

> Work's with --tcp flag too. This is a default option because i'm tired. 

```bash
sigil listening

+------+-------+-----------+----------+
| PID  | PORT  |  PROCESS  | PROTOCOL |
+------+-------+-----------+----------+
| 3731 | 49508 | Notion    | TCP      |
+------+-------+-----------+----------+
| 4369 | 15443 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 30021 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 31027 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 31624 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 61606 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 15021 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 30080 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 30443 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 |   443 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 |    80 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 |  8080 | com.docke | TCP      |
+------+-------+-----------+----------+
| 4369 | 15012 | com.docke | TCP      |
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


### Contributing

> Read this horrible code and send the PR 

#### Tests

```
sudo we don't need tests
```