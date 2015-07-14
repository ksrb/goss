#Setting up IntelliJ with Go

##[Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

##[A tour of Go](https://tour.golang.org/welcome/1)

##Notes
* [Plugin version list](https://plugins.jetbrains.com/plugins/alpha/list)

##Warning
* The IntelliJ go [plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) is extremely unstable on windows
* Plugin 0.9.371 cannot run single files from IntelliJ
* Plugin 0.9.223 can run single files from IntelliJ
* Plugin does not work with go 1.5 beta1 as [the selected directory is not a valid home for Go SDK](https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/1502)

##Option 1
Install 0.9.371 (<b>warning</b> you cannot run single files in IntelliJ)

* Install [go](https://golang.org/dl/)
* In IntelliJ:
    * ctrl+alt+shift+s &rarr; Plugins &rarr; Browser Repositories &rarr; Manage Repository &rarr; Add Repository &rarr;
     https://plugins.jetbrains.com/plugins/alpha/5047
     
##Option 2
Install 0.9.223 (<b>warning</b> this version is old, but allows you to run single files it is possible that 0.9.2x+ may be able to run single files in IntelliJ) 

* Download [0.9.223](https://plugins.jetbrains.com/plugin/download?updateId=19173)
* ctrl+alt+shift+s &rarr; Plugins &rarr; Install Plugins from Disk
* Navigate to where you downloaded 0.9.223

