# Setup
1. Download libmpv from [here](https://sourceforge.net/projects/mpv-player-windows/files/libmpv/)
2. Extract libmpv-2.dll to wails-mpv-example/
3. Rename libmpv-2.dll to libmpv.dll 

# Running

## Standalone

`go run -tags nocgo standalone/main.go`

## App

`wails dev -tags "nocgo"`