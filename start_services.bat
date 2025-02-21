@echo off
title zhihu服务启动脚本
echo 正在启动所有服务...

cd /d %~dp0

:: 设置最大等待时间(秒)
set TIMEOUT=30

:: 启动用户服务
echo 正在启动 User Service...
start "User Service" cmd /k "cd app\user && go run . -f etc/user.yaml"
call :wait_for_service 8000

:: 启动视频服务
echo 正在启动 Video Service...
start "Video Service" cmd /k "cd app\video && go run video.go -f etc/video.yaml"
call :wait_for_service 8001

:: 启动点赞服务
echo 正在启动 Like Service...
start "Like Service" cmd /k "cd app\like && go run like.go -f etc/like.yaml"
call :wait_for_service 8002

:: 启动评论服务
echo 正在启动 Comment Service...
start "Comment Service" cmd /k "cd app\comment && go run comment.go -f etc/comment.yaml"
call :wait_for_service 8003

:: 启动关注服务
echo 正在启动 Follow Service...
start "Follow Service" cmd /k "cd app\follow && go run follow.go -f etc/follow.yaml"
call :wait_for_service 8004

:: 启动聊天服务
echo 正在启动 Chat Service...
start "Chat Service" cmd /k "cd app\chat && go run chat.go -f etc/chat.yaml"
call :wait_for_service 8005

:: 启动推荐服务
echo 正在启动 Feed Service...
start "Feed Service" cmd /k "cd app\feed && go run feed.go -f etc/feed.yaml"
call :wait_for_service 8006

@REM :: 启动通知服务
@REM echo 正在启动 Notification Service...
@REM start "Notification Service" cmd /k "cd app\notification && go run notification.go -f etc/notification.yaml"
@REM call :wait_for_service 8007

:: 最后启动 API 网关
echo 正在启动 API Gateway...
start "API Gateway" cmd /k "cd app\applet && go run applet.go -f etc/applet-api.yaml"
call :wait_for_service 8888

echo 所有服务已成功启动！
pause
exit /b 0

:wait_for_service
:: 等待服务就绪
set port=%~1
set /a count=0
:check_loop
timeout /t 1 >nul
netstat -an | find ":%port%" | find "LISTENING" >nul
if %errorlevel% equ 0 (
    echo 服务已就绪: %port%
    exit /b 0
)
set /a count+=1
if %count% lss %TIMEOUT% goto check_loop
echo 等待服务超时: %port%
exit /b 1
