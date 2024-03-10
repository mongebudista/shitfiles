#!/bin/bash

SESSION=dev
tmux new-session -d -s $SESSION
tmux split-window -v -l 30%
tmux split-window -h -l 66%
tmux split-window -h -l 50%

tmux attach -t $SESSION

