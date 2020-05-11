# Ben's Binary (bb)

Rather than a series of bash scripts and various environment variables being set, this project will hold cobra commands that replace the need for those scripts.

Configuration can be stored in `~/.bb/` such that each function can use shared configuration and customize it further. 

### Features:

 - Version command that checks current version and compares to repository of source codes releases
 - Versioned Releases using semantic versions, powered by github actions and github packages for docker images.
 
 ---
 
### Credits

 - #### [Go Generic Binary](https://github.com/Benbentwo/go-bin-generic/)
    This is my starter pack for go based command line tools. It provides semantic versioned releases on merges of pull requests. Allowing labels to determine the next release.
    