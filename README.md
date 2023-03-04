# task

A very simple shell task manager for Linux

# installation

### 1. clone git repo & compile go code yourself

```sh
git clone https://github.com/kstdx/task
cd task
go build main.go
# then the compiled bin's path is "./main"
```

### 2. add bin(named "task") to your dir & export path

```sh
mv ./main [YOUR_DIR]/task
```

Append to config file(like .bashrc):

```sh
export $PATH="$HOME/[YOUR_DIR]"
```

### 3. testing

```sh
task
# then this app will create a file "task.json"
```

# how to use

### `task`

If you already have "task.json" file, it shows you list of tasks.
But if you don't have "task.json" file, it inits "task.json" file.

### `task [TASK_NAME]`

It runs task command in "task.json" file
