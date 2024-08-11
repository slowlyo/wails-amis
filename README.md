# Wails amis

- [Wails](https://github.com/wailsapp/wails)
- [amis](https://github.com/baidu/amis)

## Before You Start

- You need to download the [sdk.tar.gz](https://github.com/baidu/amis/releases/latest) file and extract it to the
  `frontend/dist/public` directory.
    - After completing this operation, you should see the following file: `frontend/dist/public/sdk/sdk.js`.

## Notes

- Due to the fact that the amis SDK does not require compilation, the `wails dev` command is not applicable.
    - Instead, you can launch a file server to preview the frontend by running `go run dev.go` in the `frontend` directory.
- When performing a `wails build`, it is necessary to bypass the frontend compilation phase.
    - To do this, execute: `wails build -s`.
