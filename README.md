# Remotepipe

allows to create remote pipes


```bash
echo "andurs" | ./remotepipe localhost "cat /dev/stdin"

```


```bash
cat input.mp4 | ./remotepipe localhost "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" | ffplay -
```



```bash
ffmpeg -i biginput.mp4  -codec copy -map 0 -movflags +faststart -f matroska pipe:1 | 
./remotepipe localhost "ffmpeg -y -i pipe:0 -c:v libx264 -f flv pipe:1" > newfile.mp4
```
