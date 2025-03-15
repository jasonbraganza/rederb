# This file records raw thoughts, tasks, changes, and progress
## How I see the program functioning
- `rederb` with a path will just create a feed in the path (`rederb create`)
- `rederb` with nothing should show help
- `rederb create` will take the files in the current folder and create a feed according to defaults
- Empty, the site defaults to `http://localhost/audio-files-folder/feed-name.xml`
- Command line flags, that'll let me set site and path. (the final part of the path will always be the folder name)
    - for example: `http://localhost/fiction/audio-files-folder/feed-name.xml`
    - only one level for now. site -> category -> book-folder
- A config file, (yaml/env) that will hold defaults? or multiple choices?
    - go create --category=fiction will use fiction category in the url?
    - config file will also hold settings for feed, such as author, email, and site url





## Tasks as I see them
- [X] ~~Learn how pull requests work in go land~~
- [X] ~~Find dependencies, like I use in Python~~  
  godotenv(MIT)=python-dotenv  
  dhowden/tag(BSD-2)=tinytag  
  gorilla/feeds(BSD-3)-or-eduncan911/podcast(MIT)=feedgen).   
  Nothing set in stone, yet
- [X] ~~Update Project LICENSE once I figure out above~~ (BSD 2 Clause)
- [X] ~~Learn the basic go toolchain. does not have to be perfect. Learn industry standards later. Right now I desperately need derb to work on the Pi~~
- [X] ~~Figure out how to read in environment variables from a .env file~~
- [X] ~~Figure out to read metadata from audio files~~
- [X] Change the whole project into a Cobra app (2025-02-23) (done: 2025-02-24)
- [X] ~~Create an app that just saves config~~ (2025-03-15)
- [X] ~~Get path from command line~~
- [ ] verify path,
- [ ] split path to get end folder to tag on to url
- [ ] and prints things out
- [ ] Figure out how to write images from metadata to disk
- [ ] Write slowly. Learn how types work as you go.
- [ ] Blah blah blah  … Lots of missing bits (writing tests, et al)
- [ ] First run!
- [ ] Accept any folder, read the files in there and then write the feed back in there
- [ ] Learn how github releases work, so I can distribute stuff
- [ ] a `rederb init` command that creates a new config file in the current folder? or saves to `~/.config/rederb/rederb.yaml`

---