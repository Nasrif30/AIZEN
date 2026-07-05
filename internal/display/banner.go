package display

import (
    "github.com/fatih/color"
)

const Warning = `
[!] WARNING: This tool is for educational and research purposes only.
[!] Use only in isolated lab environments. Do not deploy on production systems.
[!] Author: Nasrif30 — https://github.com/Nasrif30/AIZEN
`

const BannerRaw = `
  █████  ██ ███████ ███████ ███    ██ 
 ██   ██ ██    ███  ██      ████   ██ 
 ███████ ██   ███   █████   ██ ██  ██ 
 ██   ██ ██  ███    ██      ██  ██ ██ 
 ██   ██ ██ ███████ ███████ ██   ████ 

A — Adaptive      (shifts shape, evades detection)
I — Injection     (process hollowing, memory-only)
Z — Zero-Trust    (bypasses firewall & heuristics)
E — Execution     (no disk writes, no traces)
N — Network       (DNS/HTTPS/ICMP tunneling)

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⢺⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⡿⠉⠉⠉⠉⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⡜⢹⣿⣿⣿⣿⣵⡶⣤⣤⣀⣀⢻⣿⣧⣹⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠁⢸⣿⣿⣿⣿⡛⠙⠛⠓⠛⠙⣿⢧⡿⣿⠳⠏⠛⡟⢻⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠰⠋⣽⢻⣿⡿⠬⣄⣀⣀⣀⡼⠛⠀⢷⣌⣂⣀⣀⣠⢿⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢸⣿⡅⠀⠀⠀⠀⠀⠀⠀⢄⠀⠀⠀⠀⠀⠀⣷⣿
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠏⢷⠀⠀⠀⠀⠀⠀⡀⣄⠀⠀⠀⠀⠀⣰⣿⠉
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣧⡀⠀⠀⠀⠀⡈⠁⠀⠀⠀⠀⣰⢿⣿⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡠⢻⣧⢳⡄⠈⠩⣛⠛⠓⣒⡢⠉⣴⡟⣾⡏⠂
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⡀⢻⣦⠀⠀⠀⠀⠀⢀⣾⡿⠀⡏⣷⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣿⢣⡇⠈⠻⣷⣤⣤⣤⣴⣿⠟⠁⠀⡇⣿⣷
⠀⠀⠀⠀⠀⠀⢀⣠⣴⣾⣿⣿⣿⣿⢸⠀⠀⠀⠙⡛⠟⢹⣿⠃⠀⢀⠀⠃⣿⣿
⠀⠀⢀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⠸⠀⠀⠀⠀⠐⣶⡿⠁⠀⠀⡸⠰⢰⣿⣿
⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⡆⠀⠀⠀⠀⠀⠀⠀⠀⢰⠇⠆⣼⣿⣿
`


var Banner = Warning + "\n" + BannerRaw


var GreenBanner = color.GreenString(Warning + "\n" + BannerRaw)