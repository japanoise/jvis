# jvis - json visualiser

Quick and dirty json visualiser. Up and down moves through the current node,
right goes in deeper, left goes up a level; vi and emacs keys work.

## Usage

`jvis` models a json stream as a tree; when the program starts it is at the root
node. The breadcrumb trail at the top of the screen shows how you got to your
current position. Navigating the tree should be comfortable to anyone familiar
with the vi or emacs editors.

- `q`, `C-c`: Quit
- `C-n`, `j`, `DOWN`: Go down
- `C-p`, `k`, `UP`: Go up
- `M-v`, `prior`, `page up`: Go up one screen
- `C-v`, `next`, `page down`: Go down one screen
- `M-<`, `g`: Go to start of current node
- `M->`, `G`: Go to end of current node
- `RET`, `l`, `C-f`, `RIGHT`: Browse child node
- `C-l`, `C-g`, `C-b`, `h`, `LEFT`: Go to parent node (or quit, if at root)
- `C-s`, `/`, `n`: Search forward/next search result
- `C-r`, `?`, `p`: Search back/previous search result
- `C-x`, `x`: Export current node to tsv (tab-seperated data; form `Key\tValue`)

## License

Copyright Japanoise 2018, licensed under the MIT license.

## Thanks

Thanks to jsonfui, which showed me what was possible.
