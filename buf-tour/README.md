# Buf Tour

In buf's default input mode, it assumes there is a `buf.yaml` in your current directory,
or uses the default values in lieu of a buf.yaml file. We recommend always having a `buf.yaml` file at the root of your `.proto` files hierarchy, as this is how .proto import paths are resolved.

```bash
# buf is configured with a buf.yaml file, which you can create with this command:
buf mod init
```