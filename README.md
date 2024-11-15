Photo utils
---

Pull photos from Ricoh through wifi connection[[1]](https://github.com/clyang/GRsync):
```bash
python ricoh/ricoh.py -a
```

Converting JPG files to JXL:
```bash
go run go_images_xl/main.go ~/Desktop/Photos/2024_04_misc
```

Converting MOV files to mp4:
```bash
for mov in $(find . -type f -name "*.MOV"); do ffmpeg -i ${mov%.*}.MOV -vcodec libx264 -preset medium -b:v 1500k -acodec aac -b:a 128k ${mov%.*}.mp4; done
```

Pushing photos to local backup cluster:
```bash
for r in $(ls | grep -e "2024_10"); do rsync -z --ignore-existing --progress -h -r $r node-00@node-00:/home/node-00/dysio/Photos/; done
```

