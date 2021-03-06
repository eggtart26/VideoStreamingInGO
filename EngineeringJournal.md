# EngineeringJournal

## API design


                                    User  [user/:username]
                                  /      \
   upload/watch/download/delete  /        \ post
                                /          \
                            Recourse  <---- Comments
                            (Video)
    [user/:username/videos/:vid-id]        [/videos/:vid-id/comments/:comment-id]


create user : URL:/user Method: POST, SC: 201, 400, 500
User Login : URL:/user/:username Method: POST, SC: 200, 400, 500
Get user infomation : URL:/user/:username Method: GET, SC: 200, 400, 401, 403, 500
Delete user : URL:/user/:username Method: DELETE, SC: 204, 400, 401, 403, 500

List all videos : URL:/user/:username/videos Method: GET, SC: 200, 400, 500
Get a video : URL:/user/:username/videos/:vid-id Method: GET, SC: 200, 400, 500
Delete a video: URL:/user/:username/videos/:vid-id Method: DELETE, SC: 204, 400, 401, 403, 500

Show comments : URL:/videos/:vid-id/comments Method: GET, SC: 200, 400, 500
Post a comment : URL:/videos/:vid-id/comments Method: POST, SC: 201, 400, 500
Delete a comment : URL:/videos/:vid-id/comment/:comment-id Method: DELETE, SC: 204, 400, 401, 403, 500


## System Architecture
main.go -> handlers.go -> dbops/api.go -> defs -> response.go

handler -> validation{1.requesrt, 2.user} -> business logic -> reponse.
1. data model
2. error handling


## Database design

User
TABLE: users
id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT
login_name VARCHAR(64), UNIQUE KEY
pwd TEXT

Video  
TABLE: video_info
id VARCHAR(64), PRIMARY KEY, NOT NULL
author_id UNSIGNED INT
name TEXT
display_ctime TEXT
create_time DATETIME

Comment
TABLE: comments
id VARCHAR(64), PRIMARY KEY, NOT NULL
video_id VARCHAR(64)
author_id UNSIGNED INT
content TEXT
time DATETIME

Sessions
TABLE: sessions
session_id TINYTEXT, PRIMARY KEY, NOT NULL
TTL TINYTEXT
login_name VARCHAR(64)


Database relationship
  user      <---      session
   ^
   |
video_info  <---      comments



scheduler:
api->videoid-mysql
dispatcher->mysql-videoid->datachannel
executor->datachannel-videoid->delete videos