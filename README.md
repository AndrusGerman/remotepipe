# Remotepipe

allows to create remote pipes, you can create complex data flows and delegate heavy tasks to a remote computer


### Server Example
```bash
./remotepipe -server
```

### Scan devices
```bash
./remotepipe -scan
```


### Client Example


```bash
echo "andurs" | ./remotepipe 127.0.0.1 "cat /dev/stdin"

```


Encoding small file with ffmpeg
```bash
cat input.mp4 | ./remotepipe 127.0.0.1 "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" | ffplay -
```
```bash
cat input.mp4 | ./remotepipe 192.168.101.7 "ffmpeg -y -i pipe:0 -c:v libvpx-vp9 -f webm pipe:1" | ffplay -
```


Encoding a large file with ffmpeg
```bash
ffmpeg -i biginput.mp4  -codec copy -map 0 -movflags +faststart -f matroska pipe:1 | 
./remotepipe 192.168.101.7 "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" > newfile.mp4
```

```bash
ffmpeg -i input2.mp4 -codec copy -map 0 -movflags +faststart -f matroska pipe:1 | ./remotepipe scan "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" | ffplay -
```

searching for a text with grep
```bash
cat /home/andrusd/Documentos/app.py  | ./remotepipe 192.168.101.7 "grep hello"
```

find the first device
```bash
echo "hola" | ./remotepipe scan "cat /dev/stdin"
```


copy files client -> to server
```bash
echo "hello world" | ./remotepipe scan "cp /dev/stdin hello.txt"
```

copy files server -> to client
```bash
./remotepipe scan "cat bin.txt" > bin.txt 
```


### Motivation
With the need to render video without crashing my computer, I realized that my phone was powerful enough for this.
And I needed a way to connect parts of my pc's work with my phone's in real time, 'a data stream'
Pipes seemed to work for me, but I needed a simple way to do this over a tcp connection.
