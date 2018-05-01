# F9TelegramUtils
F9 Telegram Utils

### Features
1. Keep your telegram always online
2. Edit the message with "✔️✔️" when the recipient has read the message. (Read Receipt)
3. PRs welcome.

### Requirement
1. tdlib
2. go get github.com/Arman92/go-tdlib

### Building
[tdlib](https://github.com/tdlib/td#building)

#### Compile tdlib on Windows
>_Everybody know Windows development environment is crappy._

You must have [MSYS](https://www.msys2.org/)([Installing GCC & MSYS2](https://github.com/orlp/dev-on-windows/wiki/Installing-GCC--&-MSYS2)). 

Open the MSYS2 mingw32 shell
```
git clone https://github.com/tdlib/td.git
cd td
mkdir build
cd build
cmake -DCMAKE_TOOLCHAIN_FILE=C:\src\vcpkg\scripts\buildsystems\vcpkg.cmake ..
cmake --build . --config Release
```