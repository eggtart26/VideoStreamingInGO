#! bin/bash

# Build web UI
cd ~/go/src/github.com/eggtart26/VideoStreamingInGO/web
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui/web
cp -R ~/go/src/github.com/eggtart26/VideoStreamingInGO/templates ~/go/bin/video_server_web_ui/web


# cd ~/go/bin/video_server_web_ui/web
# ./web