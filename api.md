## 执行同步任务
### post /agent/script/sync
请求参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
interpreter | string | 是 | /bin/bash | 执行器
path | string | 是 | /tmp | 脚本执行的路径
user | string | 是 | root | 执行脚本的操作系统用户
content | string | 是 | #!/bin/bash\necho "OK" | 脚本内容
param |string | 是 | 1 2 | 脚本参数
timeout |int | 否 | 10 | 脚本超时时间
ip |string | 是 | 1.1.1.1 | 本地标示IP
task_id |string | 是 | f32ed3 | 任务ID

返回参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 任务状态码，0:正常;1:创建脚本失败;2:切换目录失败;3:切换用户失败
exit_code | int | 是 | 0 | 脚本的执行状态，-1:未执行; >=0表示脚本的返回码
stdout | string | 是 | 0 | 脚本标准输出
stderr | string | 是 | 0 | 脚本的异常输出
ip | string | 是 | 0 | 接口传入参数，原样返回
task_id | string | 是 | 0 | 任务ID
msg | string | 是 | 0 | 任务状态
start_time | string | 是 | 0 | 开始时间
finish_time | string | 是 | 0 | 结束时间

---
## 执行异步任务
### post /agent/script/async
请求参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
interpreter | string | 是 | /bin/bash | 执行器
path | string | 是 | /tmp | 脚本执行的路径
user | string | 是 | root | 执行脚本的操作系统用户
content | string | 是 | #!/bin/bash\necho "OK" | 脚本内容
param |string | 是 | 1 2 | 脚本参数
timeout |int | 否 | 10 | 脚本超时时间
ip |string | 是 | 1.1.1.1 | 本地标示IP
task_id |string | 是 | f32ed3 | 任务ID

返回参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 任务状态码，0:已接收并加入执行队列;1:未能加入队列;
msg | string | 是 | 0 | 任务状态

---
## 文件传输
### post /agent/file/transfer
请求参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
bcommand | string | 是 | cd /tmp&&ls | 文件传输前执行命令
bcommand_user | string | 是 | root | 文件传输前执行命令用户
acommand | string | 是 | cd /tmp&&ls | 文件传输后执行命令
acommand_user | string | 是 | root | 文件传输后执行命令用户
ip |string | 是 | 1.1.1.1 | 本地标示IP
task_id |string | 是 | f32ed3 | 任务ID
file_list | object[] |  |  | 文件传输列表
file_id | int | 是 | /tmp | 文件id
path | string | 是 | /tmp | 文件传输路径
user | string | 是 | root | 文件所属用户
group | string | 是 | root | 文件所属用户组
content | string | 是 | #!/bin/bash\necho "OK" | 文件内容
name | string | 是 | nginx.conf | 文件名称
md5_check | bool | 是 | true | true:校验原始文件md5;false:不校验原始文件md5，如果为true，md5校验失败就不传输文件
md5 | string | 否 | few3wc34 | 原始文件md5

返回参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 文件传输状态码
msg | string | 是 | 0 | 文件传输状态
ip | string | 是 | 0 | 接口传入参数，原样返回
command_before_exit_code | int | 是 | 0 | 文件传输前执行命令状态，-1:未执行; >=0表示脚本的返回码
command_before_stdout | string | 是 | 0 | 文件传输前执行命令标准输出
command_before_stderr | string | 是 | 0 | 文件传输前执行命令异常输出
command_after_exit_code | int | 是 | 0 | 文件传输后执行命令执行状态，-1:未执行; >=0表示脚本的返回码
command_after_stdout | string | 是 | 0 | 文件传输后执行命令执行标准输出
command_after_stderr | string | 是 | 0 | 文件传输后执行命令执行异常输出
task_id | string | 是 | 0 | 任务ID
file_list | object[] |  |  | 
file_id | int | 是 | 0 | 文件ID
state | bool | 是 | 0 | true:成功;false:失败
msg | string | 否 |  | 文件传输失败信息

---
## 拦截异步任务
### delete /agent/script/async
请求参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
ip |string | 是 | 1.1.1.1 | 本地标示IP
task_id |string | 是 | f32ed3 | 任务ID

返回参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 任务状态码，0:正常终止;1:终止失败
msg | string | 是 | 0 | 任务状态

---
## 异步任务回调
### post /agent/script/async/callback
请求参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 任务状态码，0:正常;1:创建脚本失败;2:切换目录失败;3:切换用户失败
exit_code | int | 是 | 0 | 脚本的执行状态，-1:未执行; >=0表示脚本的返回码
stdout | string | 是 | 0 | 脚本标准输出
stderr | string | 是 | 0 | 脚本的异常输出
ip | string | 是 | 0 | 接口传入参数，原样返回
task_id | string | 是 | 0 | 任务ID
msg | string | 是 | 0 | 任务状态
start_time | string | 是 | 0 | 开始时间
finish_time | string | 是 | 0 | 结束时间

返回参数
参数名 | 参数类型 | 是否必填 | 示例 | 说明
---|---|---|---|---
code | int | 是 | 0 | 任务状态码，0:正常终止;1:终止失败
msg | string | 是 | 0 | 任务状态