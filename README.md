# Remotepipe

allows to create remote pipes, you can create complex data flows and delegate heavy tasks to a remote computer


### Server Example
```bash
./remotepipe -server true
```


### Client Example


```bash
echo "andurs" | ./remotepipe localhost "cat /dev/stdin"

```


Encoding small file with ffmpeg
```bash
cat input.mp4 | ./remotepipe localhost "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" | ffplay -
```
```bash
cat input.mp4 | ./remotepipe localhost "ffmpeg -y -i pipe:0 -c:v libvpx-vp9 -f webm pipe:1" | ffplay -
```


Encoding a large file with ffmpeg
```bash
ffmpeg -i biginput.mp4  -codec copy -map 0 -movflags +faststart -f matroska pipe:1 | 
./remotepipe localhost "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" > newfile.mp4
```

searching for a text with grep
```bash
cat /home/andrusd/Documentos/app.py  | ./remotepipe localhost "grep hello"
```

### Motivation
With the need to render video without crashing my computer, I realized that my phone was powerful enough for this.
And I needed a way to connect parts of my pc's work with my phone's in real time, 'a data stream'
Pipes seemed to work for me, but I needed a simple way to do this over a tcp connection.
