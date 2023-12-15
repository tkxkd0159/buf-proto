# TOC

```bash
docker build -t jsproto .
docker run -it --name=ptest jsproto
docker run --name ptest --rm -v $(pwd):/workspace --workdir /workspace jsproto /bin/bash ./scripts/protogen.sh
```

## buf-tutorial
Practice buf & protocol buffer(any, reflection)